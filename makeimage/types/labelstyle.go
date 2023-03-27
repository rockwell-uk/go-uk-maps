package types

//go:generate stringer -type=TextFormat

import (
	"fmt"

	"github.com/llgcode/draw2d"

	"go-uk-maps/colours"
)

type LabelStyle struct {
	LabelTextStyle       TextStyle
	LabelBackgroundStyle TextStyle
}

var LabelStyles = map[TextFormat]LabelStyle{
	HEIGHTPOINT: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightgrey,
			4,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"regular",
			colours.Darkgrey,
			4,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	LABEL: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightgrey,
			6,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"regular",
			colours.Darkgrey,
			6,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	MOTORWAY: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkblue,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.Blue,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	MOTORWAY_JN: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkblue,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.Blue,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	PRIMARY_ROAD: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkteal,
			9,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.Teal,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	A_ROAD: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkpink,
			9,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.Pink,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	B_ROAD: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkorange,
			9,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.Orange,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	MINOR_ROAD: {
		LabelTextStyle: TextStyle{
			"regular",
			colours.Lightgrey,
			7,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.5},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.Lightyellow,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	LOCAL_STREET: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightgrey,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	PRIVATE_ROAD: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightgrey,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	PEDESTRIAN_STREET: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightgrey,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	HYDROGRAPHY: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightblue,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	ROAD_TUNNEL: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Lightgrey,
			6,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},

	CT_SMALL: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkgrey,
			5.5,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	CT_MEDIUM: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkgrey,
			6,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	CT_LARGE: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkgrey,
			12,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			8,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
	CT_EXTRA_LARGE: {
		LabelTextStyle: TextStyle{
			"bold",
			colours.Darkgrey,
			16,
			draw2d.StrokeStyle{Color: colours.White, Width: 1.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
		},
		LabelBackgroundStyle: TextStyle{
			"bold",
			colours.White,
			10,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.0},
			colours.White,
			draw2d.StrokeStyle{Color: colours.White, Width: 0.2},
		},
	},
}

func GetLabelStyle(t TextFormat) (LabelStyle, error) {
	if _, ok := LabelStyles[t]; !ok {
		return LabelStyles[t], fmt.Errorf("LabelStyle not found for TextFormat: %v", t)
	}

	return LabelStyles[t], nil
}
