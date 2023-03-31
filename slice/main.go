package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		row := make([]uint8, dx)
		for j := range row {
			row[j] = uint8(i^j) // uint8(i&j^(j+i))
		}
		pic[i] = row
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
