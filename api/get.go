package api

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"go-uk-maps/api/types"

	"github.com/rockwell-uk/go-nationalgrid"
	"github.com/wroge/wgs84"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	tileRequest, err := parseURL(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	tile, err := makeTile(db, tileRequest)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeTile(tile, tileRequest, w, r)
}

func parseURL(path string) (types.TileRequest, error) {
	var zoom float64
	var multiplier float64 = 10

	p := strings.Split(path, "/")

	tileSize, err := strconv.Atoi(p[1])
	if err != nil {
		return types.TileRequest{}, err
	}

	z, err := strconv.Atoi(p[2])
	if err != nil {
		return types.TileRequest{}, err
	}

	x, err := strconv.Atoi(p[3])
	if err != nil {
		return types.TileRequest{}, err
	}

	lonDpr := p[4]
	bits := strings.Split(lonDpr, ".")
	yv := bits[0]
	if strings.Contains(bits[0], "@") {
		bits := strings.Split(lonDpr, "@")
		yv = bits[0]
	}

	y, err := strconv.Atoi(yv)
	if err != nil {
		return types.TileRequest{}, err
	}

	lon := tile2lng(float64(x), float64(z))
	lat := tile2lat(float64(y), float64(z))

	wgs84 := types.TileRequest{
		Location: nationalgrid.Location{
			Type: "WGS84",
			LonLat: nationalgrid.LonLat{
				Lon: lon,
				Lat: lat,
			},
		},
		TileWidth:  float64(tileSize),
		TileHeight: float64(tileSize),
		Zoom:       float64(zoom),
	}

	eastNorth := wgs84.Location.ToOSGB36()

	switch z {
	case 13:
		multiplier = 10.0
		tileSize = 289
		zoom = 1
	case 14:
		multiplier = 2
		tileSize = 289
		zoom = 2
	}

	osgb36 := types.TileRequest{
		Location: nationalgrid.Location{
			Type: "OSGB36",
			EastNorth: nationalgrid.EastNorth{
				X: eastNorth.X + float64(tileSize)*multiplier,
				Y: eastNorth.Y - float64(tileSize)*multiplier,
			},
		},
		TileWidth:  float64(tileSize),
		TileHeight: float64(tileSize),
		Zoom:       float64(zoom),
	}

	return osgb36, nil
}

func tile2lng(x, z float64) float64 {
	return x/math.Pow(2, z)*360 - 180
}

func tile2lat(y, z float64) float64 {
	var n = math.Pi - 2*math.Pi*y/math.Pow(2, z)
	return 180 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
}

func num2deg(z, x, y int) (lat float64, long float64) {
	n := math.Pi - 2.0*math.Pi*float64(y)/math.Exp2(float64(z))
	lat = 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	long = float64(x)/math.Exp2(float64(z))*360.0 - 180.0
	return lat, long
}

// returns east, north, h
func wGS84ToOSGB36(lon, lat, h float64) (float64, float64, float64) {
	return wgs84.LonLat().To(wgs84.OSGB36NationalGrid())(lon, lat, h)
}

// returns lat, lon, h
func oSGB36ToWGS84(east, north, h float64) (float64, float64, float64) {
	return wgs84.OSGB36NationalGrid().To(wgs84.LonLat())(east, north, h)
}
