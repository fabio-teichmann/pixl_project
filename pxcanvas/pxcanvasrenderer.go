package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas
	canvasImage  *canvas.Image // part of Fyne toolkit, allows to display image present in PxCanvas struct; in PIXEL
	canvasBorder []canvas.Line // part of Fyne toolkit, use slice to draw borders around canvas
}

// WidgetRenderer interface implementation.
func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	// 4 border line objects + the actual image itself
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	return objects
}

// implementation not required
func (renderer *PxCanvasRenderer) Destroy() {}

/*
3 layout functions:
- layout
- layout of borders
- layout of image
*/

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size) // needs to be called first as it's resizing the image
	renderer.LayoutBorder(size)
}

func (renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols  // actual size of image, not on-screen representation
	imgPxHeight := renderer.pxCanvas.PxRows // actual size of image, not on-screen representation
	pxSize := renderer.pxCanvas.PxSize
	// move image to offset position (by default in the top left corner of drawing area)
	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))
	// makes pixel bigger on screen
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

func (renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := renderer.pxCanvas.CanvasOffset
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	bottom := &renderer.canvasBorder[1]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	top := &renderer.canvasBorder[3]
	top.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	top.Position2 = fyne.NewPos(offset.X, offset.Y)
}

// WidgetRenderer interface implementation
func (renderer *PxCanvasRenderer) Refresh() {
	// check if image needs to be reloaded
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.pxCanvas.reloadImage = false
	}

	renderer.Layout(renderer.pxCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}
