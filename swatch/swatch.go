package swatch

import (
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget // embedded
	Selected          bool
	Color             color.Color
	SwatchIndex       int             // to select swatch
	clickHandler      func(s *Swatch) // supply code on what to do when we click on it
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh() // redraw and update color change on screen
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}

	swatch.ExtendBaseWidget(swatch) // supplies all state information to fyne implementation
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
