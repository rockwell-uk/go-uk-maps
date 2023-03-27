package types

import (
	"testing"

	"github.com/rockwell-uk/go-nationalgrid"
)

func TestTileRequestValidate(t *testing.T) {
	control := TileRequest{
		Location: nationalgrid.Location{
			LatLon: nationalgrid.LatLon{
				Lat: 1,
				Lon: 1,
			},
		},
	}

	if err := control.Validate(); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		request TileRequest
		errors  []string
	}{
		{
			TileRequest{},
			[]string{
				"tile location is not defined",
			},
		},
	}

	for _, tt := range tests {
		actual := tt.request.Validate()
		expected := ErrorMsg(tt.errors)

		if actual == nil || expected == nil || actual.Error() != expected.Error() {
			t.Errorf("Failed: expected [%v], got [%v]", expected, actual)
		}
	}
}

func TestTileRequestBoundsGeom(t *testing.T) {
	tests := map[string]struct {
		tileRequest TileRequest
		boundsGeom  string
	}{
		"OSGB36 1": {
			tileRequest: TileRequest{
				Location: nationalgrid.Location{
					Type: "OSGB36",
					LatLon: nationalgrid.LatLon{
						Lat: 388874,
						Lon: 413459,
					},
				},
				TileWidth:  float64(600),
				TileHeight: float64(600),
				Zoom:       1.2,
				Quality:    100,
			},
			boundsGeom: "POLYGON ((386374.0000000000000000 410959.0000000000000000, 391374.0000000000000000 410959.0000000000000000, 391374.0000000000000000 415959.0000000000000000, 386374.0000000000000000 415959.0000000000000000, 386374.0000000000000000 410959.0000000000000000))",
		},
		"OSGB36 2": {
			tileRequest: TileRequest{
				Location: nationalgrid.Location{
					Type: "OSGB36",
					LatLon: nationalgrid.LatLon{
						Lat: 380626,
						Lon: 413800,
					},
				},
				TileWidth:  float64(600),
				TileHeight: float64(600),
				Zoom:       100,
				Quality:    100,
			},
			boundsGeom: "POLYGON ((380596.0000000000000000 413770.0000000000000000, 380656.0000000000000000 413770.0000000000000000, 380656.0000000000000000 413830.0000000000000000, 380596.0000000000000000 413830.0000000000000000, 380596.0000000000000000 413770.0000000000000000))",
		},
	}

	for tname, tt := range tests {
		actual, err := tt.tileRequest.BoundsGeom()
		if err != nil {
			t.Fatalf("%v: %v %v", tname, actual, err)
		}

		if tt.boundsGeom != actual.String() {
			t.Fatalf("%v:\nexpected %+v\ngot %+v", tname, tt.boundsGeom, actual)
		}
	}
}
