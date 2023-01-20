package types

//go:generate stringer -type=TextFormat

import (
	"fmt"
	"image/color"

	"github.com/llgcode/draw2d"
)

type TextFormat int

const (
	HEIGHTPOINT TextFormat = iota
	LABEL
	MOTORWAY
	PRIMARY_ROAD
	A_ROAD
	B_ROAD
	MINOR_ROAD
	LOCAL_STREET
	PRIVATE_ROAD
	PEDESTRIAN_STREET
	ROAD_TUNNEL
	MOTORWAY_JN
	HYDROGRAPHY
	CT_SMALL
	CT_MEDIUM
	CT_LARGE
	CT_EXTRA_LARGE
)

type TextStyle struct {
	FontName string
	Color    color.RGBA
	Size     float64
	Stroke   draw2d.StrokeStyle
	BgCol    color.RGBA
	BgStroke draw2d.StrokeStyle
}

var TextSizes = map[string]LabelStyle{
	"Small":       LabelStyles[CT_SMALL],
	"Medium":      LabelStyles[CT_MEDIUM],
	"Large":       LabelStyles[CT_LARGE],
	"Extra Large": LabelStyles[CT_EXTRA_LARGE],
}

func GetLabelStyleByTextSize(textSize string) (LabelStyle, error) {
	if _, ok := TextSizes[textSize]; !ok {
		return LabelStyle{}, fmt.Errorf("unknown font size name [%v]", textSize)
	}

	return TextSizes[textSize], nil
}
