package makeimage_test

import (
	"fmt"
	"image"
	"image/draw"
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

func TestDrawBackground(t *testing.T) {
	type testLabel struct {
		name string
		wkt  string
	}

	tests := map[string]struct {
		tileCenter nationalgrid.Location
		dim        int
		zoom       float64
		testLabels []testLabel
		fontData   draw2d.FontData
		fontSize   float64
		fontStroke float64
	}{
		"collision": {
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 292897.53,
					Lon: 65571.92,
				},
			},
			dim:  600,
			zoom: float64(1),
			testLabels: []testLabel{
				{
					name: "Combe Pafford",
					wkt:  "POINT(291342 66584)",
				},
				{
					name: "Lummaton Hill",
					wkt:  "POINT(291158 66575)",
				},
				{
					name: "St Marychurch",
					wkt:  "POINT(292033 65838)",
				},
			},
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:   float64(8),
			fontStroke: float64(1),
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

		strokeStyle := draw2d.StrokeStyle{
			Color: colours.White,
			Width: tt.fontStroke,
		}

		backgroundStrokeStyle := draw2d.StrokeStyle{
			Color: colours.White,
			Width: tt.fontStroke,
		}
		rotation := 0.0

		scale := func(x, y float64) (float64, float64) {
			x = envelope.Px(x) * tileWidth
			y = tileHeight - (envelope.Py(y) * tileHeight)
			return x, y
		}

		m := image.NewRGBA(image.Rect(0, 0, tt.dim, tt.dim))
		draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
		gc := draw2dimg.NewGraphicContext(m)
		gc.SetDPI(72)

		for _, l := range tt.testLabels {
			g, err := gctx.NewGeomFromWKT(l.wkt)
			if err != nil {
				t.Fatal(err)
			}

			typeFace := fonts.TypeFace{
				StrokeStyle:           strokeStyle,
				Color:                 colours.Pink,
				Size:                  tt.fontSize,
				FontData:              tt.fontData,
				Face:                  fonts.GetFace(gc, tt.fontData, tt.fontSize),
				BackgroundColor:       colours.Pink,
				BackgroundStrokeStyle: backgroundStrokeStyle,
			}

			ml := types.MapLabel{
				Label:         l.name,
				TypeFace:      typeFace,
				Rotation:      rotation,
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

			// mark the label center
			ctr := *ml.Dimensions.Center
			gc.SetFillColor(colours.Black)
			err = geom.DrawDot(gc, 1, ctr[0], ctr[1])
			if err != nil {
				t.Fatal(err)
			}
		}

		err = savePNG(fmt.Sprintf("test-output/named_place_test_%v.png", name), m)
		if err != nil {
			t.Fatal(err)
		}
	}
}
