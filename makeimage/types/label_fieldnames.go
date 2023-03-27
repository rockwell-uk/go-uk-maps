package types

import (
	"fmt"
)

// LABELFIELDNAMES.
var LabelFieldNames = map[DataType][]string{
	NO_LABEL:         {},
	SITE:             {"DISTNAME"},
	POINT:            {"HEIGHT"},
	LINE:             {"DISTNAME", "ROADNUMBER"},
	JUNCTION:         {"JUNCTNUM"},
	CARTOGRAPHIC:     {"DISTNAME"},
	EDUCATION_SCHOOL: {"DISTNAME"},
	POLICE_STATION:   {"DISTNAME"},
	HOSPITAL:         {"DISTNAME"},
	LEISURE_CENTRE:   {"DISTNAME"},
	AIRPORT:          {"DISTNAME"},
	PLACE_OF_WORSHIP: {"DISTNAME"},
	EDUCATION_HIGHER: {"DISTNAME"},
	WATER_TRANSPORT:  {"DISTNAME"},
	ROAD_TRANSPORT:   {"DISTNAME"},
	ROAD_SERVICES:    {"DISTNAME"},
}

func GetLabelFieldNames(d DataType) ([]string, error) {
	if _, ok := LabelFieldNames[d]; !ok {
		return LabelFieldNames[d], fmt.Errorf("LabelFieldNames not found for DataType: %v", d)
	}

	return LabelFieldNames[d], nil
}
