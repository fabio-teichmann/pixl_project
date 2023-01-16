package pxcanvas

import "fyne.io/fyne/v2"

// fyne.PointEvent holds mouse events
func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff
	pxCanvas.Refresh()
}

// scaling: up or down
func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}
	default:
		pxCanvas.PxSize = 10
	}
}
