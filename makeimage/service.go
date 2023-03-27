package makeimage

import (
	"fmt"
	"image"
	"image/draw"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-logger/logger"
	"github.com/twpayne/go-geos"

	apitypes "go-uk-maps/api/types"
	"go-uk-maps/colours"
	"go-uk-maps/layerdata"
	"go-uk-maps/makeimage/types"
)

var (
	gctx = geos.NewContext()

	featCodeOrder []types.FeatCode
)

func DrawImage(r apitypes.TileRequest, layerOrder []string, layerData map[string][]layerdata.LayerData) (*image.Image, error) {
	var tileWidth float64 = r.TileWidth
	var tileHeight float64 = r.TileHeight
	var tileImage image.Image

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("TileRequest %+v\n", r),
	)

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("layerOrder %+v\n", layerOrder),
	)

	zoom := r.Zoom
	if zoom == float64(0) {
		zoom = float64(1)
	}

	bounds, err := r.BoundsGeom()
	if err != nil {
		return &tileImage, err
	}

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("bounds %+v\n", bounds),
	)

	envelope, err := geom.ToEnvelope(bounds)
	if err != nil {
		return &tileImage, err
	}

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("envelope %+v\n", envelope),
	)

	scale := func(x, y float64) (float64, float64) {
		x = envelope.Px(x) * tileWidth
		y = tileHeight - (envelope.Py(y) * tileHeight)
		return x, y
	}

	m := image.NewRGBA(image.Rect(0, 0, int(r.TileWidth), int(r.TileHeight)))
	draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
	gc := draw2dimg.NewGraphicContext(m)
	gc.SetDPI(72)

	for _, layerName := range layerOrder {
		err := drawLayer(r, m, gc, zoom, layerName, layerData[layerName], scale)
		if err != nil {
			return &tileImage, err
		}
	}
	tileImage = m

	return &tileImage, nil
}

func GetLabelText(labelTexts map[string]string, key string) (string, error) {
	if label, ok := labelTexts[key]; ok {
		return strings.TrimSpace(label), nil
	}

	return "", fmt.Errorf("label field not found [%v]", key)
}

func getStringField(l layerdata.LayerData, fieldname string) (string, error) {
	switch fieldname {
	case "DISTNAME":
		return l.DISTNAME.String, nil
	case "HEIGHT":
		return l.HEIGHT.String, nil
	case "ROADNUMBER":
		return l.ROADNUMBER.String, nil
	case "FONTHEIGHT":
		return l.FONTHEIGHT.String, nil
	case "JUNCTNUM":
		return l.JUNCTNUM.String, nil
	}
	return "", fmt.Errorf("string field not found [%v]", fieldname)
}

func getFloatField(l layerdata.LayerData, fieldname string) (float64, error) {
	if fieldname == "ORIENTATIO" {
		return l.ORIENTATIO, nil
	}

	return float64(0), fmt.Errorf("float64 field not found [%v]", fieldname)
}

func containsFeat(s []types.FeatCode, e types.FeatCode) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
