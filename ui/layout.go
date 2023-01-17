package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	SetUpMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	// order of border container = top, bottom, left, right, everything else included in center
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas)

	app.PixlWindow.SetContent(appLayout)
}
