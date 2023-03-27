package makeimage

import (
	"fmt"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-logger/logger"
	"github.com/twpayne/go-geos"

	"go-uk-maps/makeimage/types"
)

func drawLayerGeometries(gc *draw2dimg.GraphicContext, mg types.MapLayerGeometries, scale func(x, y float64) (float64, float64)) error {
	for _, geoms := range mg.Lines {
		for _, g := range geoms {
			err := DrawGeom(gc, g, scale)
			if err != nil {
				return err
			}
		}
	}

	for _, geoms := range mg.Polygons {
		for _, g := range geoms {
			err := DrawGeom(gc, g, scale)
			if err != nil {
				return err
			}
		}
	}

	for _, geoms := range mg.Multi {
		for _, g := range geoms {
			err := DrawGeom(gc, g, scale)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func DrawGeom(gc *draw2dimg.GraphicContext, mg types.MapGeom, scale func(x, y float64) (float64, float64)) error {
	var drawfn func(geoms ...*geos.Geom) error

	drawfn = func(geoms ...*geos.Geom) error {
		for _, g := range geoms {
			_type := g.TypeID()
			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("type %+v [%+v] %+v %+v\n", _type, mg.DataType, mg.FillCol, mg.Thickness),
			)
			switch _type {
			case geos.TypeIDPoint:
				err := geom.DrawPoint(gc, g, mg.Radius, mg.FillCol, mg.Thickness, mg.LineCol, scale)
				if err != nil {
					return err
				}
			case geos.TypeIDLineString, geos.TypeIDLinearRing:
				err := geom.DrawLine(gc, g, mg.Thickness, mg.FillCol, mg.StrokeWidth, mg.LineCol, scale)
				if err != nil {
					return err
				}
			case geos.TypeIDPolygon:
				err := geom.DrawPolygon(gc, g, mg.FillCol, mg.LineCol, mg.Thickness, scale)
				if err != nil {
					return err
				}
			case geos.TypeIDMultiPoint, geos.TypeIDMultiLineString, geos.TypeIDMultiPolygon, geos.TypeIDGeometryCollection:
				n := g.NumGeometries()
				var subgeoms []*geos.Geom
				for i := 0; i < n; i++ {
					subgeoms = append(subgeoms, g.Geometry(i))
				}
				err := drawfn(subgeoms...)
				if err != nil {
					return err
				}
			default:
				logger.Log(
					logger.LVL_FATAL,
					fmt.Sprintf("unknown geometry type %v", _type),
				)
			}
		}

		return nil
	}

	return drawfn(mg.Geometry)
}
