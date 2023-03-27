package types

import (
	"reflect"

	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-nationalgrid"
	"github.com/twpayne/go-geos"
)

type TileRequest struct {
	Location   nationalgrid.Location
	TileWidth  float64
	TileHeight float64
	Zoom       float64
	Layers     []string
	Format     string
	Quality    int
	TileGeom   *geos.Geom
}

func (r TileRequest) Validate() error {
	var errorList []string

	if !r.LocationIsValid() {
		errorList = append(errorList, "tile location is not defined")
	}

	return ErrorMsg(errorList)
}

func (r TileRequest) LocationIsValid() bool {
	return !reflect.DeepEqual(r.Location, nationalgrid.Location{})
}

func (r TileRequest) Envelope() *geos.Bounds {
	tileWidth, tileHeight := r.Dims()
	tileOrigin := r.Origin()

	return geos.NewBounds(tileOrigin.LatLon.Lat, tileOrigin.LatLon.Lon, tileOrigin.LatLon.Lat+tileWidth, tileOrigin.LatLon.Lon+tileHeight)
}

func (r TileRequest) Dims() (float64, float64) {
	var tileWidth, tileHeight float64

	zoom := r.Zoom
	if zoom == float64(0) {
		zoom = float64(1)
	}

	switch r.Location.Type {
	case nationalgrid.NATIONALGRID.String():

		gridRef, _ := nationalgrid.ParseGridRef(r.Location.GridRef)

		switch {
		case gridRef.Quadrant != "":

			tileWidth = nationalgrid.QuadrantSize
			tileHeight = nationalgrid.QuadrantSize

		case gridRef.SubSquare != "":

			tileWidth = nationalgrid.SubSquareSize
			tileHeight = nationalgrid.SubSquareSize

		case gridRef.Square != "":

			tileWidth = nationalgrid.SquareSize
			tileHeight = nationalgrid.SquareSize
		}
	default:
		tileWidth = r.TileWidth * 10 / zoom
		tileHeight = r.TileHeight * 10 / zoom
	}

	return tileWidth, tileHeight
}

func (r TileRequest) Origin() nationalgrid.OSGB36LatLon {
	tileWidth, tileHeight := r.Dims()
	tileCenter := r.Location.ToOSGB36()

	return nationalgrid.OSGB36LatLon{
		LatLon: nationalgrid.LatLon{
			Lat: tileCenter.LatLon.Lat - tileWidth/2,
			Lon: tileCenter.LatLon.Lon - tileHeight/2,
		},
	}
}

func (r TileRequest) BoundsCoords() ([]float64, []float64) {
	tileWidth, tileHeight := r.Dims()
	tileOrigin := r.Origin()

	return []float64{
			tileOrigin.LatLon.Lat,
			tileOrigin.LatLon.Lon + tileHeight,
		}, []float64{
			tileOrigin.LatLon.Lat + tileWidth,
			tileOrigin.LatLon.Lon,
		}
}

func (r TileRequest) BoundsGeom() (*geos.Geom, error) {
	tileWidth, tileHeight := r.Dims()
	tileOrigin := r.Origin()

	g, err := geom.BoundsGeom(
		tileOrigin.LatLon.Lat,
		tileOrigin.LatLon.Lat+tileWidth,
		tileOrigin.LatLon.Lon,
		tileOrigin.LatLon.Lon+tileHeight,
	)
	if err != nil {
		return &geos.Geom{}, err
	}

	return g, nil
}
