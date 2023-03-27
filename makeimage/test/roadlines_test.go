package makeimage_test

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-text/fonts"
	"github.com/twpayne/go-geos"

	"go-uk-maps/colours"
	"go-uk-maps/makeimage"
	"go-uk-maps/makeimage/types"
)

func TestLineTest(t *testing.T) {
	title := "Line Styles"

	width := 1080
	height := 1920

	var fontSize float64
	var fontCol color.RGBA
	fontName := "bold"

	lineHeight := 70.0
	border := 50.0
	zoomedLineWidthPercent := 60.0
	unZoomedLineWidthPercent := 40.0

	y := border
	zoomedLineZoom := 6.0
	zoomedLineLength := (float64(width) - (border * 1.5)) * zoomedLineWidthPercent / 100
	lineLength := (float64(width) - (border * 1.5)) * unZoomedLineWidthPercent / 100
	zoomedLineEndX := border + zoomedLineLength

	fontSize = 12.0

	// setup image
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{colours.White}, image.Point{0, 0}, draw.Src)
	gc := draw2dimg.NewGraphicContext(m)
	gc.SetDPI(72)

	shouldSplitFn := func(s string) bool {
		return false
	}

	noscale := func(x, y float64) (float64, float64) {
		return x, y
	}

	// strokestyle
	strokeStyle := draw2d.StrokeStyle{
		Color: colours.White,
		Width: 1.0,
	}

	// font
	fontData := draw2d.FontData{
		Name:   fontName,
		Family: draw2d.FontFamilySans,
		Style:  draw2d.FontStyleNormal,
	}

	typeFace := fonts.TypeFace{
		StrokeStyle: strokeStyle,
		Color:       colours.Black,
		Size:        fontSize,
		FontData:    fontData,
		Face:        fonts.GetFace(gc, fontData, fontSize),
	}

	// document title
	titlePos := []float64{
		border,
		border,
	}

	ml := types.MapLabel{
		Label:         title,
		TypeFace:      typeFace,
		Rotation:      0.0,
		ShouldSplitFn: shouldSplitFn,
	}
	ml.Build(noscale) //nolint:errcheck
	ml.SetCenter(titlePos)

	err := ml.DrawText(gc)
	if err != nil {
		t.Fatal(err)
	}

	y += fontSize * 2

	featCodes := []types.FeatCode{
		// Road.shp
		types.FEAT_MOTORWAY,                                // Motorway : Line
		types.FEAT_PRIMARY_ROAD,                            // Primary Road : Line
		types.FEAT_A_ROAD,                                  // A Road : Line
		types.FEAT_B_ROAD,                                  // B Road : Line
		types.FEAT_MINOR_ROAD,                              // Minor Road : Line
		types.FEAT_LOCAL_STREET,                            // Local Street : Line
		types.FEAT_PRIVATE_ROAD,                            // Private Road, Public Access : Line
		types.FEAT_PEDESTRIAN_STREET,                       // Pedestrianised Street : Line
		types.FEAT_MOTORWAY_COLLAPSED_DUAL_CARRIAGEWAY,     // Motorway, Collapsed Dual Carriageway : Line
		types.FEAT_PRIMARY_ROAD_COLLAPSED_DUAL_CARRIAGEWAY, // Primary Road, Collapsed Dual Carriageway : Line
		types.FEAT_A_ROAD_COLLAPSED_DUAL_CARRIAGEWAY,       // A Road, Collapsed Dual Carriageway : Line
		types.FEAT_B_ROAD_COLLAPSED_DUAL_CARRIAGEWAY,       // B Road, Collapsed Dual Carriageway : Line
		types.FEAT_MINOR_ROAD_COLLAPSED_DUAL_CARRIAGEWAY,   // Minor Road, Collapsed Dual Carriageway : Line

		// RailwayStation.shp
		types.FEAT_LIGHT_RAPID_TRANSIT_STATION,                                // Light Rapid Transit Station
		types.FEAT_RAILWAY_STATION,                                            // Railway Station
		types.FEAT_LONDON_UNDERGROUND_STATION,                                 // London Underground Station
		types.FEAT_RAILWAY_STATION_AND_LONDON_UNDERGROUND_STATION,             // Railway Station And London Underground Station
		types.FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_RAILWAY_STATION,            // Light Rapid Transit Station And Railway Station
		types.FEAT_LIGHT_RAPID_TRANSIT_STATION_AND_LONDON_UNDERGROUND_STATION, // Light Rapid Transit Station And London Underground Station

		// Foreshore.shp
		types.FEAT_FORESHORE, // Foreshore : Line

		// AdministrativeBoundary.shp
		types.FEAT_NATIONAL_BOUNDARY,                   // National Boundary
		types.FEAT_PARISH_OR_COMMUNITY_BOUNDARY,        // Parish or Community Boundary
		types.FEAT_DISTRICT_OR_LONDON_BOROUGH_BOUNDARY, // District or London Borough Boundary
		types.FEAT_COUNTY_REGION_OR_ISLAND_BOUNDARY,    // County, Region Or Island Boundary
	}

	// draw the lines
	for id, featCode := range featCodes {
		shapeType, err := types.GetShapeType(featCode)
		if err != nil {
			t.Fatal(err)
		}

		renderType, err := types.GetRenderType(shapeType.RenderType)
		if err != nil {
			t.Fatal(err)
		}

		labelStyle, err := types.GetLabelStyle(shapeType.TextFormat)
		if err != nil {
			t.Fatal(err)
		}

		a := types.LayerArtifact{
			ID:         strconv.Itoa(id),
			LayerType:  "lines_test",
			FeatCode:   featCode,
			GeomType:   geos.TypeIDLineString,
			ShapeType:  shapeType,
			RenderType: renderType,
			LabelFieldNames: []string{
				"DISTNAME",
			},
			LabelText: map[string]string{
				"DISTNAME": fmt.Sprintf("[%v] %v", strconv.Itoa(int(featCode)), featCode.String()),
			},
		}

		dataType := a.ShapeType.DataType

		// draw 2 versions of the line
		// one zoomed in
		// one actual size

		// zoomed
		zoomedLineThickness := a.RenderType.LineThickness * zoomedLineZoom

		textHeight := labelStyle.LabelTextStyle.Size
		remainingHeight := (lineHeight - textHeight) / 2
		adj := (remainingHeight / 2) + zoomedLineThickness/2

		wktZoomedLine := fmt.Sprintf("LINESTRING(%v %v, %v %v)", border, y+adj, zoomedLineEndX, y+adj)
		geomZoomedLine, err := gctx.NewGeomFromWKT(wktZoomedLine)
		if err != nil {
			t.Fatal(err)
		}

		mgZoomed := types.MapGeom{
			ID:          a.ID,
			LayerType:   a.LayerType,
			Geometry:    geomZoomedLine,
			DataType:    dataType,
			FeatCode:    featCode,
			LineCol:     a.ShapeType.LineCol,
			FillCol:     a.ShapeType.FillCol,
			Thickness:   zoomedLineThickness,
			StrokeWidth: a.RenderType.StrokeWidth * zoomedLineZoom,
		}
		err = makeimage.DrawGeom(gc, mgZoomed, noscale)
		if err != nil {
			t.Fatal(err)
		}

		// un zoomed
		lineThickness := a.RenderType.LineThickness
		wktUnZoomedLine := fmt.Sprintf("LINESTRING(%v %v, %v %v)", zoomedLineEndX+border, y+adj, zoomedLineEndX+lineLength-border/2, y+adj)
		geomUnZoomedLine, err := gctx.NewGeomFromWKT(wktUnZoomedLine)
		if err != nil {
			t.Fatal(err)
		}

		mg := types.MapGeom{
			ID:          a.ID,
			LayerType:   a.LayerType,
			Geometry:    geomUnZoomedLine,
			DataType:    dataType,
			FeatCode:    featCode,
			LineCol:     a.ShapeType.LineCol,
			FillCol:     a.ShapeType.FillCol,
			Thickness:   lineThickness,
			StrokeWidth: a.RenderType.StrokeWidth,
		}
		err = makeimage.DrawGeom(gc, mg, noscale)
		if err != nil {
			t.Fatal(err)
		}

		// 2 labels, one for each line
		// first with large uniform font size

		// style
		fontName := labelStyle.LabelTextStyle.FontName
		fontSpacing := 0.0

		label, err := makeimage.GetLabelText(a.LabelText, a.LabelFieldNames[0])
		if err != nil {
			t.Fatal(err)
		}

		// set face
		fontSize = 11.0
		fontCol = colours.Darkgrey

		typeFace := fonts.TypeFace{
			Name:        fontName,
			Spacing:     fontSpacing,
			StrokeStyle: strokeStyle,
			Color:       fontCol,
			Size:        fontSize,
			FontData:    fontData,
			Face:        fonts.GetFace(gc, fontData, fontSize),
		}

		// position
		posA := []float64{
			border,
			y + fontSize/2,
		}

		fonts.SetFont(gc, typeFace)
		err = geom.DrawString(gc, posA, 0, label)
		if err != nil {
			t.Fatal(err)
		}

		// set face
		fontSize := labelStyle.LabelTextStyle.Size
		fontCol = labelStyle.LabelTextStyle.Color

		typeFace = fonts.TypeFace{
			Name:        fontName,
			Spacing:     fontSpacing,
			StrokeStyle: strokeStyle,
			Color:       fontCol,
			Size:        fontSize,
			FontData:    fontData,
			Face:        fonts.GetFace(gc, fontData, fontSize),
		}

		// position
		posB := []float64{
			zoomedLineEndX + border,
			y + fontSize/2,
		}

		fonts.SetFont(gc, typeFace)
		err = geom.DrawString(gc, posB, 0, label)
		if err != nil {
			t.Fatal(err)
		}

		// next row
		y += lineHeight
	}

	// Return the output filename
	err = savePNG("test-output/roadlines_test.png", m)
	if err != nil {
		t.Fatal(err)
	}
}

func savePNG(fname string, m image.Image) error {
	dir, _ := path.Split(fname)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	return draw2dimg.SaveToPngFile(fname, m)
}
