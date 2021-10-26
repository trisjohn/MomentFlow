package gui

import (
	"fmt"
	"image/color"

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

type DynamicGrid struct {
	widget.BaseWidget
	Tilesize float32
}

func(grid *DynamicGrid) Layout( size fyne.Size) {
	fmt.Println("Dynamic Grid Layout called.")
}
func(grid *DynamicGrid) MinSize() fyne.Size {
	return fyne.NewSize(500,500)
}
func(grid *DynamicGrid) Refresh() {
	fmt.Println("Dynamic Grid refresh called.")
}

func(grid *DynamicGrid) Objects() []fyne.CanvasObject {
	
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