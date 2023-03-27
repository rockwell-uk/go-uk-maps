package makeimage

import (
	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-text/fonts"

	"go-uk-maps/makeimage/types"
)

func GetTypeFace(gc *draw2dimg.GraphicContext, l *types.ImageLayer, t types.TextStyle) fonts.TypeFace {
	var fontSpacing float64
	var fontData draw2d.FontData
	var fontDataStyle draw2d.FontStyle

	var fontName string = t.FontName
	var fontSize float64 = t.Size * l.Zoom
	var fillCol color.RGBA = t.Color

	switch fontName {
	case "regular":
		fontDataStyle = draw2d.FontStyleNormal
	default:
		fontDataStyle = draw2d.FontStyleBold
	}

	fontData = draw2d.FontData{
		Name:   fontName,
		Family: draw2d.FontFamilySans,
		Style:  fontDataStyle,
	}

	return fonts.TypeFace{
		Name:  fontName,
		Size:  fontSize,
		Color: fillCol,
		StrokeStyle: draw2d.StrokeStyle{
			Color: t.Stroke.Color,
			Width: t.Stroke.Width * l.Zoom,
		},
		BackgroundColor: t.BgCol,
		BackgroundStrokeStyle: draw2d.StrokeStyle{
			Color: t.BgStroke.Color,
			Width: t.BgStroke.Width * l.Zoom,
		},
		Spacing:  fontSpacing,
		FontData: fontData,
		Face:     fonts.GetFace(gc, fontData, fontSize),
	}
}
