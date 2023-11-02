package api

import (
	"fmt"
	"image"
	"os"
	"path"
	"reflect"
	"testing"

	"go-uk-maps/api/types"
	"go-uk-maps/database"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-nationalgrid"
)

func TestParseURL(t *testing.T) {
	tests := map[string]struct {
		path        string
		tileRequest types.TileRequest
	}{
		"basic 2645": {
			path: "/256/13/4046/2645.png",
			tileRequest: types.TileRequest{
				Location: nationalgrid.Location{
					Type: "OSGB36",
					EastNorth: nationalgrid.EastNorth{
						X: 389938.5218501827,
						Y: 410687.6916311788,
					},
				},
				TileWidth:  289,
				TileHeight: 289,
				Zoom:       1,
			},
		},
		"basic 2646": {
			path: "/256/13/4046/2646.png",
			tileRequest: types.TileRequest{
				Location: nationalgrid.Location{
					Type: "OSGB36",
					EastNorth: nationalgrid.EastNorth{
						X: 389930.47144396976,
						Y: 407786.73510286136,
					},
				},
				TileWidth:  289,
				TileHeight: 289,
				Zoom:       1,
			},
		},
		"zoomed": {
			path: "/256/14/8096/5290.png",
			tileRequest: types.TileRequest{
				Location: nationalgrid.Location{
					Type: "OSGB36",
					EastNorth: nationalgrid.EastNorth{
						X: 191188.14654721774,
						Y: 441204.99190356,
					},
				},
				TileWidth:  289,
				TileHeight: 289,
				Zoom:       11,
			},
		},
	}

	for tname, tt := range tests {
		tr, err := parseURL(tt.path)
		if err != nil {
			t.Fatalf("%v: %v %v", tname, tr, err)
		}
		//tileCenter := tr.Location.ToOSGB36()
		//t.Fatalf("%v", tileCenter)

		if !reflect.DeepEqual(tt.tileRequest, tr) {
			t.Fatalf("%v:\nexpected %+v\ngot %+v", tname, tt.tileRequest, tr)
		}
	}
}

func TestOSGB36ToWGS84(t *testing.T) {
	tests := map[string]struct {
		tileRequest types.TileRequest
		east        float64
		north       float64
	}{
		"Rochdale 1": {
			tileRequest: types.TileRequest{
				Location: nationalgrid.Location{
					Type: "OSGB36",
					EastNorth: nationalgrid.EastNorth{
						X: 389721,
						Y: 413215,
					},
				},
				TileWidth:  600,
				TileHeight: 600,
				Zoom:       1,
			},
			east:  -2.1568529911185794, // -2.1561
			north: 53.61537906042884,   // 53.6097136
		},
	}

	for tname, tt := range tests {
		east, north, _ := oSGB36ToWGS84(
			tt.tileRequest.Location.EastNorth.X,
			tt.tileRequest.Location.EastNorth.Y,
			0.0)

		if tt.east != east {
			t.Fatalf("%v: east - expected %v got %v", tname, tt.east, east)
		}
		if tt.north != north {
			t.Fatalf("%v: north - expected %v got %v", tname, tt.north, north)
		}
	}
}

func TestOSMTile(t *testing.T) {

	connectDB()

	tests := map[string]struct {
		path string
	}{
		"rw_13_4046_2245": {
			path: "/256/13/4046/2645.png",
		},
		"rw_13_4046_2246": {
			path: "/256/13/4046/2646.png",
		},
		"rw_14_8096_5290": {
			path: "/256/14/8096/5290.png",
		},
	}

	for tname, tt := range tests {
		tileRequest, err := parseURL(tt.path)
		if err != nil {
			t.Fatalf("%v: %v %v", tname, tileRequest, err)
		}

		tile, err := makeTile(db, tileRequest)
		if err != nil {
			t.Fatalf("%v: %v %v", tname, tileRequest, err)
			return
		}

		/*
			tileSize := 256
			m := image.NewRGBA(image.Rect(0, 0, tileSize, tileSize))
			draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
			gc := draw2dimg.NewGraphicContext(m)
			gc.SetDPI(72)
		*/

		// Return the output filename
		err = savePNG(fmt.Sprintf("test-output/%s.png", tname), *tile)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func connectDB() {
	// Database
	dbConfig := database.Config{
		Engine:        "mysql",
		Host:          "127.0.0.1",
		Port:          "3307",
		User:          "osdata",
		Pass:          "osdata",
		Schema:        "osdata",
		StorageFolder: "db",
		Timeout:       10,
	}
	db = database.Start(dbConfig)
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
