package makeimage

import (
	"fmt"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/twpayne/go-geos"
	"golang.org/x/image/draw"

	"go-uk-maps/layerdata"
	"go-uk-maps/makeimage/types"
)

func prepareLayerArtifacts(gc *draw2dimg.GraphicContext, m draw.Image, l *types.ImageLayer, d []layerdata.LayerData, scale func(x, y float64) (float64, float64)) (types.AssetsAdded, error) {
	assetsAdded := make(types.AssetsAdded)

	for _, ld := range d {
		a, err := artifactFromLayerData(l, ld)
		if err != nil {
			return types.AssetsAdded{}, err
		}

		if !containsFeat(featCodeOrder, a.FeatCode) {
			featCodeOrder = append(featCodeOrder, a.FeatCode)
		}

		assetsAdded, err = prepareGeom(m, l, assetsAdded, a, scale)
		if err != nil {
			return types.AssetsAdded{}, err
		}

		err = prepareLabel(gc, l, a, scale)
		if err != nil {
			return types.AssetsAdded{}, err
		}
	}

	return assetsAdded, nil
}

func artifactFromLayerData(l *types.ImageLayer, ld layerdata.LayerData) (types.LayerArtifact, error) {
	var geomType geos.TypeID

	geom, err := gctx.NewGeomFromWKT(ld.WKT.String)
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("%v [%+v]", err.Error(), ld)
	}

	geomType = geom.TypeID()

	featCode, err := types.GetFeatCode(ld.FEATCODE)
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("[%v] %v: %v", l.LayerType, err.Error(), ld)
	}

	shapeType, err := types.GetShapeType(featCode)
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("[%v] %v: %v", l.LayerType, err.Error(), ld)
	}

	renderType, err := types.GetRenderType(shapeType.RenderType)
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("[%v] %v: %v", l.LayerType, err.Error(), ld)
	}

	labelStyle, err := types.GetLabelStyle(shapeType.TextFormat)
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("[%v] %v: %v", l.LayerType, err.Error(), ld)
	}

	labelFieldNames, err := types.GetLabelFieldNames(shapeType.DataType)
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("[%v] %v: %v", l.LayerType, err.Error(), ld)
	}

	labelText := map[string]string{}
	for _, fieldname := range labelFieldNames {
		labelString, err := getStringField(ld, fieldname)
		if err != nil {
			return types.LayerArtifact{}, fmt.Errorf("%v [%+v]", err.Error(), ld)
		}
		labelText[fieldname] = fmt.Sprintf("%s ", labelString)
	}

	cartoFontHeight, err := getStringField(ld, "FONTHEIGHT")
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("%v [%+v]", err.Error(), ld)
	}
	cartoOrientation, err := getFloatField(ld, "ORIENTATIO")
	if err != nil {
		return types.LayerArtifact{}, fmt.Errorf("%v [%+v]", err.Error(), ld)
	}

	artifact := types.LayerArtifact{
		ID:              ld.ID,
		LayerType:       l.LayerType,
		FeatCode:        types.FeatCode(ld.FEATCODE),
		GeomType:        geomType,
		Geometry:        geom,
		ShapeType:       shapeType,
		RenderType:      renderType,
		LabelStyle:      labelStyle,
		LabelFieldNames: labelFieldNames,
		LabelText:       labelText,
		CartographicFontDetails: types.CartographicFontDetails{
			FontHeight:  cartoFontHeight,
			Orientation: cartoOrientation,
		},
	}

	return artifact, nil
}
