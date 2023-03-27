package makeimage

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-text/fonts"
	"github.com/twpayne/go-geos"

	"go-uk-maps/makeimage/types"
)

func prepareLabel(gc *draw2dimg.GraphicContext, l *types.ImageLayer, a types.LayerArtifact, scale func(x, y float64) (float64, float64)) error {
	var ml types.MapLabel
	var dataType types.DataType
	var g *geos.Geom
	var typeFace fonts.TypeFace
	var rotation float64

	dataType = a.ShapeType.DataType
	g = a.Geometry
	_type := g.TypeID()

	// Named Places - we have font sizes defined so override the "default"
	if dataType == types.CARTOGRAPHIC {
		var textSize string = a.CartographicFontDetails.FontHeight

		labelStyle, err := types.GetLabelStyleByTextSize(textSize)
		if err != nil {
			return err
		}
		a.LabelStyle = labelStyle

		_, err = GetLabelText(a.LabelText, a.LabelFieldNames[0])
		if err != nil {
			return err
		}

		rotation = a.CartographicFontDetails.Orientation
		if err != nil {
			return err
		}
	}

	// construct the basic maplabel
	mlBase := &types.MapLabel{
		ID:            a.ID,
		LayerType:     a.LayerType,
		FeatCode:      a.FeatCode,
		Geometry:      g,
		DataType:      dataType,
		Rotation:      rotation,
		LineCol:       a.ShapeType.LineCol,
		FillCol:       a.ShapeType.FillCol,
		Zoom:          l.Zoom,
		Radius:        a.RenderType.DotRadius * l.Zoom,
		Thickness:     a.RenderType.LineThickness * l.Zoom,
		ShouldSplitFn: types.ShouldSplit,
	}

	if _type == geos.TypeIDPoint {
		// first element of LabelText is DISTNAME
		label, err := GetLabelText(a.LabelText, a.LabelFieldNames[0])
		if err != nil {
			return err
		}

		if label != "" {
			typeFace = GetTypeFace(gc, l, a.LabelStyle.LabelTextStyle)

			labelType := types.LABEL_POINT
			if a.FeatCode == types.FEAT_MOTORWAY_JN {
				labelType = types.LABEL_NUMBER
			}

			// build label for a point
			ml = *mlBase
			ml.LabelType = labelType
			ml.Label = label
			ml.TypeFace = typeFace

			err := ml.Build(scale)
			if err != nil {
				return err
			}

			if dataType == types.CARTOGRAPHIC {
				if _, exists := l.Labels.Labels[a.FeatCode]; !exists {
					l.Labels.Labels[a.FeatCode] = make(map[string]types.MapLabel)
				}

				l.Labels.Labels[a.FeatCode][a.ID] = ml
				l.Labels.LabelsOrder[a.FeatCode] = append(l.Labels.LabelsOrder[a.FeatCode], a.ID)
			} else {
				if _, exists := l.Labels.Points[a.FeatCode]; !exists {
					l.Labels.Points[a.FeatCode] = make(map[string]types.MapLabel)
				}

				l.Labels.Points[a.FeatCode][a.ID] = ml
				l.Labels.PointsOrder[a.FeatCode] = append(l.Labels.PointsOrder[a.FeatCode], a.ID)
			}
		}
	}

	if _type == geos.TypeIDMultiLineString || _type == geos.TypeIDLineString {
		if dataType == types.LINE {
			// first element of LabelText is DISTNAME
			label, err := GetLabelText(a.LabelText, a.LabelFieldNames[0])
			if err != nil {
				return err
			}

			if label != "" {
				typeFace = GetTypeFace(gc, l, a.LabelStyle.LabelTextStyle)

				// build label for a name
				ml = *mlBase
				ml.Label = label
				ml.LabelType = types.LABEL_NAME
				ml.TypeFace = typeFace

				err := ml.Build(scale)
				if err != nil {
					return err
				}

				if _, exists := l.Labels.Names[a.FeatCode]; !exists {
					l.Labels.Names[a.FeatCode] = make(map[string]types.MapLabel)
				}

				l.Labels.Names[a.FeatCode][a.ID] = ml
				l.Labels.NamesOrder[a.FeatCode] = append(l.Labels.NamesOrder[a.FeatCode], a.ID)
			}

			// second element of LabelText is ROADNUMBER
			label, err = GetLabelText(a.LabelText, a.LabelFieldNames[1])
			if err != nil {
				return err
			}

			if label != "" {
				typeFace = GetTypeFace(gc, l, a.LabelStyle.LabelBackgroundStyle)

				// build label for a number
				ml = *mlBase
				ml.Label = label
				ml.LabelType = types.LABEL_NUMBER
				ml.TypeFace = typeFace

				err := ml.Build(scale)
				if err != nil {
					return err
				}

				if _, exists := l.Labels.Numbers[a.FeatCode]; !exists {
					l.Labels.Numbers[a.FeatCode] = make(map[string]types.MapLabel)
				}

				l.Labels.Numbers[a.FeatCode][a.ID] = ml
				l.Labels.NumbersOrder[a.FeatCode] = append(l.Labels.NumbersOrder[a.FeatCode], a.ID)
			}
		}
	}

	return nil
}
