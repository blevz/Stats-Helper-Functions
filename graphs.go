package main

import (
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"image"
	_ "image/color"
)

type graph struct {
	data        *randomVarStore
	graphBounds image.Rectangle
	dataBounds  rect
}

type point struct {
	X, Y float64
}

type rect struct {
	Min, Max point
}

func (r rect) Dx() float64 {
	return r.Max.X - r.Min.X
}

func (r rect) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

func (g graph) init() {

}

func drawRectOnC(c *draw2d.ImageGraphicContext, r image.Rectangle) {
	c.MoveTo(float64(r.Min.X), float64(r.Min.Y))
	c.LineTo(float64(r.Min.X), float64(r.Max.Y))
	c.LineTo(float64(r.Max.X), float64(r.Max.Y))
	c.LineTo(float64(r.Max.X), float64(r.Min.Y))
	c.LineTo(float64(r.Min.X), float64(r.Min.Y))
	c.Stroke()
}

func (g *graph) toImage() (m *image.RGBA) {
	//Make image and set graph bounds inside the image
	m = image.NewRGBA(image.Rect(0, 0, 300, 300))
	g.graphBounds.Min = m.Bounds().Max.Div(20)
	g.graphBounds.Max = g.graphBounds.Min.Mul(19)

	c := draw2d.NewGraphicContext(m)
	drawRectOnC(c, g.graphBounds)

	//Get the min and max x and y and set g.data.bounds
	sk := g.data.getSortedKeys()
	xmin := sk[0]
	xmax := sk[len(sk)-1]
	var ymin, ymax float64 = g.data.data[sk[0]], g.data.data[sk[0]]
	for _, k := range sk {
		if ymin > g.data.data[k] {
			ymin = g.data.data[k]
		} else if ymax < g.data.data[k] {
			ymax = g.data.data[k]
		}
	}
	g.dataBounds.Min = point{xmin, ymin}
	g.dataBounds.Max = point{xmax, ymax}

	// Draw the line, scaling done in fucntion
	lastX := sk[0]
	lastY := g.data.data[sk[0]]
	for _, k := range sk {
		pa := point{lastX, lastY}
		pb := point{k, g.data.data[k]}
		lastX = k
		lastY = g.data.data[k]
		g.drawGraphLine(m, c, pa, pb)
	}
	return
}

func (g graph) localizePoint(p point) (toRet point) {
	toRet.X = float64(g.graphBounds.Min.X) + (p.X-g.dataBounds.Min.X)*float64(g.graphBounds.Dx())/g.dataBounds.Dx()
	toRet.Y = float64(g.graphBounds.Max.Y) - (p.Y-g.dataBounds.Min.Y)*float64(g.graphBounds.Dy())/g.dataBounds.Dy()
	fmt.Println(toRet.X, toRet.Y)
	return
}

func (g *graph) drawGraphLine(m *image.RGBA, c *draw2d.ImageGraphicContext, ps, pe point) {
	p1 := g.localizePoint(ps)
	p2 := g.localizePoint(pe)
	if Verbosity >= Image_Verbosity {
		fmt.Printf("Convert line: (%f, %f)->(%f, %f)\n", ps.X, ps.Y, pe.X, pe.Y)
		fmt.Printf("To line: (%f, %f)->(%f, %f)\n\n", p1.X, p1.Y, p2.X, p2.Y)
	}
	c.MoveTo(float64(p2.X), float64(p2.Y))
	c.LineTo(float64(p1.X), float64(p1.Y))
	c.Stroke()
	return
}
