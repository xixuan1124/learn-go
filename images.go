package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	width, height int
	ColorAt func (x, y int) uint8
}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (i Image) At(x int, y int) color.Color{
	c := i.ColorAt(x, y)
	return color.RGBA{c, c, 255, 255}
}

func (i Image) Bounds() image.Rectangle{
	return image.Rect(0, 0, i.width, i.height)
}

func main() {
	m := Image{200, 200, func (x, y int) uint8 {return uint8 (x * y)}}
	pic.ShowImage(m)
}
