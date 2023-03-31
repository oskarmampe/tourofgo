package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	pixels [][]uint8
}

func createImage() Image {
	pic := make([][]uint8, 255)
	for i := range pic {
		row := make([]uint8, 255)
		for j := range row {
			row[j] = uint8(i ^ j) // uint8(i&j^(j+i))
		}
		pic[i] = row
	}
	return Image{pic}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{255, 255}}
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, i.pixels[x][y], i.pixels[x][y]}
}

func main() {
	m := createImage()
	pic.ShowImage(m)
}
