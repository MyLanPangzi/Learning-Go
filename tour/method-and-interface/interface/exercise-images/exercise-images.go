package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
	v    uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{
		R: i.v,
		G: i.v,
		B: 255,
		A: 255,
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

//import "golang.org/x/tour/pic"
//
//type Image struct{}
//
//func main() {
//	m := Image{}
//	pic.ShowImage(m)
//}
