package customButton

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"

	"image/color"
)

type CustomTappableRectangle struct {
	widget.BaseWidget
	rect     *canvas.Rectangle
	onTapped func()
}

var _ fyne.Tappable = (*CustomTappableRectangle)(nil)

func NewCustomTappableRectangle(color color.Color, tappedCallback func()) *CustomTappableRectangle {
	r := &CustomTappableRectangle{
		rect:     canvas.NewRectangle(color),
		onTapped: tappedCallback,
	}
	r.ExtendBaseWidget(r)
	return r
}

func (c *CustomTappableRectangle) MinSize() fyne.Size {
	return fyne.NewSize(200, 200)
}

func (c *CustomTappableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return &customRectangleRenderer{
		rectangle: c.rect,
		objects:   []fyne.CanvasObject{c.rect},
	}
}

func (c *CustomTappableRectangle) Tapped(event *fyne.PointEvent) {
	c.onTapped()
}

type customRectangleRenderer struct {
	rectangle *canvas.Rectangle
	objects   []fyne.CanvasObject
}

func (r *customRectangleRenderer) MinSize() fyne.Size {
	return r.rectangle.MinSize()
}

func (r *customRectangleRenderer) Layout(size fyne.Size) {
	r.rectangle.Resize(size)
}

func (r *customRectangleRenderer) Refresh() {
	r.rectangle.Refresh()
}

func (r *customRectangleRenderer) Hide() {
	r.rectangle.Hide()
}

func (r *customRectangleRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *customRectangleRenderer) Destroy() {
}
