package makeimage

import (
	"fmt"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-text/fonts"
	"github.com/rockwell-uk/go-text/text"

	apitypes "go-uk-maps/api/types"
	"go-uk-maps/makeimage/types"
)

func drawLayerLabels(r apitypes.TileRequest, assetsAdded types.AssetsAdded, gc *draw2dimg.GraphicContext, ml types.MapLayerLabels, scale func(x, y float64) (float64, float64)) error {
	for _, f := range featCodeOrder {
		for _, key := range ml.LabelsOrder[f] {
			err := drawLabel(r, assetsAdded, gc, ml.Labels[f][key], scale)
			if err != nil {
				return err
			}
		}

		for _, key := range ml.NamesOrder[f] {
			err := drawLabel(r, assetsAdded, gc, ml.Names[f][key], scale)
			if err != nil {
				return err
			}
		}

		for _, key := range ml.NumbersOrder[f] {
			err := drawLabel(r, assetsAdded, gc, ml.Numbers[f][key], scale)
			if err != nil {
				return err
			}
		}

		for _, key := range ml.PointsOrder[f] {
			err := drawLabel(r, assetsAdded, gc, ml.Points[f][key], scale)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func drawLabel(r apitypes.TileRequest, assetsAdded types.AssetsAdded, gc *draw2dimg.GraphicContext, l types.MapLabel, scale func(x, y float64) (float64, float64)) error {
	if l.TypeFace.Face == nil {
		l.TypeFace.Face = fonts.GetFace(gc, l.TypeFace.FontData, l.TypeFace.Size)
	}

	g := l.Geometry

	fonts.SetFont(gc, l.TypeFace)

	switch l.LabelType {
	case types.LABEL_POINT:

		if shouldBeAdded(r, assetsAdded, l) {
			err := geom.DrawPoint(gc, g, l.Radius, l.FillCol, l.Thickness, l.LineCol, scale)
			if err != nil {
				return err
			}
			err = l.DrawText(gc)
			if err != nil {
				return err
			}
		}

	case types.LABEL_NAME:

		g, err := geom.ScaleLine(g, scale)
		if err != nil {
			return err
		}

		// road numbers arent bold
		fontData := draw2d.FontData{
			Name:   l.TypeFace.Name,
			Family: draw2d.FontFamilySans,
			Style:  draw2d.FontStyleNormal,
		}

		face := fonts.GetFace(gc, fontData, l.TypeFace.Size)
		l.TypeFace.Face = face
		fonts.SetFont(gc, l.TypeFace)

		if shouldBeAdded(r, assetsAdded, l) {
			glyphs, _ := text.TextAlongLine(gc, l.Label, *geom.GetPoints(g), l.TypeFace)

			for _, glyph := range glyphs {
				err := geom.DrawRune(gc, glyph.Pos, face, glyph.Rotation, glyph.Char)
				if err != nil {
					return err
				}
			}
		}

	case types.LABEL_NUMBER:

		if shouldBeAdded(r, assetsAdded, l) {
			// draw the label background
			err := l.DrawBackground(gc)
			if err != nil {
				return err
			}

			err = l.DrawText(gc)
			if err != nil {
				return err
			}
		}

	case types.LABEL_ICON:
		return nil

	default:
		return fmt.Errorf("LABEL TYPE UNKNOWN: [%+v] %+v", l.LabelType, l)
	}

	return nil
}
