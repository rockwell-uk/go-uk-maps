package types

import (
	"fmt"
)

type RenderType struct {
	StrokeWidth   float64
	LineThickness float64
	DotRadius     float64
}

// RENDERTYPES array["StrokeWidth", "LineThickness", "DotRadius"].
var RenderTypes = map[string]RenderType{
	"building":   {0.1, 0.4, 5.0},
	"glasshouse": {0.1, 0.1, 5.0},
	"woodland":   {0.1, 0.4, 5.0},
	"water":      {0.1, 0.1, 5.0},
	"landform":   {0.1, 0.1, 5.0},
	"water_line": {0.0, 0.1, 5.0},
	"boundary":   {0.0, 0.1, 5.0},
	"point":      {0.0, 0.0, 5.0},
	"height":     {0.0, 0.1, 0.5},
	"label":      {0.0, 0.0, 5.0},

	"rline_multi":  {0.0, 2.0, 5.0},
	"rline_single": {0.0, 1.0, 5.0},
	"rline_narrow": {0.0, 0.5, 5.0},
	"rline_tunnel": {0.0, 2.0, 5.0},

	"railway_station": {0.1, 1.0, 5.0},
	"lrt_station":     {0.1, 1.0, 5.0},
	"lu_station":      {0.1, 1.0, 5.0},

	"motorway":      {0.1, 4.0, 5.0},
	"primary_road":  {0.5, 2.5, 5.0},
	"a_road":        {0.5, 2.5, 5.0},
	"b_road":        {0.5, 2.0, 5.0},
	"minor_road":    {0.1, 2.0, 5.0},
	"local_street":  {0.1, 2.0, 5.0},
	"priv_road":     {0.1, 2.0, 5.0},
	"pedest_street": {0.1, 2.0, 5.0},

	"roundabout_primary_road":  {0.1, 0.1, 4.0},
	"roundabout_a_road":        {0.1, 0.1, 4.0},
	"roundabout_b_road":        {0.1, 0.1, 4.0},
	"roundabout_minor_road":    {0.1, 0.1, 3.0},
	"roundabout_local_street":  {0.1, 0.1, 3.0},
	"roundabout_priv_road":     {0.1, 0.1, 3.0},
	"roundabout_pedest_street": {0.1, 0.1, 2.0},

	"motorway_dual_carriageway":     {0.5, 4.0, 5.0},
	"primary_road_dual_carriageway": {0.5, 2.5, 5.0},
	"a_road_dual_carriageway":       {0.5, 2.5, 5.0},
	"b_road_dual_carriageway":       {0.5, 2.5, 5.0},
	"minor_road_dual_carriageway":   {0.5, 2.5, 5.0},

	"road_tunnel":       {0.1, 2.0, 5.0},
	"motorway_junction": {0.0, 0.0, 5.0},

	"electricity_transmission_line": {0.0, 0.1, 5.0},

	"high_water_mark": {0.0, 1.0, 5.0},
	"low_water_mark":  {0.0, 1.0, 5.0},

	"foreshore": {0.0, 1.0, 5.0},

	"text": {0.0, 0.0, 0.0},
}

func GetRenderType(s string) (RenderType, error) {
	if _, ok := RenderTypes[s]; !ok {
		return RenderTypes[s], fmt.Errorf("RenderType not found for ShapeType: %v", s)
	}

	return RenderTypes[s], nil
}
