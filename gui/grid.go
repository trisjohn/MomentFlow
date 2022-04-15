package gui

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

type GridWithMenu struct {

} 

func(g *GridWithMenu) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()
		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w,h)
}

// Must be in order of grid > new obj panel > property panel
func(g *GridWithMenu) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	// Layout grid on the left side, and New object buttons on the top right, with property window on the buttom right
	pos := fyne.NewPos(0,0)
	maxHigh := containerSize.Height
	maxWidth := containerSize.Width
	if len(objects) != 3 {
		RaiseFatal("GridWithMenu container is passed wrong size of objects, it needs:\nGrid,ObjSearch,ObjProperty\nReceived:", len(objects))
	}

	// Instead of a for loop, use individual logics
	objSize := objects[1].MinSize()
	propSize := objects[2].MinSize()

	trueWidth := objSize.Width
	if propSize.Width > trueWidth {
		trueWidth = propSize.Width
	}

	size := fyne.NewSize(maxWidth-trueWidth, maxHigh)
	objects[0].Resize(size)
	objects[0].Move(pos)

	pos = fyne.NewPos(size.Width,0)
	objects[1].Resize(objSize)
	objects[1].Move(pos)

	pos = fyne.NewPos(maxWidth-trueWidth, maxHigh-propSize.Height)
	objects[2].Resize(propSize)
	objects[2].Move(pos)

}

type DyanmicGridLayout struct {
	tiles []*canvas.Rectangle
	canvas fyne.CanvasObject
	tilesize float32
}

func (g *DyanmicGridLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(500,500)
}

// Create a new square (canvas.Rectangle) at the specified position
func newGridTile(tilesize float32, pos fyne.Position) (r *canvas.Rectangle) {
	r = &canvas.Rectangle{}
	r.FillColor = color.Gray{0x70}
	r.StrokeWidth = 1
	r.StrokeColor = color.Gray{0xE0}
	r.Resize(fyne.NewSize(tilesize,tilesize))
	r.Move(pos)
	r.Show()
	r.Refresh()
	return r
}

func (g *DyanmicGridLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	rows := float32(math.RoundToEven(float64(size.Height / g.tilesize)))
	cols := float32(math.RoundToEven(float64(size.Width / g.tilesize)))
	fmt.Println("Dynamic Grid Layout called.", size, len(g.tiles), rows, cols)
	// Save the objects to the g. Everytime, check if current g has the rigth number of anchors
	number := int(math.RoundToEven(float64(rows * cols)))
	l := len(g.tiles)
	for l != number {
		l = len(g.tiles)
		if l == 0 {
			new_pos := fyne.NewPos(0,0)
			g.tiles = append(g.tiles, newGridTile(g.tilesize, new_pos))
		} else if l < number {
			// Too few tiles
			last := g.tiles[l-1]
			last_pos := last.Position()
			new_pos := fyne.NewPos(last_pos.X+g.tilesize,last_pos.Y)
			if last_pos.X/g.tilesize + 1 > cols {
				// Too far right
				new_pos = fyne.NewPos(0, last_pos.Y+g.tilesize)
			}
			rect := newGridTile(g.tilesize,new_pos)
			// new_anchor := AnchorObject{rect}
			// new_anchor.Rectangle.Move(new_pos)
			g.tiles = append(g.tiles, rect)
		} else if l > number {
			// Too many tiles
			new_tiles := make([]*canvas.Rectangle, 0)
			fmt.Println(size, l, number)
			// Loop through tiles and only add tiles that fit within size
			for _, v := range g.tiles {
				pos := v.Position()
				if pos.X <= size.Width && pos.Y <= size.Height{
					new_tiles = append(new_tiles, v)
				} 
			}
			g.tiles = new_tiles
			fmt.Println("Tiles deleted.", len(g.tiles), number)
			break
		}
	}
}

func (g *DyanmicGridLayout) Render() *fyne.Container {
	tiles := make([]fyne.CanvasObject,0)
	for _, v := range g.tiles {
		tiles = append(tiles, v)
	}
	contain := container.NewWithoutLayout(tiles...)
	contain.Layout = g
	g.canvas = contain
	if g.tilesize <= 0 {
		g.tilesize = 10
	}
	canvas.Refresh(g.canvas)
	return contain
}

type AnchorObject struct {
	*canvas.Rectangle
}

type DynamicGrid struct {
	widget.BaseWidget
	Tilesize float32
	tiles []fyne.CanvasObject
}

func NewDynamicGrid(tilesize float32) DynamicGrid {
	d := DynamicGrid{}
	d.Tilesize = tilesize
	return d
}

// Grid renderer
type DynamicGridRender struct {
	objects []fyne.CanvasObject
	grid *DynamicGrid
}

func(grid DynamicGridRender) Layout( size fyne.Size) {
	fmt.Println("Dynamic Grid Layout called.")
}
func(grid DynamicGridRender) MinSize() fyne.Size {
	return fyne.NewSize(500,500)
}
func(grid DynamicGridRender) Refresh() {
	fmt.Println("Dynamic Grid refresh called.")
}

func(grid DynamicGridRender) Destroy() {
	fmt.Println("Dynamic Grid destroy called.")
}

// Return a grid, given a number of rows by a number of columns
func CreateGrid(rows, columns int) *fyne.Container {
	grid := container.NewGridWithColumns(columns)

	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			bg := canvas.NewRectangle(color.Gray{0x90})
			grid.Add(bg)
		}
	}

	fmt.Println("New Grid:",rows, columns)
	return grid
}