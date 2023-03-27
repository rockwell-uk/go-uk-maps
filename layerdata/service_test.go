package layerdata

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"path"
	"testing"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-nationalgrid"
	"github.com/twpayne/go-geos"

	"go-uk-maps/api/types"
	"go-uk-maps/colours"
)

func TestIsInViewPolygonsMix(t *testing.T) {
	tests := map[string]struct {
		bigTileWKTs  []string
		tilePolyWKTs []string
	}{
		"polygon": {
			bigTileWKTs: []string{
				"POLYGON((380000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 410000.0000000000000000))",
				"POLYGON((300000 400000, 400000 400000, 400000 500000, 300000 500000, 300000 400000))",
			},
			tilePolyWKTs: []string{
				"POLYGON((380000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 410000.0000000000000000))",
				"POLYGON((390000.0000000000000000 410000.0000000000000000, 400000.0000000000000000 410000.0000000000000000, 400000.0000000000000000 420000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 390000.0000000000000000 410000.0000000000000000))",
			},
		},
		"mix": {
			bigTileWKTs: []string{
				"POLYGON ((380000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 410000.0000000000000000))",
				"POLYGON ((300000 400000, 400000 400000, 400000 500000, 300000 500000, 300000 400000))",
			},
			tilePolyWKTs: []string{
				"LINESTRING(380000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 410000.0000000000000000)",
				"LINESTRING(390000.0000000000000000 410000.0000000000000000, 400000.0000000000000000 410000.0000000000000000, 400000.0000000000000000 420000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 390000.0000000000000000 410000.0000000000000000)",
			},
		},
	}

	var targetTilePoly string = "POLYGON((387221.1985319799860008 410715.0784210899728350, 392221.1985319799860008 410715.0784210899728350, 392221.1985319799860008 415715.0784210899728350, 387221.1985319799860008 415715.0784210899728350, 387221.1985319799860008 410715.0784210899728350))"
	targetTileGeom, _ := gctx.NewGeomFromWKT(targetTilePoly)

	for tname, tt := range tests {
		for key, bigTileWKT := range tt.bigTileWKTs {
			bigTileGeom, _ := gctx.NewGeomFromWKT(bigTileWKT)
			name := fmt.Sprintf("%v-bigtile-%v", tname, key)

			err := drawImage(name, bigTileGeom, targetTileGeom)
			if err != nil {
				t.Fatal(err)
			}

			inView := isInView(bigTileGeom.Bounds(), targetTileGeom.Bounds())
			if !inView {
				t.Fatalf("%v expected bigTileWKT in view %v", tname, bigTileWKT)
			}
		}

		for key, tilePolyWKT := range tt.tilePolyWKTs {
			tilePolyGeom, _ := gctx.NewGeomFromWKT(tilePolyWKT)
			name := fmt.Sprintf("%v-tilepoly-%v", tname, key)

			err := drawImage(name, targetTileGeom, tilePolyGeom)
			if err != nil {
				t.Fatal(err)
			}

			inView := isInView(targetTileGeom.Bounds(), tilePolyGeom.Bounds())
			if !inView {
				t.Fatalf("%v expected tilePolyWKT in view %v", tname, tilePolyWKT)
			}
		}
	}
}

func TestIsInViewLinestrings(t *testing.T) {
	// test to prove that overlapping linestrings now works
	targetTilePoly := "LINESTRING(387221.1985319799860008 410715.0784210899728350, 392221.1985319799860008 410715.0784210899728350, 392221.1985319799860008 415715.0784210899728350, 387221.1985319799860008 415715.0784210899728350, 387221.1985319799860008 410715.0784210899728350)"
	targetTileGeom, _ := gctx.NewGeomFromWKT(targetTilePoly)

	tests := map[string]struct {
		bigTileWKT   string
		expectInView bool
	}{
		"large": {
			"LINESTRING(380000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 410000.0000000000000000)",
			true,
		},
		"small": {
			"LINESTRING(300000 400000, 400000 400000, 400000 500000, 300000 500000, 300000 400000)",
			true,
		},
	}

	for name, tt := range tests {
		bigTileGeom, _ := gctx.NewGeomFromWKT(tt.bigTileWKT)
		name := fmt.Sprintf("linestring-bigtile-%v", name)

		err := drawImage(name, bigTileGeom, targetTileGeom)
		if err != nil {
			t.Fatal(err)
		}

		inView := isInView(bigTileGeom.Bounds(), targetTileGeom.Bounds())
		if tt.expectInView && !inView {
			t.Fatalf("%v expected bigTileWKT in view %v", name, tt.bigTileWKT)
		}
		if !tt.expectInView && inView {
			t.Fatalf("%v expected bigTileWKT not in view %v", name, tt.bigTileWKT)
		}
	}

	addnTests := map[string]string{
		"a": "LINESTRING(380000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 410000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 420000.0000000000000000, 380000.0000000000000000 410000.0000000000000000)",
		"b": "LINESTRING(390000.0000000000000000 410000.0000000000000000, 400000.0000000000000000 410000.0000000000000000, 400000.0000000000000000 420000.0000000000000000, 390000.0000000000000000 420000.0000000000000000, 390000.0000000000000000 410000.0000000000000000)",
	}

	for name, tilePolyWKT := range addnTests {
		tilePolyGeom, _ := gctx.NewGeomFromWKT(tilePolyWKT)
		name := fmt.Sprintf("linestring-tilepoly-%v", name)

		err := drawImage(name, targetTileGeom, tilePolyGeom)
		if err != nil {
			t.Fatal(err)
		}

		inView := isInView(targetTileGeom.Bounds(), tilePolyGeom.Bounds())
		if !inView {
			t.Fatalf("%v expected tilePolyWKT in view %v", name, tilePolyWKT)
		}
	}
}

func drawImage(name string, a, b *geos.Geom) error {
	dim := 1000
	tileCenter := nationalgrid.Location{
		Type: "OSGB36",
		LatLon: nationalgrid.LatLon{
			Lat: 388874,
			Lon: 413459,
		},
	}
	zoom := .12

	m := image.NewRGBA(image.Rect(0, 0, dim, dim))
	draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
	gc := draw2dimg.NewGraphicContext(m)
	gc.SetDPI(72)

	r := types.TileRequest{
		Location:   tileCenter,
		TileWidth:  float64(dim),
		TileHeight: float64(dim),
		Zoom:       zoom,
		Quality:    100,
	}

	bounds, err := r.BoundsGeom()
	if err != nil {
		return err
	}

	envelope, err := geom.ToEnvelope(bounds)
	if err != nil {
		return err
	}

	tileHeight := r.TileHeight
	tileWidth := r.TileWidth

	scale := func(x, y float64) (float64, float64) {
		x = envelope.Px(x) * tileWidth
		y = tileHeight - (envelope.Py(y) * tileHeight)
		return x, y
	}

	fillColor := colours.White
	strokeColor := colours.Black
	strokeWidth := 1.0
	thickness := 1.0

	_type := b.TypeID()

	if _type == geos.TypeIDLineString {
		err = geom.DrawLine(gc, a, thickness, fillColor, strokeWidth, strokeColor, scale)
		if err != nil {
			return err
		}
		err = geom.DrawLine(gc, b, thickness, fillColor, strokeWidth, strokeColor, scale)
		if err != nil {
			return err
		}
	}
	if _type == geos.TypeIDPolygon {
		err = geom.DrawPolygon(gc, a, fillColor, strokeColor, strokeWidth, scale)
		if err != nil {
			return err
		}
		err = geom.DrawPolygon(gc, b, fillColor, strokeColor, strokeWidth, scale)
		if err != nil {
			return err
		}
	}

	// finally draw the image
	err = savePNG(fmt.Sprintf("test-output/%v.png", name), m)
	if err != nil {
		return err
	}

	return nil
}

func savePNG(fname string, m image.Image) error {
	dir, _ := path.Split(fname)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	return draw2dimg.SaveToPngFile(fname, m)
}
