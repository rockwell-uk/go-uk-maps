package makeimage_test

import (
	"fmt"
	"image"
	"image/draw"
	"strings"
	"testing"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-nationalgrid"
	"github.com/rockwell-uk/go-text/fonts"

	apitypes "go-uk-maps/api/types"
	"go-uk-maps/colours"
	"go-uk-maps/makeimage/types"
)

func TestRoadLabel(t *testing.T) {
	tests := map[string]struct {
		dim        int
		fontData   draw2d.FontData
		fontSize   float64
		fontStroke draw2d.StrokeStyle
		text       string
		geomWKT    string
		tileCenter nationalgrid.Location
		zoom       float64
	}{
		"A6060": {
			dim: 600,
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleNormal,
			},
			zoom:     1.2,
			fontSize: 9,
			fontStroke: draw2d.StrokeStyle{
				Color: colours.White,
				Width: float64(9) / float64(5),
			},
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 389721.19853198,
					Lon: 413215.07842109,
				},
			},
			text:    "A6060",
			geomWKT: "MULTILINESTRING((388888.99999999994 413242.9999997675,388874 413258.9999997683))",
		},
	}

	for name, tt := range tests {
		r := apitypes.TileRequest{
			Location:   tt.tileCenter,
			TileWidth:  float64(tt.dim),
			TileHeight: float64(tt.dim),
			Zoom:       tt.zoom,
			Quality:    100,
		}

		bounds, err := r.BoundsGeom()
		if err != nil {
			t.Fatal(err)
		}

		envelope, err := geom.ToEnvelope(bounds)
		if err != nil {
			t.Fatal(err)
		}

		tileHeight := r.TileHeight
		tileWidth := r.TileWidth

		scale := func(x, y float64) (float64, float64) {
			x = envelope.Px(x) * tileWidth
			y = tileHeight - (envelope.Py(y) * tileHeight)
			return x, y
		}

		m := image.NewRGBA(image.Rect(0, 0, tt.dim, tt.dim))
		draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
		gc := draw2dimg.NewGraphicContext(m)
		gc.SetDPI(72)

		// road label
		g, err := gctx.NewGeomFromWKT(tt.geomWKT)
		if err != nil {
			t.Fatal(err)
		}

		// draw the envelope - this should be covered
		e := g.Bounds()
		fillColor := colours.White
		strokeColor := colours.Blue
		width := 1.0
		err = geom.DrawPolygon(gc, e.Geom(), fillColor, strokeColor, width, scale)
		if err != nil {
			t.Fatal(err)
		}

		// face first
		face := fonts.GetFace(gc, tt.fontData, tt.fontSize)

		rotation := 0.0

		// strokestyle
		strokeStyle := draw2d.StrokeStyle{
			Color: colours.White,
			Width: 1.0,
		}

		backgroundStrokeStyle := draw2d.StrokeStyle{
			Color: colours.White,
			Width: 1.0,
		}

		// font
		fontData := draw2d.FontData{
			Name:   "bold",
			Family: draw2d.FontFamilySans,
			Style:  draw2d.FontStyleNormal,
		}

		typeFace := fonts.TypeFace{
			StrokeStyle:           strokeStyle,
			Color:                 colours.Pink,
			Size:                  tt.fontSize,
			FontData:              fontData,
			Face:                  face,
			BackgroundColor:       colours.Pink,
			BackgroundStrokeStyle: backgroundStrokeStyle,
		}

		ml := types.MapLabel{
			Label:    tt.text,
			TypeFace: typeFace,
			Rotation: rotation,
			Geometry: g,
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

		err = savePNG(fmt.Sprintf("test-output/roadlabel_test_%v.png", strings.ToLower(name)), m)
		if err != nil {
			t.Fatal(err)
		}
	}
}
