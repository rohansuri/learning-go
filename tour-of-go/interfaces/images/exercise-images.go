package main

import "golang.org/x/tour/pic"
import "image/color"
import "image"

// ideally BlackImage would be simply a RGBA value and not a type
// but just to quickly get done with the exercise
type BlackImage struct {
	image *image.RGBA
}

func NewBlackImage(rect image.Rectangle) *BlackImage {
	img := new(BlackImage)
	img.image = image.NewRGBA(rect)
	colorItBlack(img)
	return img
}

func colorItBlack(img *BlackImage) {
	color := color.Gray{0}
	rect := img.image.Bounds()
	for i := rect.Min.X; i <= rect.Max.X; i++ {
		for j := rect.Min.Y; j <= rect.Max.Y; j++ {
			img.image.Set(i, j, color)
		}
	}
}

func (i *BlackImage) ColorModel() color.Model {
	return i.image.ColorModel()
}

func (i *BlackImage) Bounds() image.Rectangle {
	return i.image.Bounds()
}

func (i *BlackImage) At(x, y int) color.Color {
	return i.image.At(x, y)
}

func main() {
	m := NewBlackImage(image.Rect(0, 0, 50, 50))
	pic.ShowImage(m)
}
