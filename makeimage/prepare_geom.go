package makeimage

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/twpayne/go-geos"
	"golang.org/x/image/draw"

	"go-uk-maps/makeimage/types"
)

func prepareGeom(m draw.Image, l *types.ImageLayer, assetsAdded types.AssetsAdded, a types.LayerArtifact, scale func(x, y float64) (float64, float64)) (types.AssetsAdded, error) {
	featCode := a.FeatCode
	dataType := a.ShapeType.DataType

	gtype := a.Geometry.TypeID()

	if icon, ok := types.CommunityServices[featCode]; ok {
		src, _ := png.Decode(bytes.NewReader(icon))

		iconWidth := 16 * l.Zoom
		iconHeight := 16 * l.Zoom

		decIcon := image.NewRGBA(image.Rect(0, 0, int(iconWidth), int(iconHeight)))
		draw.NearestNeighbor.Scale(decIcon, decIcon.Rect, src, src.Bounds(), draw.Over, nil)

		x := a.Geometry.X()
		y := a.Geometry.Y()
		x, y = scale(x, y)

		offset := image.Pt(int(x-iconWidth/2), int(y-iconHeight/2))

		opacity := 35
		alpha := (256 * opacity) / 100

		mask := image.NewUniform(color.Alpha{uint8(alpha)})
		zeroPoint := image.Point{X: 0, Y: 0}
		draw.DrawMask(m, decIcon.Bounds().Add(offset), decIcon, zeroPoint, mask, zeroPoint, draw.Over)

		if _, exists := assetsAdded[a.FeatCode]; !exists {
			assetsAdded[a.FeatCode] = make(map[types.LabelType]map[string]types.ImageAsset)
		}
		if _, exists := assetsAdded[a.FeatCode][types.LABEL_ICON]; !exists {
			assetsAdded[a.FeatCode][types.LABEL_ICON] = make(map[string]types.ImageAsset)
		}

		assetsAdded[a.FeatCode][types.LABEL_ICON][a.ID] = types.ImageAsset{
			ID:        a.ID,
			LayerType: a.LayerType,
			FeatCode:  a.FeatCode,
			Geometry:  a.Geometry,
		}
	} else {
		mg := types.MapGeom{
			ID:          a.ID,
			LayerType:   a.LayerType,
			Geometry:    a.Geometry,
			DataType:    dataType,
			FeatCode:    featCode,
			LineCol:     a.ShapeType.LineCol,
			FillCol:     a.ShapeType.FillCol,
			Radius:      a.RenderType.DotRadius * l.Zoom,
			Thickness:   a.RenderType.LineThickness * l.Zoom,
			StrokeWidth: a.RenderType.StrokeWidth * l.Zoom,
		}

		switch gtype {
		case geos.TypeIDPoint:
			l.Geometries.Points[dataType] = append(l.Geometries.Points[dataType], mg)
		case geos.TypeIDLineString, geos.TypeIDLinearRing:
			l.Geometries.Lines[dataType] = append(l.Geometries.Lines[dataType], mg)
		case geos.TypeIDPolygon:
			l.Geometries.Polygons[dataType] = append(l.Geometries.Polygons[dataType], mg)
		case geos.TypeIDMultiPoint, geos.TypeIDMultiLineString, geos.TypeIDMultiPolygon, geos.TypeIDGeometryCollection:
			l.Geometries.Multi[dataType] = append(l.Geometries.Multi[dataType], mg)
		default:
			return map[types.FeatCode]map[types.LabelType]map[string]types.ImageAsset{}, fmt.Errorf("unknown geometry type %v", gtype)
		}
	}

	return assetsAdded, nil
}
