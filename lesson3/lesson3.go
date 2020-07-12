package main

import (
	"bufio"
	"encoding/base64"
	"image"
	"image/png"
	"io"
	"os"
)

func Show(f func(dx, dy int) [][]uint8) {
	const (
		dx = 256 * 5
		dy = 256 
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = 0
			m.Pix[i+1] = v
			m.Pix[i+2] = v
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dx)
	for i := 0; i < dy; i++ {
		picture[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			picture[i][j] = uint8((i ^ j) * 2)
		}
	}
	return picture
}

func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}

func main() {
	Show(Pic)
}
