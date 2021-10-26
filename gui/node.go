// Nodes represent each game object. These are Game objects that need connections to other game objects

package gui

import "fyne.io/fyne/v2/widget"

// Returns a label that acts a place holder for any given object, useful for building
func PlaceHolder(title string) *widget.Label {
	w := widget.NewLabel(title)
	return w
}