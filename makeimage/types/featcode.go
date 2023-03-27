package types

//go:generate stringer -type=FeatCode

type FeatCode int

const (
	// Building.shp.
	FEAT_BUILDING FeatCode = 25014

	// Glasshouse.shp.
	FEAT_GLASSHOUSE FeatCode = 25016

	// Road.shp.
	FEAT_MOTORWAY                                FeatCode = 25710
	FEAT_PRIMARY_ROAD                            FeatCode = 25723
	FEAT_A_ROAD                                  FeatCode = 25729
	FEAT_B_ROAD                                  FeatCode = 25743
	FEAT_MINOR_ROAD                              FeatCode = 25750
	FEAT_LOCAL_STREET                            FeatCode = 25760
	FEAT_PRIVATE_ROAD                            FeatCode = 25780
	FEAT_PEDESTRIAN_STREET                       FeatCode = 25790
	FEAT_MOTORWAY_COLLAPSED_DUAL_CARRIAGEWAY     FeatCode = 25719
	FEAT_PRIMARY_ROAD_COLLAPSED_DUAL_CARRIAGEWAY FeatCode = 25735
	FEAT_A_ROAD_COLLAPSED_DUAL_CARRIAGEWAY       FeatCode = 25739
	FEAT_B_ROAD_COLLAPSED_DUAL_CARRIAGEWAY       FeatCode = 25749
	FEAT_MINOR_ROAD_COLLAPSED_DUAL_CARRIAGEWAY   FeatCode = 25759

	// RoadTunnel.shp.
	FEAT_ROAD_TUNNEL FeatCode = 25792

	// MotorwayJunction.shp.
	FEAT_MOTORWAY_JN FeatCode = 25796

	// Roundabout.shp.
	FEAT_ROUNDABOUT_PRIMARY_ROAD FeatCode = 25703
	FEAT_ROUNDABOUT_A_ROAD       FeatCode = 25704
	FEAT_ROUNDABOUT_B_ROAD       FeatCode = 25705
	FEAT_ROUNDABOUT_MINOR_ROAD   FeatCode = 25706
	FEAT_ROUNDABOUT_LOCAL_STREET FeatCode = 25707
	FEAT_ROUNDABOUT_PRIVATE_ROAD FeatCode = 25708

	// SurfaceWater_Line.shp.
	FEAT_SURFACE_WATER_LINE FeatCode = 25600

	// SurfaceWater_Area.shp.
	FEAT_SURFACE_WATER_AREA FeatCode = 25609

	// TidalWater.shp.
	FEAT_HIGH_WATER_MARK FeatCode = 25608

	// TidalBoundary.shp.
	FEAT_HIGH_WATER_MARK_LOW_WATER_MARK FeatCode = 25604
	FEAT_LOW_WATER_MARK                 FeatCode = 25605

	// Foreshore.shp.
	FEAT_FORESHORE FeatCode = 25612

	// AdministrativeBoundary.shp.
	FEAT_NATIONAL_BOUNDARY                   FeatCode = 25204
	FEAT_PARISH_OR_COMMUNITY_BOUNDARY        FeatCode = 25200
	FEAT_DISTRICT_OR_LONDON_BOROUGH_BOUNDARY FeatCode = 25201
	FEAT_COUNTY_REGION_OR_ISLAND_BOUNDARY    FeatCode = 25202

	// RailwayTrack.shp.
	FEAT_MULTI_TRACK_RAILWAY  FeatCode = 25300
	FEAT_SINGLE_TRACK_RAILWAY FeatCode = 25301
	FEAT_NARROW_GAGUE_RAILWAY FeatCode = 25302

	// RailwayTunnel.shp.
	FEAT_RAIL_TUNNEL_ALLIGNMENT FeatCode = 25303

	// RailwayStation.shp.
	FEAT_LIGHT_RAPID_TRANSIT_STATION                                FeatCode = 25420
	FEAT_RAILWAY_STATION                                            FeatCode = 25422
	FEAT_LONDON_UNDERGROUND_STATION                                 FeatCode = 25423
	FEAT_RAILWAY_STATION_AND_LONDON_UNDERGROUND_STATION             FeatCode = 25424
	FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_RAILWAY_STATION            FeatCode = 25425
	FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_LONDON_UNDERGROUND_STATION FeatCode = 25426

	// FunctionalSite.shp.
	FEAT_EDUCATION_FACILITY_SCHOOL FeatCode = 25250
	FEAT_POLICE_STATION            FeatCode = 25251
	FEAT_MEDICAL_CARE              FeatCode = 25252
	FEAT_PLACE_OF_WORSHIP          FeatCode = 25253
	FEAT_LEISURE_OR_SPORTS_CENTRE  FeatCode = 25254
	FEAT_AIR_TRANSPORT             FeatCode = 25255
	FEAT_EDUCATION_FACILITY_HIGHER FeatCode = 25256
	FEAT_WATER_TRANSPORT           FeatCode = 25257
	FEAT_ROAD_TRANSPORT            FeatCode = 25258
	FEAT_ROAD_SERVICES             FeatCode = 25259

	// Woodland.shp.
	FEAT_WOODLAND FeatCode = 25999

	// Ornament.shp.
	FEAT_CUSTOM_LANDFORM FeatCode = 25550

	// ElectricityTransmissionLine.shp.
	FEAT_ELECTRICITY_TRANSMISSION_LINE FeatCode = 25102

	// NamedPlace.shp.
	FEAT_POPULATED_PLACE    FeatCode = 25801
	FEAT_LANDFORM           FeatCode = 25802
	FEAT_WOODLAND_OR_FOREST FeatCode = 25803
	FEAT_HYDROGRAPHY        FeatCode = 25804
	FEAT_LANCOVER           FeatCode = 25805

	// SpotHeight.shp.
	FEAT_HEIGHTED_POINT FeatCode = 25810 //nolint:misspell
)

func GetFeatCode(f int) (FeatCode, error) {
	fc := FeatCode(f)

	return fc, nil
}
