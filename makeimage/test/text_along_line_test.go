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
	"github.com/rockwell-uk/go-text/fonts"
	"github.com/rockwell-uk/go-text/text"

	"go-uk-maps/api/types"
	"go-uk-maps/colours"
)

func TestGetLineData(t *testing.T) {
	tests := map[string]struct {
		wkt        string
		tileCenter nationalgrid.Location
		expected   string
		dim        int
		zoom       float64
	}{
		"mellor_street": {
			"MULTILINESTRING((388874 413258.9999997683,388844.99999999994 413290.99999976775,388740.99999999994 413427.9999997701,388659.00000000006 413499.99999976833,388648 413512.9999997685,388583.00000000006 413820.9999997686))",
			nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388874,
					Lon: 413459,
				},
			},
			"Mellor Street",
			600,
			1.2,
		},
	}

	lineWidth := 1.0

	for name, tt := range tests {
		m := image.NewRGBA(image.Rect(0, 0, tt.dim, tt.dim))
		draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
		gc := draw2dimg.NewGraphicContext(m)
		gc.SetDPI(72)

		r := types.TileRequest{
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

		noscale := func(x, y float64) (float64, float64) {
			return x, y
		}

		// wkt from the original geometry
		g, err := gctx.NewGeomFromWKT(tt.wkt)
		if err != nil {
			t.Fatal(err)
		}

		// this line gets overwritten by the blue line
		err = geom.DrawLine(gc, g, lineWidth, colours.Red, 1.0, colours.Black, scale)
		if err != nil {
			t.Fatal(err)
		}

		// scale the original geometry
		tl, err := geom.ScaleLine(g, scale)
		if err != nil {
			t.Fatal(err)
		}

		// transformed line
		err = geom.DrawLine(gc, tl, lineWidth, colours.Blue, 1.0, colours.Black, noscale)
		if err != nil {
			t.Fatal(err)
		}

		scaled := geom.TransposeMultiLineData(text.GetLineData(*geom.GetPoints(tl)))

		// set the origin to the same point as the blue and green line start from
		origin := []float64{
			300,
			324,
		}
		sl, err := scaled.ToGeom(origin)
		if err != nil {
			t.Fatal(err)
		}

		err = geom.DrawLine(gc, sl, lineWidth, colours.Green, 1.0, colours.Black, noscale)
		if err != nil {
			t.Fatal(err)
		}

		// finally draw the image
		err = savePNG(fmt.Sprintf("test-output/linedata/%v.png", name), m)
		if err != nil {
			t.Fatal(err)
		}

		// all 3 lines overlap, so we only see the green line
	}
}

func TestDrawRoadLines(t *testing.T) {
	tests := map[string]struct {
		dim          int
		fontData     draw2d.FontData
		fontSize     float64
		fontSpacing  float64
		fontStroke   float64
		fontColor    color.RGBA
		text         string
		geomWKT      string
		tileCenter   nationalgrid.Location
		showOutlines bool
	}{
		"MellorStreet": {
			dim: 600,
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:    8,
			fontSpacing: 0.5,
			fontStroke:  1,
			fontColor:   colours.Pink,
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388700,
					Lon: 413500,
				},
			},
			text:         "Mellor Street",
			geomWKT:      "MULTILINESTRING((388874 413258.9999997683,388844.99999999994 413290.99999976775,388740.99999999994 413427.9999997701,388659.00000000006 413499.99999976833,388648 413512.9999997685,388583.00000000006 413820.9999997686))",
			showOutlines: true,
		},
		"TurfHillRoad": {
			dim: 600,
			fontData: draw2d.FontData{
				Name:   "regular",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleNormal,
			},
			fontSize:    2,
			fontSpacing: 0.5,
			fontStroke:  0.2,
			fontColor:   colours.Darkgrey,
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 390902,
					Lon: 411492,
				},
			},
			text:         "Turf Hill Road",
			geomWKT:      "MULTILINESTRING((390902 411492.9999997673,390951.00000000006 411523.99999976787,391010 411571.999999767,391052 411608.9999997665,391092.99999999994 411656.99999976583))",
			showOutlines: false,
		},
		"RochValleyWay": {
			dim: 600,
			fontData: draw2d.FontData{
				Name:   "bold",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleBold,
			},
			fontSize:    6,
			fontSpacing: 0.5,
			fontStroke:  1,
			fontColor:   colours.Teal,
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 388279,
					Lon: 412467,
				},
			},
			text:         "Roch Valley Way",
			geomWKT:      "MULTILINESTRING((388379.00000000006 412267.9999997671,388332 412288.99999976665,388267.00000000006 412331.9999997683,388218 412381.99999976764,388100.99999999994 412517.9999997671,388058 412585.99999976833,388043.00000000006 412627.999999768,388034 412683.99999976694,388035.00000000006 412742.9999997681))",
			showOutlines: true,
		},
	}

	for name, tt := range tests {
		for zoom := 1; zoom <= 10; zoom++ {
			r := types.TileRequest{
				Location:   tt.tileCenter,
				TileWidth:  float64(tt.dim),
				TileHeight: float64(tt.dim),
				Zoom:       float64(zoom),
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

			polyGeom, err := gctx.NewGeomFromWKT(tt.geomWKT)
			if err != nil {
				t.Fatal(err)
			}

			lineString, err := geom.ToLineString(polyGeom)
			if err != nil {
				t.Fatal(err)
			}

			// draw the line
			err = geom.DrawLine(gc, lineString, 1.0, colours.Blue, 1.0, colours.Black, scale)
			if err != nil {
				t.Fatal(err)
			}

			fontSize := tt.fontSize * float64(zoom)
			fontStroke := tt.fontStroke * float64(zoom)
			fontSpacing := tt.fontSpacing * float64(zoom)
			strokeStyle := draw2d.StrokeStyle{
				Color: colours.White,
				Width: fontStroke,
			}

			face := fonts.GetFace(gc, tt.fontData, fontSize)

			typeFace := fonts.TypeFace{
				StrokeStyle: strokeStyle,
				Color:       colours.Pink,
				Size:        fontSize,
				FontData:    tt.fontData,
				Face:        fonts.GetFace(gc, tt.fontData, fontSize),
				Spacing:     fontSpacing,
			}
			fonts.SetFont(gc, typeFace)

			scaledGeom, err := geom.ScaleLine(polyGeom, scale)
			if err != nil {
				t.Fatal(err)
			}

			if tt.showOutlines {
				// draw an outline of the bounds for each rune
				err = text.DrawGlyphOutlines(gc, tt.text, *geom.GetPoints(scaledGeom), typeFace)
				if err != nil {
					t.Fatal(err)
				}
			}

			// text along line
			glyphs, err := text.TextAlongLine(gc, tt.text, *geom.GetPoints(scaledGeom), typeFace)
			if err != nil {
				t.Fatal(err)
			}
			for _, glyph := range glyphs {
				err = geom.DrawRune(gc, glyph.Pos, face, glyph.Rotation, glyph.Char)
				if err != nil {
					t.Fatal(err)
				}
			}

			err = savePNG(fmt.Sprintf("test-output/roadline/%v_%v.png", strings.ToLower(name), zoom), m)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

func TestTurfHillRoadRochdale(t *testing.T) {
	tests := map[string]struct {
		dim          int
		fontData     draw2d.FontData
		fontSize     float64
		fontSpacing  float64
		fontStroke   float64
		fontColor    color.RGBA
		zoom         float64
		text         string
		geomWKT      string
		tileCenter   nationalgrid.Location
		showOutlines bool
	}{
		"TurfHillRoad_Zoom_1": {
			dim: 600,
			fontData: draw2d.FontData{
				Name:   "regular",
				Family: draw2d.FontFamilySans,
				Style:  draw2d.FontStyleNormal,
			},
			zoom:        1,
			fontSize:    34,
			fontSpacing: 0.0,
			fontStroke:  1.0,
			fontColor:   colours.Darkgrey,
			tileCenter: nationalgrid.Location{
				Type: "OSGB36",
				LatLon: nationalgrid.LatLon{
					Lat: 390902,
					Lon: 411492,
				},
			},
			text:         "Turf Hill Road",
			geomWKT:      "MULTILINESTRING((390902 411492.9999997673,390951.00000000006 411523.99999976787,391010 411571.999999767,391052 411608.9999997665,391092.99999999994 411656.99999976583))",
			showOutlines: false,
		},
	}

	for name, tt := range tests {
		r := types.TileRequest{
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

		polyGeom, err := gctx.NewGeomFromWKT(tt.geomWKT)
		if err != nil {
			t.Fatal(err)
		}

		lineString, err := geom.ToLineString(polyGeom)
		if err != nil {
			t.Fatal(err)
		}

		// draw the line
		err = geom.DrawLine(gc, lineString, 1.0, colours.Blue, 1.0, colours.Black, scale)
		if err != nil {
			t.Fatal(err)
		}

		fontSize := tt.fontSize * tt.zoom
		fontStroke := tt.fontStroke * tt.zoom
		fontSpacing := tt.fontSpacing * tt.zoom
		strokeStyle := draw2d.StrokeStyle{
			Color: colours.White,
			Width: fontStroke,
		}
		face := fonts.GetFace(gc, tt.fontData, fontSize)
		scaledGeom, err := geom.ScaleLine(polyGeom, scale)
		if err != nil {
			t.Fatal(err)
		}

		typeFace := fonts.TypeFace{
			Name:        "regular",
			Spacing:     fontSpacing,
			StrokeStyle: strokeStyle,
			Color:       tt.fontColor,
			Size:        fontSize,
			FontData:    tt.fontData,
			Face:        face,
		}

		if tt.showOutlines {
			// draw an outline of the bounds for each rune
			err = text.DrawGlyphOutlines(gc, tt.text, *geom.GetPoints(scaledGeom), typeFace)
			if err != nil {
				t.Fatal(err)
			}
		}

		// text along line
		glyphs, _ := text.TextAlongLine(gc, tt.text, *geom.GetPoints(scaledGeom), typeFace)
		for _, glyph := range glyphs {
			err = geom.DrawRune(gc, glyph.Pos, face, glyph.Rotation, glyph.Char)
			if err != nil {
				t.Fatal(err)
			}
		}

		err = savePNG(fmt.Sprintf("test-output/rochdale-labels/%v.png", strings.ToLower(name)), m)
		if err != nil {
			t.Fatal(err)
		}
	}
}
