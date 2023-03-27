package types

import (
	"errors"
	"fmt"
	"image/color"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/rockwell-uk/go-geos-draw/geom"
	"github.com/rockwell-uk/go-text/fonts"
	"github.com/rockwell-uk/go-text/text"
	"github.com/twpayne/go-geos"
)

type LineData struct {
	Ascent  float64
	Descent float64
	Width   float64
}

type MapLabel struct {
	ID            string
	LayerType     string
	LabelType     LabelType
	DataType      DataType
	FeatCode      FeatCode
	Geometry      *geos.Geom
	Label         string
	Lines         []string
	LineData      map[string]LineData
	Rotation      float64
	LineCol       color.RGBA
	FillCol       color.RGBA
	Zoom          float64
	Radius        float64
	Thickness     float64
	TypeFace      fonts.TypeFace
	FaceMetrics   *fonts.FaceMetrics
	Dimensions    MapLabelDimensions
	Built         *bool
	ShouldSplitFn func(s string) bool
}

type MapLabelDimensions struct {
	Center     *[]float64
	Width      *float64
	Height     *float64
	TextWidth  *float64
	TextHeight *float64
	LineHeight *float64
	Border     *float64
	Envelope   *geos.Geom
}

func (d MapLabelDimensions) String() string {
	return fmt.Sprintf("Center %v\n"+
		"Width %v\n"+
		"Height %v\n"+
		"TextWidth %v\n"+
		"TextHeight %v\n"+
		"LineHeight %v\n"+
		"Border %v\n"+
		"Envelope %s\n",
		*d.Center,
		*d.Width,
		*d.Height,
		*d.TextWidth,
		*d.TextHeight,
		*d.LineHeight,
		*d.Border,
		d.Envelope,
	)
}

func (l *MapLabel) Build(scale func(x, y float64) (float64, float64)) error {
	if l.TypeFace.Face == nil {
		return errors.New("set font before building label")
	}

	// face metrics
	faceMetrics := l.getFaceMetrics()
	fontHeight := faceMetrics.Height
	border := fontHeight * 0.05
	l.FaceMetrics = &faceMetrics
	l.Dimensions.LineHeight = &fontHeight
	l.Dimensions.Border = &border

	// text width and number of lines
	l.getLineDetails()

	// text height
	textHeight := l.getTextHeight()
	l.Dimensions.TextHeight = &textHeight

	// width
	width := l.getWidth()
	l.Dimensions.Width = &width

	// height
	height := l.getHeight()
	l.Dimensions.Height = &height

	// label center
	center, err := l.getCenter(scale)
	if err != nil {
		return err
	}
	l.SetCenter(center)

	// envelope
	l.Dimensions.Envelope = l.getEnvelope()

	built := true
	l.Built = &built

	return nil
}

func (l *MapLabel) SetCenter(center []float64) {
	l.Dimensions.Center = &center

	// rebuild envelope
	l.Dimensions.Envelope = l.getEnvelope()
}

func (l *MapLabel) getFaceMetrics() fonts.FaceMetrics {
	if l.FaceMetrics == nil {
		fm := fonts.GetFaceMetrics(l.TypeFace)
		l.FaceMetrics = &fm
	}

	return *l.FaceMetrics
}

func (l *MapLabel) getLineDetails() {
	if l.ShouldSplitFn == nil {
		l.ShouldSplitFn = ShouldSplit
	}

	var longestLineWidth float64
	var lines []string = text.SplitStringInTwo(l.Label, l.ShouldSplitFn)
	var lineData = make(map[string]LineData)

	for _, v := range lines {
		lineTextWidth := fonts.GetTextWidth(l.TypeFace, v)
		if lineTextWidth > longestLineWidth {
			longestLineWidth = lineTextWidth + (l.TypeFace.StrokeStyle.Width * 2)
		}

		ascent, descent := fonts.GetTextHeight(l.TypeFace, v)
		lineData[v] = LineData{
			Ascent:  ascent,
			Descent: descent,
			Width:   lineTextWidth,
		}
	}

	l.Dimensions.TextWidth = &longestLineWidth
	l.Lines = lines
	l.LineData = lineData
}

func (l *MapLabel) getCenter(scale func(x, y float64) (float64, float64)) ([]float64, error) {
	if l.Geometry == nil {
		return []float64{}, errors.New("label geometry not set")
	}

	center, err := geom.GetGeometryCenter(l.Geometry, scale)
	if err != nil {
		return []float64{}, err
	}

	return center, nil
}

func (l *MapLabel) getWidth() float64 {
	return *l.Dimensions.TextWidth + *l.Dimensions.Border
}

func (l *MapLabel) getHeight() float64 {
	return *l.Dimensions.TextHeight + (*l.Dimensions.Border * 2)
}

func (l *MapLabel) getTextHeight() float64 {
	return l.FaceMetrics.Height * float64(len(l.Lines))
}

func (l MapLabel) getEnvelope() *geos.Geom {
	labelWidth := *l.Dimensions.Width
	labelHeight := *l.Dimensions.Height
	center := *l.Dimensions.Center

	expansionX := labelWidth / 2
	expansionY := labelHeight / 2

	xMin := center[0] - expansionX
	xMax := center[0] + expansionX
	yMin := center[1] - expansionY
	yMax := center[1] + expansionY

	g, _ := geom.BoundsGeom(
		xMin,
		xMax,
		yMin,
		yMax,
	)

	return g
}

func (l MapLabel) DrawBackground(gc *draw2dimg.GraphicContext) error {
	if l.TypeFace.BackgroundStrokeStyle.Color == nil {
		return errors.New("label TypeFace.BackgroundStrokeStyle.Color cannot be nil")
	}

	return geom.DrawPolygon(gc,
		l.Dimensions.Envelope,
		l.TypeFace.BackgroundColor,
		l.TypeFace.BackgroundStrokeStyle.Color,
		l.TypeFace.BackgroundStrokeStyle.Width,
		func(x, y float64) (float64, float64) { return x, y },
	)
}

func (l MapLabel) DrawText(gc *draw2dimg.GraphicContext) error {
	// set font before label
	fonts.SetFont(gc, l.TypeFace)

	labelCenter := *l.Dimensions.Center
	longestLineWidth := *l.Dimensions.TextWidth
	lineHeight := *l.Dimensions.LineHeight
	border := *l.Dimensions.Border

	// set start point
	var yOffset float64
	if len(l.Lines) == 1 {
		yOffset = l.FaceMetrics.Descent/2 - lineHeight/2
	}
	for _, w := range l.LineData {
		yOffset += w.Descent / 2
	}

	startX := labelCenter[0] - longestLineWidth/2 + border // indent by border
	startY := labelCenter[1] - yOffset

	// write the lines
	for i, w := range l.Lines {
		lineSpaceDiff := longestLineWidth - l.LineData[w].Width

		lineX := startX + lineSpaceDiff/2
		lineY := startY
		if i > 0 {
			lineY += lineHeight
		}

		err := geom.DrawString(gc, []float64{lineX, lineY}, l.Rotation, w)
		if err != nil {
			return err
		}
	}

	return nil
}

func ShouldSplit(s string) bool {
	// dont split short strings
	if len(s) < 12 {
		return false
	}

	n := strings.Count(s, " ")

	// if there arent any spaces dont split
	if n == 0 {
		return false
	}

	words := strings.Split(s, " ")

	// if there are 2 words
	// if the first or secord word is short dont split
	if len(words) == 2 {
		if len(words[0]) < 4 {
			return false
		}
		if len(words[1]) < 4 {
			return false
		}
	}

	return true
}
