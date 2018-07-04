package main

import (
	"fmt"
	"image"
	"time"
)

//0.24186126600000002
//0.224485236
//0.010330342000000001

func main() {
	w, h := 600, 480
	w *= 4
	h *= 4
	rect := image.Rect(0, 0, w, h)

	rgba := image.NewRGBA(rect)

	//	image.RGBA
	{
		m := rgba
		t1 := time.Now()
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				c := m.At(i, j)
				r, g, b, a := c.RGBA()
				r /= 256
				g /= 256
				b /= 256
				_, _, _, _ = r, g, b, a
			}
		}
		t2 := time.Now()
		fmt.Println(t2.Sub(t1).Seconds())
	}

	//	image.Image
	{
		m := image.Image(rgba)
		t1 := time.Now()
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				c := m.At(i, j)
				r, g, b, a := c.RGBA()
				r /= 256
				g /= 256
				b /= 256
				_, _, _, _ = r, g, b, a
			}
		}
		t2 := time.Now()
		fmt.Println(t2.Sub(t1).Seconds())
	}

	//	raw
	//	interface と color の変換モデルがない場合は体感的に20倍ほど速いかな
	{
		t1 := time.Now()
		pix := m.Pix
		stride := m.Stride
		//		off := 0
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				off := i*4 + stride*j
				r, g, b, a := pix[off], pix[off+1], pix[off+2], pix[off+3]
				_, _, _, _ = r, g, b, a
				//				off += 4
			}
		}
		t2 := time.Now()
		fmt.Println(t2.Sub(t1).Seconds())
	}
}
