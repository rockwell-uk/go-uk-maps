package makeimage

import (
	"fmt"

	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-logger/logger"
	"github.com/rockwell-uk/go-text/text"
	"github.com/twpayne/go-geos"

	apitypes "go-uk-maps/api/types"
	"go-uk-maps/makeimage/types"
)

const (
	roadNumberDispersionPixels = 70
	roadNumberCirclePoints     = 10
)

func shouldBeAdded(r apitypes.TileRequest, assetsAdded types.AssetsAdded, l types.MapLabel) bool {
	// if its a communityservices label, allow
	if _, ok := types.CommunityServices[l.FeatCode]; ok {
		return true
	}

	var g *geos.Geom
	var envelope *geos.Geom = l.Dimensions.Envelope

	switch l.LabelType {
	case types.LABEL_ICON:
		return true

	case types.LABEL_NUMBER:

		// if fully visible
		if !IsFullyVisible(envelope, r.TileGeom) {
			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("%v Not Fully Visible [%+v]\n", l.LabelType.String(), l.Label),
			)
			return false
		}

		// has already been added
		for _, featcode := range types.RoadTypes {
			for _, assets := range assetsAdded[featcode] {
				for _, existing := range assets {
					if existing.Label == l.Label {
						logger.Log(
							logger.LVL_INTERNAL,
							fmt.Sprintf("%v Already Added [%+v]\n", l.LabelType.String(), l.Label),
						)
						return false
					}
				}
			}
		}

		if l.FeatCode == types.FEAT_MOTORWAY_JN {
			return true
		} else {
			// expanded envelope for road numbers
			g = GetCircularEnvelope(*l.Dimensions.Center, roadNumberDispersionPixels)

			for _, featcode := range types.RoadTypes {
				for _, assets := range assetsAdded[featcode] {
					for _, existing := range assets {
						if l.LabelType == existing.LabelType && IsOvelapping(g, existing.Geometry) {
							logger.Log(
								logger.LVL_INTERNAL,
								fmt.Sprintf("%v Too Close [%+v] [%+v] [%+v] [%+v]\n", l.LabelType.String(), l.Label, l.LabelType, existing.Label, existing.LabelType),
							)
							return false
						}
					}
				}
			}
		}

	case types.LABEL_NAME:

		// if fully visible
		if !IsFullyVisible(envelope, r.TileGeom) {
			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("%v Not Fully Visible [%+v]\n", l.LabelType.String(), l.Label),
			)
			return false
		}

		// has already been added
		if _, exists := assetsAdded[l.FeatCode][l.LabelType][l.Label]; exists {
			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("%v Already Added [%+v]\n", l.LabelType.String(), l.Label),
			)
			return false
		}

		// check if the label will be written
		_, err := text.GetLetterPositions(l.Label, *geom.GetPoints(l.Geometry), l.TypeFace)
		if err != nil {
			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("Name Doesnt Fit [%+v]\n", l.Label),
			)
			return false
		}

		g = envelope

	case types.LABEL_POINT:

		// has already been added
		if _, exists := assetsAdded[l.FeatCode][l.LabelType][l.Label]; exists {
			logger.Log(
				logger.LVL_INTERNAL,
				fmt.Sprintf("%v Already Added [%+v]\n", l.LabelType.String(), l.Label),
			)
			return false
		}

		g = envelope

		if l.FeatCode == types.FEAT_HEIGHTED_POINT { //nolint:misspell
			// is the label too close to an existing label
			for _, featcode := range featCodeOrder {
				for _, assets := range assetsAdded[featcode] {
					for _, existing := range assets {
						if IsOvelapping(envelope, existing.Geometry) {
							logger.Log(
								logger.LVL_INTERNAL,
								fmt.Sprintf("%v Too Close [%+v] (%v) [%+v] (%v)\n", l.LabelType.String(), l.Label, l.FeatCode, existing.Label, existing.FeatCode),
							)
							return false
						}
					}
				}
			}
		} else {
			// is the label too close to an existing label of the same name
			for _, assets := range assetsAdded[l.FeatCode] {
				for _, existing := range assets {
					if IsOvelapping(envelope, existing.Geometry) {
						logger.Log(
							logger.LVL_INTERNAL,
							fmt.Sprintf("%v Too Close [%+v] [%+v]\n", l.LabelType.String(), l.Label, existing.Label),
						)
						return false
					}
				}
			}
		}

	default:
		logger.Log(
			logger.LVL_INTERNAL,
			fmt.Sprintf("Unknown LabelType [%+v]]\n", l.LabelType),
		)
	}

	// cleared for takeoff
	addToMap(assetsAdded, types.ImageAsset{
		ID:        l.ID,
		LayerType: l.LayerType,
		LabelType: l.LabelType,
		DataType:  l.DataType,
		FeatCode:  l.FeatCode,
		Geometry:  g,
		Label:     l.Label,
	})

	logger.Log(
		logger.LVL_INTERNAL,
		fmt.Sprintf("%v Added [%+v]\n", l.LabelType.String(), l.Label),
	)

	return true
}

// do the geometries overlap?
func IsOvelapping(item *geos.Geom, target *geos.Geom) bool {
	intersects := item.Intersects(target)
	return intersects
}

// is the polyWkt within the target tile poly?
func IsFullyVisible(item *geos.Geom, target *geos.Geom) bool {
	within := item.Within(target)
	return within
}

func GetCircularEnvelope(pos []float64, expansion float64) *geos.Geom {
	g, err := geom.CircleGeom(
		pos,
		expansion,
		roadNumberCirclePoints,
	)
	if err != nil {
		logger.Log(
			logger.LVL_FATAL,
			err.Error(),
		)
	}

	return g
}

func addToMap(assetMap map[types.FeatCode]map[types.LabelType]map[string]types.ImageAsset, asset types.ImageAsset) {
	if _, exists := assetMap[asset.FeatCode]; !exists {
		assetMap[asset.FeatCode] = make(map[types.LabelType]map[string]types.ImageAsset)
	}

	if _, exists := assetMap[asset.FeatCode][asset.LabelType]; !exists {
		assetMap[asset.FeatCode][asset.LabelType] = make(map[string]types.ImageAsset)
	}

	assetMap[asset.FeatCode][asset.LabelType][asset.ID] = asset
}
