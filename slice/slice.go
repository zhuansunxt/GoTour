// Generating  2d-array representing a picture
// The purpose of this exercise is to get familiar with allocating memory
// space for golang slices 

package main 

import "golang.org/x/tour/pic"

func Pic(dy, dx int) [][]uint8 {
	var pic = make([][]uint8, dy)
	for i := 0; i < dy ; i++ {
		pic[i] = make([]uint8, dx)
	}

	for i:= 0; i < len(pic); i++ {
		for j := 0; j < len(pic[0]); j++ {
			pic[i][j] = uint8((i^j))
		}	
	}

	return pic;
}

func main() {
	pic.Show(Pic)
}