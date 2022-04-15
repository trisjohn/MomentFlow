package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type TestLayout struct {
	line *canvas.Line
	circle *canvas.Circle
	Rect []*canvas.Rectangle
	canvas *fyne.Container
	lastSize fyne.Size
}

func (t * TestLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	
	size = fyne.NewSize(diameter,diameter)
	stroke := diameter / 40
	t.circle.Resize(size)
	t.circle.Move(fyne.NewPos(100,100))
	t.line.StrokeWidth = stroke
	t.line.Resize(fyne.NewSize(300,300))
	
	for i,v := range t.Rect {
		w := i % int(size.Width)
		// fmt.Println("W:",w, "I:",i)
		v.Resize(fyne.NewSize(50,50))
		v.Move(fyne.NewPos(float32(i*50),float32(w*50)))
	}
	fmt.Println("Layout complete.", t.lastSize, size, t.lastSize==size)
	// new_rect := &canvas.Rectangle{FillColor: color.Gray{0x50}}
	// new_rect.Resize(fyne.NewSize(30,30))
	// for i,v := range t.Rect {
	// 	v.Resize(fyne.NewSize(50,50))
	// 	v.Move(fyne.NewPos(float32(i*50),float32(i*50)))
	// 	new_rect.Move(fyne.NewPos(float32(i*50+30), float32(i*50+30)))
	// }
	// if t.lastSize != size {
	// 	t.lastSize = size
	// 	t.Rect = append(t.Rect, new_rect)
	// 	fmt.Println("New Rectangle added", new_rect.Position(), len(t.Rect))
	// }
	// t.canvas.Add(new_rect)
}

func(t *TestLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(500,500)
}

// Return an array of canvas rectangle pointers
func Rectangles(fillcol, strokecol color.Color, num int) []*canvas.Rectangle {
	rects := make([]*canvas.Rectangle,0)
	for i := 0; i < num; i++ {
		rects = append(rects, &canvas.Rectangle{StrokeColor: strokecol, FillColor: fillcol, StrokeWidth: 2})
	}
	return rects
}

func(t *TestLayout) Render() *fyne.Container {
	fmt.Println("Render test layout")
	t.line = &canvas.Line{StrokeColor: color.Gray{0xE0}, StrokeWidth: 2}
	t.circle = &canvas.Circle{StrokeColor: color.Gray{0xE0}, StrokeWidth: 2}
	t.Rect = Rectangles(color.Gray{0x70},color.Gray{0xE0}, 50)
	con := container.NewWithoutLayout(t.line, t.circle)
	t.lastSize = t.MinSize(nil)
	for _,v := range t.Rect {
		con.Add(v)
	}
	con.Layout = t
	t.canvas = con
	return con
}