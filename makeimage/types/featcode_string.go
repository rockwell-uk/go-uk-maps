// Code generated by "stringer -type=FeatCode"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FEAT_BUILDING-25014]
	_ = x[FEAT_GLASSHOUSE-25016]
	_ = x[FEAT_MOTORWAY-25710]
	_ = x[FEAT_PRIMARY_ROAD-25723]
	_ = x[FEAT_A_ROAD-25729]
	_ = x[FEAT_B_ROAD-25743]
	_ = x[FEAT_MINOR_ROAD-25750]
	_ = x[FEAT_LOCAL_STREET-25760]
	_ = x[FEAT_PRIVATE_ROAD-25780]
	_ = x[FEAT_PEDESTRIAN_STREET-25790]
	_ = x[FEAT_MOTORWAY_COLLAPSED_DUAL_CARRIAGEWAY-25719]
	_ = x[FEAT_PRIMARY_ROAD_COLLAPSED_DUAL_CARRIAGEWAY-25735]
	_ = x[FEAT_A_ROAD_COLLAPSED_DUAL_CARRIAGEWAY-25739]
	_ = x[FEAT_B_ROAD_COLLAPSED_DUAL_CARRIAGEWAY-25749]
	_ = x[FEAT_MINOR_ROAD_COLLAPSED_DUAL_CARRIAGEWAY-25759]
	_ = x[FEAT_ROAD_TUNNEL-25792]
	_ = x[FEAT_MOTORWAY_JN-25796]
	_ = x[FEAT_ROUNDABOUT_PRIMARY_ROAD-25703]
	_ = x[FEAT_ROUNDABOUT_A_ROAD-25704]
	_ = x[FEAT_ROUNDABOUT_B_ROAD-25705]
	_ = x[FEAT_ROUNDABOUT_MINOR_ROAD-25706]
	_ = x[FEAT_ROUNDABOUT_LOCAL_STREET-25707]
	_ = x[FEAT_ROUNDABOUT_PRIVATE_ROAD-25708]
	_ = x[FEAT_SURFACE_WATER_LINE-25600]
	_ = x[FEAT_SURFACE_WATER_AREA-25609]
	_ = x[FEAT_HIGH_WATER_MARK-25608]
	_ = x[FEAT_HIGH_WATER_MARK_LOW_WATER_MARK-25604]
	_ = x[FEAT_LOW_WATER_MARK-25605]
	_ = x[FEAT_FORESHORE-25612]
	_ = x[FEAT_NATIONAL_BOUNDARY-25204]
	_ = x[FEAT_PARISH_OR_COMMUNITY_BOUNDARY-25200]
	_ = x[FEAT_DISTRICT_OR_LONDON_BOROUGH_BOUNDARY-25201]
	_ = x[FEAT_COUNTY_REGION_OR_ISLAND_BOUNDARY-25202]
	_ = x[FEAT_MULTI_TRACK_RAILWAY-25300]
	_ = x[FEAT_SINGLE_TRACK_RAILWAY-25301]
	_ = x[FEAT_NARROW_GAGUE_RAILWAY-25302]
	_ = x[FEAT_RAIL_TUNNEL_ALLIGNMENT-25303]
	_ = x[FEAT_LIGHT_RAPID_TRANSIT_STATION-25420]
	_ = x[FEAT_RAILWAY_STATION-25422]
	_ = x[FEAT_LONDON_UNDERGROUND_STATION-25423]
	_ = x[FEAT_RAILWAY_STATION_AND_LONDON_UNDERGROUND_STATION-25424]
	_ = x[FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_RAILWAY_STATION-25425]
	_ = x[FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_LONDON_UNDERGROUND_STATION-25426]
	_ = x[FEAT_EDUCATION_FACILITY_SCHOOL-25250]
	_ = x[FEAT_POLICE_STATION-25251]
	_ = x[FEAT_MEDICAL_CARE-25252]
	_ = x[FEAT_PLACE_OF_WORSHIP-25253]
	_ = x[FEAT_LEISURE_OR_SPORTS_CENTRE-25254]
	_ = x[FEAT_AIR_TRANSPORT-25255]
	_ = x[FEAT_EDUCATION_FACILITY_HIGHER-25256]
	_ = x[FEAT_WATER_TRANSPORT-25257]
	_ = x[FEAT_ROAD_TRANSPORT-25258]
	_ = x[FEAT_ROAD_SERVICES-25259]
	_ = x[FEAT_WOODLAND-25999]
	_ = x[FEAT_CUSTOM_LANDFORM-25550]
	_ = x[FEAT_ELECTRICITY_TRANSMISSION_LINE-25102]
	_ = x[FEAT_POPULATED_PLACE-25801]
	_ = x[FEAT_LANDFORM-25802]
	_ = x[FEAT_WOODLAND_OR_FOREST-25803]
	_ = x[FEAT_HYDROGRAPHY-25804]
	_ = x[FEAT_LANCOVER-25805]
	_ = x[FEAT_HEIGHTED_POINT-25810]
}

const _FeatCode_name = "FEAT_BUILDINGFEAT_GLASSHOUSEFEAT_ELECTRICITY_TRANSMISSION_LINEFEAT_PARISH_OR_COMMUNITY_BOUNDARYFEAT_DISTRICT_OR_LONDON_BOROUGH_BOUNDARYFEAT_COUNTY_REGION_OR_ISLAND_BOUNDARYFEAT_NATIONAL_BOUNDARYFEAT_EDUCATION_FACILITY_SCHOOLFEAT_POLICE_STATIONFEAT_MEDICAL_CAREFEAT_PLACE_OF_WORSHIPFEAT_LEISURE_OR_SPORTS_CENTREFEAT_AIR_TRANSPORTFEAT_EDUCATION_FACILITY_HIGHERFEAT_WATER_TRANSPORTFEAT_ROAD_TRANSPORTFEAT_ROAD_SERVICESFEAT_MULTI_TRACK_RAILWAYFEAT_SINGLE_TRACK_RAILWAYFEAT_NARROW_GAGUE_RAILWAYFEAT_RAIL_TUNNEL_ALLIGNMENTFEAT_LIGHT_RAPID_TRANSIT_STATIONFEAT_RAILWAY_STATIONFEAT_LONDON_UNDERGROUND_STATIONFEAT_RAILWAY_STATION_AND_LONDON_UNDERGROUND_STATIONFEAT_LIGHT_RAPID_TRANSIT_STATION_AND_RAILWAY_STATIONFEAT_LIGHT_RAPID_TRANSIT_STATION_AND_LONDON_UNDERGROUND_STATIONFEAT_CUSTOM_LANDFORMFEAT_SURFACE_WATER_LINEFEAT_HIGH_WATER_MARK_LOW_WATER_MARKFEAT_LOW_WATER_MARKFEAT_HIGH_WATER_MARKFEAT_SURFACE_WATER_AREAFEAT_FORESHOREFEAT_ROUNDABOUT_PRIMARY_ROADFEAT_ROUNDABOUT_A_ROADFEAT_ROUNDABOUT_B_ROADFEAT_ROUNDABOUT_MINOR_ROADFEAT_ROUNDABOUT_LOCAL_STREETFEAT_ROUNDABOUT_PRIVATE_ROADFEAT_MOTORWAYFEAT_MOTORWAY_COLLAPSED_DUAL_CARRIAGEWAYFEAT_PRIMARY_ROADFEAT_A_ROADFEAT_PRIMARY_ROAD_COLLAPSED_DUAL_CARRIAGEWAYFEAT_A_ROAD_COLLAPSED_DUAL_CARRIAGEWAYFEAT_B_ROADFEAT_B_ROAD_COLLAPSED_DUAL_CARRIAGEWAYFEAT_MINOR_ROADFEAT_MINOR_ROAD_COLLAPSED_DUAL_CARRIAGEWAYFEAT_LOCAL_STREETFEAT_PRIVATE_ROADFEAT_PEDESTRIAN_STREETFEAT_ROAD_TUNNELFEAT_MOTORWAY_JNFEAT_POPULATED_PLACEFEAT_LANDFORMFEAT_WOODLAND_OR_FORESTFEAT_HYDROGRAPHYFEAT_LANCOVERFEAT_HEIGHTED_POINTFEAT_WOODLAND"

var _FeatCode_map = map[FeatCode]string{
	25014: _FeatCode_name[0:13],
	25016: _FeatCode_name[13:28],
	25102: _FeatCode_name[28:62],
	25200: _FeatCode_name[62:95],
	25201: _FeatCode_name[95:135],
	25202: _FeatCode_name[135:172],
	25204: _FeatCode_name[172:194],
	25250: _FeatCode_name[194:224],
	25251: _FeatCode_name[224:243],
	25252: _FeatCode_name[243:260],
	25253: _FeatCode_name[260:281],
	25254: _FeatCode_name[281:310],
	25255: _FeatCode_name[310:328],
	25256: _FeatCode_name[328:358],
	25257: _FeatCode_name[358:378],
	25258: _FeatCode_name[378:397],
	25259: _FeatCode_name[397:415],
	25300: _FeatCode_name[415:439],
	25301: _FeatCode_name[439:464],
	25302: _FeatCode_name[464:489],
	25303: _FeatCode_name[489:516],
	25420: _FeatCode_name[516:548],
	25422: _FeatCode_name[548:568],
	25423: _FeatCode_name[568:599],
	25424: _FeatCode_name[599:650],
	25425: _FeatCode_name[650:702],
	25426: _FeatCode_name[702:765],
	25550: _FeatCode_name[765:785],
	25600: _FeatCode_name[785:808],
	25604: _FeatCode_name[808:843],
	25605: _FeatCode_name[843:862],
	25608: _FeatCode_name[862:882],
	25609: _FeatCode_name[882:905],
	25612: _FeatCode_name[905:919],
	25703: _FeatCode_name[919:947],
	25704: _FeatCode_name[947:969],
	25705: _FeatCode_name[969:991],
	25706: _FeatCode_name[991:1017],
	25707: _FeatCode_name[1017:1045],
	25708: _FeatCode_name[1045:1073],
	25710: _FeatCode_name[1073:1086],
	25719: _FeatCode_name[1086:1126],
	25723: _FeatCode_name[1126:1143],
	25729: _FeatCode_name[1143:1154],
	25735: _FeatCode_name[1154:1198],
	25739: _FeatCode_name[1198:1236],
	25743: _FeatCode_name[1236:1247],
	25749: _FeatCode_name[1247:1285],
	25750: _FeatCode_name[1285:1300],
	25759: _FeatCode_name[1300:1342],
	25760: _FeatCode_name[1342:1359],
	25780: _FeatCode_name[1359:1376],
	25790: _FeatCode_name[1376:1398],
	25792: _FeatCode_name[1398:1414],
	25796: _FeatCode_name[1414:1430],
	25801: _FeatCode_name[1430:1450],
	25802: _FeatCode_name[1450:1463],
	25803: _FeatCode_name[1463:1486],
	25804: _FeatCode_name[1486:1502],
	25805: _FeatCode_name[1502:1515],
	25810: _FeatCode_name[1515:1534],
	25999: _FeatCode_name[1534:1547],
}

func (i FeatCode) String() string {
	if str, ok := _FeatCode_map[i]; ok {
		return str
	}
	return "FeatCode(" + strconv.FormatInt(int64(i), 10) + ")"
}
