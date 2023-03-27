package types

import (
	"fmt"
	"image/color"

	"go-uk-maps/colours"
)

type ShapeType struct {
	LineCol    color.RGBA
	FillCol    color.RGBA
	DataType   DataType
	RenderType string
	TextFormat TextFormat
}

// SHAPETYPES array["linecol", "fillcol", "datatype", "rendertype", "textformat"].
var ShapeTypes = map[FeatCode]ShapeType{
	// Building.shp
	FEAT_BUILDING: {colours.Darkpeach, colours.Peach, LINE, "building", LABEL}, // Building

	// Glasshouse.shp
	FEAT_GLASSHOUSE: {colours.Lightgrey, colours.Lightgrey, LINE, "glasshouse", LABEL}, // Glasshouse

	// Road.shp
	FEAT_MOTORWAY:                            {colours.Midgrey, colours.Lightblue, LINE, "motorway", MOTORWAY},                // Motorway : Line
	FEAT_PRIMARY_ROAD:                        {colours.Midgrey, colours.Lightteal, LINE, "primary_road", PRIMARY_ROAD},        // Primary Road : Line
	FEAT_A_ROAD:                              {colours.Midgrey, colours.Lightpink, LINE, "a_road", A_ROAD},                    // A Road : Line
	FEAT_B_ROAD:                              {colours.Midgrey, colours.Lightorange, LINE, "b_road", B_ROAD},                  // B Road : Line
	FEAT_MINOR_ROAD:                          {colours.Midgrey, colours.Lightyellow, LINE, "minor_road", MINOR_ROAD},          // Minor Road : Line
	FEAT_LOCAL_STREET:                        {colours.Midgrey, colours.Vlightgrey, LINE, "local_street", LOCAL_STREET},       // Local Street : Line
	FEAT_PRIVATE_ROAD:                        {colours.Midgrey, colours.Vlightgrey, LINE, "priv_road", PRIVATE_ROAD},          // Private Road, Public Access : Line
	FEAT_PEDESTRIAN_STREET:                   {colours.Midgrey, colours.Vlightgrey, LINE, "pedest_street", PEDESTRIAN_STREET}, // Pedestrianised Street : Line
	FEAT_MOTORWAY_COLLAPSED_DUAL_CARRIAGEWAY: {colours.Midgrey, colours.Skyblue, LINE, "motorway_dual_carriageway", MOTORWAY}, // Motorway, Collapsed Dual Carriageway : Line
	FEAT_PRIMARY_ROAD_COLLAPSED_DUAL_CARRIAGEWAY: {colours.Midgrey, colours.Lightteal, LINE, "primary_road_dual_carriageway", PRIMARY_ROAD}, // Primary Road, Collapsed Dual Carriageway : Line
	FEAT_A_ROAD_COLLAPSED_DUAL_CARRIAGEWAY:       {colours.Midgrey, colours.Lightpink, LINE, "a_road_dual_carriageway", A_ROAD},             // A Road, Collapsed Dual Carriageway : Line
	FEAT_B_ROAD_COLLAPSED_DUAL_CARRIAGEWAY:       {colours.Midgrey, colours.Lightorange, LINE, "b_road_dual_carriageway", B_ROAD},           // B Road, Collapsed Dual Carriageway : Line
	FEAT_MINOR_ROAD_COLLAPSED_DUAL_CARRIAGEWAY:   {colours.Midgrey, colours.Lightyellow, LINE, "minor_road_dual_carriageway", MINOR_ROAD},   // Minor Road, Collapsed Dual Carriageway : Line

	// RoadTunnel.shp
	FEAT_ROAD_TUNNEL: {colours.Midgrey, colours.Lightgrey, LINE, "road_tunnel", ROAD_TUNNEL}, // Road Tunnel : Dashed Line

	// MotorwayJunction.shp
	FEAT_MOTORWAY_JN: {colours.Midgrey, colours.Skyblue, JUNCTION, "motorway_junction", MOTORWAY_JN}, // MotorwayJunction : Text Label

	// Roundabout.shp
	FEAT_ROUNDABOUT_PRIMARY_ROAD: {colours.Midgrey, colours.Lightteal, POINT, "roundabout_primary_road", PRIMARY_ROAD},  // Primary Road : Line
	FEAT_ROUNDABOUT_A_ROAD:       {colours.Midgrey, colours.Lightpink, POINT, "roundabout_a_road", A_ROAD},              // A Road : Line
	FEAT_ROUNDABOUT_B_ROAD:       {colours.Midgrey, colours.Lightorange, POINT, "roundabout_b_road", B_ROAD},            // B Road : Line
	FEAT_ROUNDABOUT_MINOR_ROAD:   {colours.Midgrey, colours.Lightyellow, POINT, "roundabout_minor_road", MINOR_ROAD},    // Minor Road : Line
	FEAT_ROUNDABOUT_LOCAL_STREET: {colours.Midgrey, colours.Vlightgrey, POINT, "roundabout_local_street", LOCAL_STREET}, // Local Street : Line
	FEAT_ROUNDABOUT_PRIVATE_ROAD: {colours.Midgrey, colours.Vlightgrey, POINT, "roundabout_priv_road", PRIVATE_ROAD},    // Private Road, Public Access : Line

	// SurfaceWater_Line.shp
	FEAT_SURFACE_WATER_LINE: {colours.Vlightblue, colours.Lightblue, NO_LABEL, "water_line", LABEL}, // Water : Line

	// SurfaceWater_Area.shp
	FEAT_SURFACE_WATER_AREA: {colours.Seablue, colours.Seablue, NO_LABEL, "water", LABEL}, // Water : Polygon

	// TidalWater.shp
	FEAT_HIGH_WATER_MARK: {colours.Seablue, colours.Seablue, NO_LABEL, "high_water_mark", LABEL}, // High Water Mark : Line

	// TidalBoundary.shp
	FEAT_HIGH_WATER_MARK_LOW_WATER_MARK: {colours.Vlightblue, colours.Lightblue, NO_LABEL, "high_water_mark", LABEL}, // High Water Mark Low Water Mark : Line
	FEAT_LOW_WATER_MARK:                 {colours.Vlightblue, colours.Lightblue, NO_LABEL, "low_water_mark", LABEL},  // Low Water Mark : Line

	// Foreshore.shp
	FEAT_FORESHORE: {colours.Darkbeachgrey, colours.Beachgrey, LINE, "foreshore", LABEL}, // Foreshore : Line

	// AdministrativeBoundary.shp
	FEAT_NATIONAL_BOUNDARY:                   {colours.Lightgrey, colours.Lightgrey, LINE, "boundary", LABEL}, // National Boundary
	FEAT_PARISH_OR_COMMUNITY_BOUNDARY:        {colours.Lightgrey, colours.Lightgrey, LINE, "boundary", LABEL}, // Parish or Community Boundary
	FEAT_DISTRICT_OR_LONDON_BOROUGH_BOUNDARY: {colours.Lightgrey, colours.Lightgrey, LINE, "boundary", LABEL}, // District or London Borough Boundary
	FEAT_COUNTY_REGION_OR_ISLAND_BOUNDARY:    {colours.Lightgrey, colours.Lightgrey, LINE, "boundary", LABEL}, // County, Region Or Island Boundary

	// RailwayTrack.shp
	FEAT_MULTI_TRACK_RAILWAY:  {colours.White, colours.Darkgrey, NO_LABEL, "rline_multi", LABEL},  // Multi Track Railway : Line
	FEAT_SINGLE_TRACK_RAILWAY: {colours.White, colours.Darkgrey, NO_LABEL, "rline_single", LABEL}, // Single Track Railway : Line
	FEAT_NARROW_GAGUE_RAILWAY: {colours.White, colours.Darkgrey, NO_LABEL, "rline_narrow", LABEL}, // Narrow Gauge Railway : Line

	// RailwayTunnel.shp
	FEAT_RAIL_TUNNEL_ALLIGNMENT: {colours.White, colours.Darkgrey, NO_LABEL, "rline_tunnel", LABEL}, // Rail Tunnel Alignment : Line

	// RailwayStation.shp
	FEAT_LIGHT_RAPID_TRANSIT_STATION:                                {colours.Darkgrey, colours.Lightyellow, LINE, "lrt_station", LABEL},   // Light Rapid Transit Station
	FEAT_RAILWAY_STATION:                                            {colours.Darkgrey, colours.Lightpink, LINE, "railway_station", LABEL}, // Railway Station
	FEAT_LONDON_UNDERGROUND_STATION:                                 {colours.Darkgrey, colours.Lightgrey, LINE, "lu_station", LABEL},      //  London Underground Station
	FEAT_RAILWAY_STATION_AND_LONDON_UNDERGROUND_STATION:             {colours.Darkgrey, colours.Lightpink, LINE, "railway_station", LABEL}, // Railway Station And London Underground Station
	FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_RAILWAY_STATION:            {colours.Darkgrey, colours.Lightpink, LINE, "railway_station", LABEL}, // Light Rapid Transit Station And Railway Station
	FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_LONDON_UNDERGROUND_STATION: {colours.Darkgrey, colours.Lightgrey, LINE, "lu_station", LABEL},      // Light Rapid Transit Station And London Underground Station

	// FunctionalSite.shp
	FEAT_EDUCATION_FACILITY_SCHOOL: {colours.Lightgrey, colours.Lightgrey, EDUCATION_SCHOOL, "point", LABEL}, // Education Facility - School : Point
	FEAT_POLICE_STATION:            {colours.Lightgrey, colours.Lightgrey, POLICE_STATION, "point", LABEL},   // Police Station : Point
	FEAT_MEDICAL_CARE:              {colours.Lightgrey, colours.Lightgrey, HOSPITAL, "point", LABEL},         // Medical Care : Point
	FEAT_PLACE_OF_WORSHIP:          {colours.Lightgrey, colours.Lightgrey, PLACE_OF_WORSHIP, "point", LABEL}, // Place Of Worship : Point
	FEAT_LEISURE_OR_SPORTS_CENTRE:  {colours.Lightgrey, colours.Lightgrey, LEISURE_CENTRE, "point", LABEL},   // eisure or Sports Centre : Point
	FEAT_AIR_TRANSPORT:             {colours.Lightgrey, colours.Lightgrey, AIRPORT, "point", LABEL},          // Air Transport : Point
	FEAT_EDUCATION_FACILITY_HIGHER: {colours.Lightgrey, colours.Lightgrey, EDUCATION_HIGHER, "point", LABEL}, // Education Facility - Higher : Point
	FEAT_WATER_TRANSPORT:           {colours.Lightgrey, colours.Lightgrey, WATER_TRANSPORT, "point", LABEL},  // Water Transport 25257
	FEAT_ROAD_TRANSPORT:            {colours.Lightgrey, colours.Lightgrey, ROAD_TRANSPORT, "point", LABEL},   // Road Transport 25258
	FEAT_ROAD_SERVICES:             {colours.Lightgrey, colours.Lightgrey, ROAD_SERVICES, "point", LABEL},    // Road Services 25259

	// Woodland.shp
	FEAT_WOODLAND: {colours.Darkgreen, colours.Lightgreen, NO_LABEL, "woodland", LABEL}, // Woodland : Polygon

	// Ornament.shp
	FEAT_CUSTOM_LANDFORM: {colours.Grey, colours.Grey, NO_LABEL, "landform", LABEL}, // Custom Landform : Polygon

	// ElectricityTransmissionLine.shp
	FEAT_ELECTRICITY_TRANSMISSION_LINE: {colours.Lightgrey, colours.Lightgrey, LINE, "electricity_transmission_line", LABEL}, // ElectricityTransmissionLine : Line

	// NamedPlace.shp
	FEAT_POPULATED_PLACE:    {colours.Lightgrey, colours.Lightgrey, CARTOGRAPHIC, "text", LABEL},       // Populated Place : Point
	FEAT_LANDFORM:           {colours.Lightgrey, colours.Lightgrey, CARTOGRAPHIC, "text", LABEL},       // Landform : Point
	FEAT_WOODLAND_OR_FOREST: {colours.Lightgrey, colours.Lightgrey, CARTOGRAPHIC, "text", LABEL},       // Woodland Or Forest : Point
	FEAT_HYDROGRAPHY:        {colours.Lightgrey, colours.Lightgrey, CARTOGRAPHIC, "text", HYDROGRAPHY}, // Hydrography : Point
	FEAT_LANCOVER:           {colours.Lightgrey, colours.Lightgrey, CARTOGRAPHIC, "text", LABEL},       // Landcover : Point

	// SpotHeight.shp
	//nolint:misspell
	FEAT_HEIGHTED_POINT: {colours.Darkgrey, colours.Darkgrey, POINT, "height", HEIGHTPOINT}, // Heighted Point : Point
}

func GetShapeType(f FeatCode) (ShapeType, error) {
	if _, ok := ShapeTypes[f]; !ok {
		return ShapeTypes[f], fmt.Errorf("ShapeType not found for FeatCode: %v", f)
	}

	return ShapeTypes[f], nil
}
