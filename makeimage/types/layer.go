package types

import (
	"image/color"

	"github.com/twpayne/go-geos"
)

type LabelType int

const (
	LABEL_POINT LabelType = iota
	LABEL_NAME
	LABEL_NUMBER
	LABEL_ICON
)

func (d LabelType) String() string {
	return [...]string{"Point", "Name", "Number", "Icon"}[d]
}

type AssetsAdded map[FeatCode]map[LabelType]map[string]ImageAsset

type LayerArtifact struct {
	ID                      string
	LayerType               string
	FeatCode                FeatCode
	GeomType                geos.TypeID
	Geometry                *geos.Geom
	ShapeType               ShapeType
	RenderType              RenderType
	LabelStyle              LabelStyle
	LabelFieldNames         []string
	LabelText               map[string]string
	CartographicFontDetails CartographicFontDetails
}

type CartographicFontDetails struct {
	FontHeight  string
	Orientation float64
}

type MapLayerGeometries struct {
	Points   map[DataType][]MapGeom
	Lines    map[DataType][]MapGeom
	Polygons map[DataType][]MapGeom
	Multi    map[DataType][]MapGeom
}

type MapLayerLabels struct {
	LabelsOrder map[FeatCode][]string
	Labels      map[FeatCode]map[string]MapLabel

	PointsOrder  map[FeatCode][]string
	Points       map[FeatCode]map[string]MapLabel
	NamesOrder   map[FeatCode][]string
	Names        map[FeatCode]map[string]MapLabel
	NumbersOrder map[FeatCode][]string
	Numbers      map[FeatCode]map[string]MapLabel
}

type ImageLayer struct {
	LayerType  string
	Zoom       float64
	Geometries MapLayerGeometries
	Labels     MapLayerLabels
}

type ImageAsset struct {
	ID        string
	LayerType string
	LabelType LabelType
	DataType  DataType
	FeatCode  FeatCode
	Geometry  *geos.Geom
	Label     string
}

type MapGeom struct {
	ID          string
	LayerType   string
	Geometry    *geos.Geom
	DataType    DataType
	FeatCode    FeatCode
	LineCol     color.RGBA
	FillCol     color.RGBA
	Radius      float64
	Thickness   float64
	StrokeWidth float64
}
