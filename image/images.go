// Define Image type, implement the necessary methods, and 
// call pic.ShowImage.
// The purpose of the exercise is to enhance the understanding on golang's
// interface concept. 

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

// Implement necessary methods of Image interface

func (img Image) ColorModel() color.Model {
	return color.RGBAModel;
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {

	image_generator := func (x, y int) uint8 {
		return uint8((x+y)/2)
	} 

	v := image_generator(x, y)

	return color.RGBA{v, v, 255, 255}
}

func main() {
	w, h := 128, 128
	m := Image{}
	pic.ShowImage(m)
}