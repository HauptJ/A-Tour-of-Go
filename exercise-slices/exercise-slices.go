package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	
	for x := range image {
		image[x] = make([]uint8, dx)
	}
	
	for x := range image {
		for y := range image[x] {
			image[x][y] = uint8(x)^uint8(y) + uint8(y)^uint8(x)
		}
	}
	
	return image
	
}

func main() {
	pic.Show(Pic)
}