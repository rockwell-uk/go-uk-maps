package makeimage_test

import (
	"testing"

	"go-uk-maps/makeimage"
)

func TestIsOvelapping(t *testing.T) {
	type testLabel struct {
		name string
		wkt  string
	}

	tests := map[string]struct {
		labelA   testLabel
		labelB   testLabel
		expected bool
	}{
		"Overlap": {
			labelA: testLabel{
				name: "Combe Pafford",
				wkt:  "POLYGON ((141.9689999999972088 209.5529999999998552, 169.2269999999972185 209.5529999999998552, 169.2269999999972185 191.5529999999998552, 141.9689999999972088 191.5529999999998552, 141.9689999999972088 209.5529999999998552))",
			},
			labelB: testLabel{
				name: "Lummaton Hill",
				wkt:  "POLYGON ((122.6019999999972185 209.9694999999997833, 160.4969999999972288 209.9694999999997833, 160.4969999999972288 191.9694999999997833, 122.6019999999972185 191.9694999999997833, 122.6019999999972185 209.9694999999997833))",
			},
			expected: true,
		},
		"No Overlap": {
			labelA: testLabel{
				name: "Lummaton Hill",
				wkt:  "POLYGON ((122.6019999999972185 209.9694999999997833, 160.4969999999972288 209.9694999999997833, 160.4969999999972288 191.9694999999997833, 122.6019999999972185 191.9694999999997833, 122.6019999999972185 209.9694999999997833))",
			},
			labelB: testLabel{
				name: "St Marychurch",
				wkt:  "POLYGON ((191.5869999999972038 283.5619999999998413, 231.8469999999972231 283.5619999999998413, 231.8469999999972231 259.5619999999998413, 191.5869999999972038 259.5619999999998413, 191.5869999999972038 283.5619999999998413))",
			},
			expected: false,
		},
	}

	for name, tt := range tests {
		ag, err := gctx.NewGeomFromWKT(tt.labelA.wkt)
		if err != nil {
			t.Fatal(err)
		}

		bg, err := gctx.NewGeomFromWKT(tt.labelB.wkt)
		if err != nil {
			t.Fatal(err)
		}

		actual := makeimage.IsOvelapping(ag, bg)

		if tt.expected != actual {
			t.Errorf("%v: Expected %v, Actual %v", name, tt.expected, actual)
		}
	}
}
