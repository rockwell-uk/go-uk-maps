package types

import (
	"go-uk-maps/icons"
)

var CommunityServices = map[FeatCode][]byte{
	FEAT_EDUCATION_FACILITY_SCHOOL: icons.EducationIcon, // School
	FEAT_POLICE_STATION:            icons.PoliceIcon,
	FEAT_MEDICAL_CARE:              icons.HospitalIcon,
	FEAT_PLACE_OF_WORSHIP:          icons.LeisureIcon, // Place Of Worship
	FEAT_LEISURE_OR_SPORTS_CENTRE:  icons.LeisureIcon,
	FEAT_AIR_TRANSPORT:             icons.AirportIcon,
	FEAT_EDUCATION_FACILITY_HIGHER: icons.EducationIcon, // Higher
	FEAT_WATER_TRANSPORT:           icons.HospitalIcon,  // Water Transport
	FEAT_ROAD_TRANSPORT:            icons.HospitalIcon,  // Road Transport
	FEAT_ROAD_SERVICES:             icons.HospitalIcon,  // Road Services
}
