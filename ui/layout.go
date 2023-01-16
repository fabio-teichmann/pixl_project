package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	// order of border container = top, bottom, left, right
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker)

	app.PixlWindow.SetContent(appLayout)
}
