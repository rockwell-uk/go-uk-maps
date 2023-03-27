package makeimage_test

import (
	"image"
	"image/draw"
	"testing"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-text/fonts"
	"github.com/twpayne/go-geos"

	"go-uk-maps/colours"
	"go-uk-maps/makeimage/types"
)

var (
	gctx = geos.NewContext()
)

func TestDrawLabel(t *testing.T) {
	dimX := 1024
	dimY := 768
	fontName := "bold"
	fontSize := 64.0

	scale := func(x, y float64) (float64, float64) {
		return x, y
	}

	m := image.NewRGBA(image.Rect(0, 0, dimX, dimY))
	draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
	gc := draw2dimg.NewGraphicContext(m)
	gc.SetDPI(72)

	// strokestyle
	strokeStyle := draw2d.StrokeStyle{
		Color: colours.White,
		Width: 1.0,
	}

	// font
	fontData := draw2d.FontData{
		Name:   fontName,
		Family: draw2d.FontFamilySans,
		Style:  draw2d.FontStyleNormal,
	}

	backgroundStrokeStyle := draw2d.StrokeStyle{
		Color: colours.Blue,
		Width: 1.0,
	}

	typeFace := fonts.TypeFace{
		StrokeStyle:           strokeStyle,
		Color:                 colours.Black,
		Size:                  fontSize,
		FontData:              fontData,
		Face:                  fonts.GetFace(gc, fontData, fontSize),
		BackgroundStrokeStyle: backgroundStrokeStyle,
	}

	labels := []struct {
		labelText string
		centerWkt string
	}{
		{
			"David Boyle",
			"POINT(700 400)",
		},
		{
			"Absolutely Astounding",
			"POINT(200 500)",
		},
		{
			"A680",
			"POINT(200 100)",
		},
	}

	for _, tt := range labels {
		g, err := gctx.NewGeomFromWKT(tt.centerWkt)
		if err != nil {
			t.Fatal(err)
		}

		ml := types.MapLabel{
			Label:         tt.labelText,
			TypeFace:      typeFace,
			Rotation:      0,
			Geometry:      g,
			ShouldSplitFn: types.ShouldSplit,
		}
		err = ml.Build(scale)
		if err != nil {
			t.Fatal(err)
		}

		// draw the label background
		err = ml.DrawBackground(gc)
		if err != nil {
			t.Fatal(err)
		}

		err = ml.DrawText(gc)
		if err != nil {
			t.Fatal(err)
		}

		err = savePNG("test-output/maplabel_draw_test.png", m)
		if err != nil {
			t.Fatal(err)
		}
	}
}
