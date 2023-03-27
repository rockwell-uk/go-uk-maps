package makeimage_test

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strings"
	"testing"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-nationalgrid"
	"github.com/rockwell-uk/go-text/text"

	apitypes "go-uk-maps/api/types"
	"go-uk-maps/colours"
	"go-uk-maps/makeimage"
	"go-uk-maps/makeimage/types"
)

func TestEnvelope(t *testing.T) {
	noSplit := func(s string) bool {
		return false
	}

	type testLabel struct {
		name             string
		wkt              string
		expectedEnvelope string
	}

	tests := map[string]struct {
		tileCenter    nationalgrid.Location
		dim           int
		zoom          float64
		testLabels    []testLabel
		fontData      draw2d.FontData
		fontSize      float64
		fontStroke    float64
		shouldSplitFn func(s string) bool
	}{
		"Spotland_Road_Zoom_5": {
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388800,
					Lon: 413800,
				},
			},
			dim:  600,
			zoom: 5,
			testLabels: []testLabel{
				{
					name:             "Spotland Road",
					wkt:              "LINESTRING(388603.3400000000256114 413816.2299999999813735, 388662.0000000000000000 413776.0000000000000000, 388682.0000000000000000 413767.0000000000000000, 388697.9799999999813735 413767.5700000000069849)",
					expectedEnvelope: "POLYGON ((210.4499999999999886 258.0000000000000000, 389.5500000000000114 258.0000000000000000, 389.5500000000000114 342.0000000000000000, 210.4499999999999886 342.0000000000000000, 210.4499999999999886 258.0000000000000000))",
				},
			},
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:      float64(8),
			fontStroke:    float64(1),
			shouldSplitFn: text.ShouldSplit,
		},
		"Spotland_Road_Zoom_05": {
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388800,
					Lon: 413800,
				},
			},
			dim:  600,
			zoom: 0.5,
			testLabels: []testLabel{
				{
					name:             "Spotland Road",
					wkt:              "LINESTRING(388603.3400000000256114 413816.2299999999813735, 388662.0000000000000000 413776.0000000000000000, 388682.0000000000000000 413767.0000000000000000, 388697.9799999999813735 413767.5700000000069849)",
					expectedEnvelope: "POLYGON ((291.0500000000000114 295.8000000000000114, 308.9499999999999886 295.8000000000000114, 308.9499999999999886 304.1999999999999886, 291.0500000000000114 304.1999999999999886, 291.0500000000000114 295.8000000000000114))",
				},
			},
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:      float64(8),
			fontStroke:    float64(1),
			shouldSplitFn: text.ShouldSplit,
		},
		"A680": {
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388800,
					Lon: 413800,
				},
			},
			dim:  600,
			zoom: 5,
			testLabels: []testLabel{
				{
					name:             "A680",
					wkt:              "POINT(388600 413800)",
					expectedEnvelope: "POLYGON ((249.8050000000000068 278.0000000000000000, 350.1949999999999932 278.0000000000000000, 350.1949999999999932 322.0000000000000000, 249.8050000000000068 322.0000000000000000, 249.8050000000000068 278.0000000000000000))",
				},
			},
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:      float64(8),
			fontStroke:    float64(1),
			shouldSplitFn: text.ShouldSplit,
		},
		"David_Boyle": {
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388800,
					Lon: 413800,
				},
			},
			dim:  600,
			zoom: 5,
			testLabels: []testLabel{
				{
					name:             "David Boyle",
					wkt:              "POINT(388600 413800)",
					expectedEnvelope: "POLYGON ((184.1150000000000091 278.0000000000000000, 415.8849999999999909 278.0000000000000000, 415.8849999999999909 322.0000000000000000, 184.1150000000000091 322.0000000000000000, 184.1150000000000091 278.0000000000000000))",
				},
			},
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:      float64(8),
			fontStroke:    float64(1),
			shouldSplitFn: noSplit,
		},
	}

	for name, tt := range tests {
		m := image.NewRGBA(image.Rect(0, 0, tt.dim, tt.dim))
		draw.Draw(m, m.Bounds(), &image.Uniform{colours.Lightgrey}, image.Point{0, 0}, draw.Src)
		gc := draw2dimg.NewGraphicContext(m)
		gc.SetDPI(72)

		i := types.ImageLayer{
			Zoom: tt.zoom,
		}

		shapeType, err := types.GetShapeType(types.FEAT_A_ROAD)
		if err != nil {
			t.Fatal(err)
		}

		labelStyle, err := types.GetLabelStyle(shapeType.TextFormat)
		if err != nil {
			t.Fatal(err)
		}

		typeFace := makeimage.GetTypeFace(gc, &i, labelStyle.LabelBackgroundStyle)

		for _, tl := range tt.testLabels {
			g, err := gctx.NewGeomFromWKT(tl.wkt)
			if err != nil {
				t.Fatal(err)
			}

			l := types.MapLabel{
				ID:            "F31F418C-F7D5-4306-B1C0-B320C720774A",
				LayerType:     "road",
				LabelType:     1,
				DataType:      3,
				FeatCode:      types.FEAT_A_ROAD,
				Geometry:      g,
				Label:         tl.name,
				Rotation:      0,
				LineCol:       color.RGBA{R: 0x8c, G: 0x8c, B: 0x8b, A: 0xff},
				FillCol:       color.RGBA{R: 0xff, G: 0xbb, B: 0xd2, A: 0xff},
				Zoom:          tt.zoom,
				Radius:        25,
				Thickness:     12.5,
				TypeFace:      typeFace,
				ShouldSplitFn: tt.shouldSplitFn,
			}

			r := apitypes.TileRequest{
				Location:   tt.tileCenter,
				TileWidth:  float64(tt.dim),
				TileHeight: float64(tt.dim),
				Zoom:       l.Zoom,
				Quality:    100,
			}

			tileHeight := r.TileHeight
			tileWidth := r.TileWidth

			bounds, err := r.BoundsGeom()
			if err != nil {
				t.Fatal(err)
			}

			tileEnvelope, err := geom.ToEnvelope(bounds)
			if err != nil {
				t.Fatal(err)
			}

			scale := func(x, y float64) (float64, float64) {
				x = tileEnvelope.Px(x) * tileWidth
				y = tileHeight - (tileEnvelope.Py(y) * tileHeight)
				return x, y
			}

			// build
			err = l.Build(scale)
			if err != nil {
				t.Fatal(err)
			}

			l.SetCenter([]float64{
				300,
				300,
			})

			envelope := l.Dimensions.Envelope
			if err != nil {
				t.Fatal(err)
			}
			actual := envelope.ToWKT()

			if tl.expectedEnvelope != actual {
				t.Errorf("%v\nExpected %v\nActual %v", name, tl.expectedEnvelope, actual)
			}

			// draw the label background
			err = l.DrawBackground(gc)
			if err != nil {
				t.Fatal(err)
			}

			err = l.DrawText(gc)
			if err != nil {
				t.Fatal(err)
			}

			err = savePNG(fmt.Sprintf("test-output/maplabel_envelope_%v.png", strings.ToLower(name)), m)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
