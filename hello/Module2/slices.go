package main

import "code.google.com/p/go-tour/pic"

fun Pic(dx , dy int)[][]uint8 {
	img := make([][]uint8 , dy)
	for y := range img{
		for x := range img[y]{
			img[y][x] = uint8(x+y)
		}
	}
	return img
}

func main(){
	pic.Show(Pic)
}


