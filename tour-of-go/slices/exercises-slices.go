package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) (pic [][]uint8) {
	pic = allocate(dx, dy)
	color(pic)
	return
}

func color(pic [][]uint8) {
	for i := 0; i < len(pic); i = i + 1 {
		for j := 0; j < len(pic[i]); j = j + 1 {
			pic[i][j] = uint8((i + j) / 2)
		}
	}
}

func allocate(dx, dy int) (pic [][]uint8) {

	/*
		slices are of type []T
		So here we ask make to give us a slice to type []uint8
		i.e. a slice to a slice.
		Actually dy slices each of which are slices.
	*/
	pic = make([][]uint8, dy)
	for i := 0; i < dy; i = i + 1 {
		pic[i] = make([]uint8, dx)
	}
	return
}

func main() {
	pic.Show(Pic)
}
