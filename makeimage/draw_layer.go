package makeimage

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/draw"

	apitypes "go-uk-maps/api/types"
	"go-uk-maps/layerdata"
	"go-uk-maps/makeimage/types"
)

func drawLayer(r apitypes.TileRequest, m draw.Image, gc *draw2dimg.GraphicContext, zoom float64, layerType string, layerData []layerdata.LayerData, scale func(x, y float64) (float64, float64)) error {
	imageLayer := types.ImageLayer{
		LayerType: layerType,
		Zoom:      zoom,
		Geometries: types.MapLayerGeometries{
			Points:   make(map[types.DataType][]types.MapGeom),
			Lines:    make(map[types.DataType][]types.MapGeom),
			Polygons: make(map[types.DataType][]types.MapGeom),
			Multi:    make(map[types.DataType][]types.MapGeom),
		},
		Labels: types.MapLayerLabels{
			LabelsOrder:  make(map[types.FeatCode][]string),
			Labels:       make(map[types.FeatCode]map[string]types.MapLabel),
			PointsOrder:  make(map[types.FeatCode][]string),
			Points:       make(map[types.FeatCode]map[string]types.MapLabel),
			NamesOrder:   make(map[types.FeatCode][]string),
			Names:        make(map[types.FeatCode]map[string]types.MapLabel),
			NumbersOrder: make(map[types.FeatCode][]string),
			Numbers:      make(map[types.FeatCode]map[string]types.MapLabel),
		},
	}

	assetsAdded, err := prepareLayerArtifacts(gc, m, &imageLayer, layerData, scale)
	if err != nil {
		return err
	}

	err = drawLayerGeometries(gc, imageLayer.Geometries, scale)
	if err != nil {
		return err
	}

	err = drawLayerLabels(r, assetsAdded, gc, imageLayer.Labels, scale)
	if err != nil {
		return err
	}

	return nil
}
