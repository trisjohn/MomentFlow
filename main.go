package main

import (
	"fmt"
	"gui/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func main() {
	fmt.Println("Moment Flow initiated.")
	a := app.New()
	w := a.NewWindow("Moment Flow")

	info := gui.PlaceHolder("No Object Selected")

	v := container.NewVBox(
		info,
		widget.NewButton("New Location", func() {
			// Create a Location node
			info.SetText("New Location")
		}),
		widget.NewButton("New Scene", func() {
			// Create a Scene node
			info.SetText("New Scene")
		}),
	)
	prop := gui.PlaceHolder("Fuck you piece of shit")
	dynGrid := &gui.TestLayout{}
	g := dynGrid.Render()
	fmt.Println("grid created.")

	// h := container.NewAdaptiveGrid(2,
	// 	g ,
	// 	v,
	// )
	h := container.New(&gui.GridWithMenu{},
		g ,
		v,
		prop,
	)
	
	w.Resize(fyne.NewSize(500,500))
	w.SetContent(h)
	w.ShowAndRun()
}