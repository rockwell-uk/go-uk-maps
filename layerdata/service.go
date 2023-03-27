package layerdata

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/rockwell-uk/go-logger/logger"
	"github.com/rockwell-uk/go-nationalgrid"
	"github.com/rockwell-uk/go-utils/sliceutils"
	"github.com/twpayne/go-geos"

	"go-uk-maps/api/types"
	"go-uk-maps/database"
)

var (
	filter    = true
	filterLvl = 0.0001
	reduce    = false
	simplify  = false
	gctx      = geos.NewContext()
)

func GetLayerData(db database.StorageEngine, r types.TileRequest) ([]string, map[string][]LayerData, error) {
	var targetTileGeom *geos.Geom
	var squares map[string][]int
	var layersToLoad []string
	var records map[string][]LayerData

	logger.Log(
		logger.LVL_DEBUG,
		fmt.Sprintf("tileRequest %+v\n", r),
	)

	tileWidth, tileHeight := r.Dims()

	logger.Log(
		logger.LVL_DEBUG,
		fmt.Sprintf("tileWidth %+v\n", tileWidth),
	)

	logger.Log(
		logger.LVL_DEBUG,
		fmt.Sprintf("tileHeight %+v\n", tileHeight),
	)

	if len(r.Layers) == 0 || reflect.DeepEqual(r.Layers, []string{"all"}) {
		layersToLoad = defaultLayers
	} else {
		for _, q := range r.Layers {
			for _, t := range database.LayerTypes {
				if q == t {
					layersToLoad = append(layersToLoad, t)
				}
			}
		}
	}

	logger.Log(
		logger.LVL_DEBUG,
		fmt.Sprintf("layersToLoad %+v\n", layersToLoad),
	)

	envelope := r.Envelope()

	logger.Log(
		logger.LVL_DEBUG,
		fmt.Sprintf("tile envelope %+v\n", envelope),
	)

	// Get the sectors relevant to the requested tile
	squares = nationalgrid.GetSubSquares(envelope)
	logger.Log(
		logger.LVL_DEBUG,
		fmt.Sprintf("squares %+v\n", squares),
	)

	targetTileGeom, err := r.BoundsGeom()
	if err != nil {
		return layersToLoad, records, err
	}

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("targetTileGeom %+v\n", targetTileGeom),
	)

	logger.Log(
		logger.LVL_INTERNAL,
		"getting records ...",
	)

	// Get the records relevant to the squares
	records, err = getRecords(db, squares, targetTileGeom, layersToLoad)
	if err != nil {
		return layersToLoad, records, err
	}

	return layersToLoad, records, nil
}

func getRecords(db database.StorageEngine, subSquaresMap map[string][]int, targetTileGeom *geos.Geom, layersToLoad []string) (map[string][]LayerData, error) {
	records := make(map[string][]LayerData)

	for _, layerName := range layersToLoad {
		items := []LayerData{}

		logger.Log(
			logger.LVL_INTERNAL,
			fmt.Sprintf("layer %+v\n", layerName),
		)

		fields, orderBy := getDBParams(layerName, db.GetGeomFn())

		for square, subSquares := range subSquaresMap {
			where := "WHERE GRIDREF IN ("
			numSubSquares := len(subSquares) - 1

			for i, subSquare := range subSquares {
				tpl := "%v,"
				if i == numSubSquares {
					tpl = "%v"
				}
				where += fmt.Sprintf(tpl, subSquare)
			}
			where += ") "

			q := fmt.Sprintf("SELECT %s FROM %s %s %s", fields, db.GetTableName(layerName, square), where, orderBy)

			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("sql %+v\n", q),
			)

			rows, err := db.GetDB(layerName).Queryx(q)
			if errors.Is(err, sql.ErrNoRows) {
				logger.Log(
					logger.LVL_WARN,
					err.Error(),
				)
				continue
			}
			defer rows.Close()

			l := LayerData{}

			for rows.Next() {
				err := rows.StructScan(&l)
				if err != nil {
					logger.Log(
						logger.LVL_FATAL,
						err.Error(),
					)
				}

				if filter {
					polyGeom, err := geos.NewGeomFromWKT(l.WKT.String)
					if err != nil {
						return records, err
					}
					polyEnv := polyGeom.Bounds()

					inView := isInView(polyEnv, targetTileGeom.Bounds())
					if !inView {
						continue
					}

					if simplify && sliceutils.ContainsString(layersToSimplify, layerName) {
						l.WKT = simplifyGeom(polyGeom, filterLvl)
					}

					if reduce {
						inView := isInView(polyEnv, targetTileGeom.Bounds())
						if inView {
							l.WKT = reducePoly(polyGeom, targetTileGeom)
						}
					}
				}

				items = append(items, l)
			}

			records[layerName] = items
		}
	}

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("finished loading records %v\n", len(records)),
	)

	return records, nil
}

// is the geometry fully or partially in view?
func isInView(item *geos.Bounds, target *geos.Bounds) bool {
	intersects := item.Intersects(target)
	within := item.Contains(target)

	return intersects || within
}

func simplifyGeom(polyGeom *geos.Geom, lvl float64) sql.NullString {
	newGeom := polyGeom.Simplify(lvl)
	wkt := newGeom.ToWKT()

	return sql.NullString{
		String: wkt,
		Valid:  true,
	}
}

func reducePoly(polyEnv *geos.Geom, tilePolyGeom *geos.Geom) sql.NullString {
	newGeom := tilePolyGeom.Intersection(polyEnv)
	wkt := newGeom.ToWKT()

	return sql.NullString{
		String: wkt,
		Valid:  true,
	}
}

func getDBParams(layerType, geomFn string) (string, string) {
	var fields, orderBy string

	switch layerType {
	case "administrative_boundary",
		"building",
		"electricity_transmission_line",
		"foreshore",
		"glasshouse",
		"ornament",
		"railway_tunnel",
		"road_tunnel",
		"surface_water_area",
		"surface_water_line",
		"tidal_water",
		"woodland":
		fields = fmt.Sprintf("id, featcode, %s(ogc_geom) AS wkt", geomFn)

	case "railway_track",
		"roundabout",
		"tidal_boundary":
		fields = fmt.Sprintf("id, classifica, featcode, %s(ogc_geom) AS wkt", geomFn)

	case "functional_site",
		"railway_station":
		fields = fmt.Sprintf("id, distname, classifica, featcode, %s(ogc_geom) AS wkt", geomFn)

	case "motorway_junction":
		fields = fmt.Sprintf("id, junctnum, featcode, %s(ogc_geom) AS wkt", geomFn)

	case "named_place":
		fields = fmt.Sprintf("id, distname, htmlname, classifica, fontheight, orientatio, featcode, %s(ogc_geom) AS wkt", geomFn)
		// orderBy = "ORDER BY FIELD(fontheight, 'Large', 'Medium', 'Small')"

	case "road":
		fields = fmt.Sprintf("id, distname, roadnumber, classifica, drawlevel, override, featcode, %s(ogc_geom) AS wkt", geomFn)

	case "spot_height":
		fields = fmt.Sprintf("id, height, featcode, %s(ogc_geom) AS wkt", geomFn)

	default:
		fields = "*"
	}

	orderBy = "ORDER BY FEATCODE DESC"

	return fields, orderBy
}
