package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

var characters = []string{"█", "▓", "▒", "░", " "}

const (
	blackpoint = 0x2000
	whitepoint = 0xcfff
)

func main() {
	reader, err := os.Open("image.jpg")
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	newWidth := 50
	bounds := m.Bounds()
	ratio := float64(bounds.Dy()) / float64(bounds.Dx())
	newHeight := int(float64(newWidth) * ratio)

	newImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := x * bounds.Dx() / newWidth
			srcY := y * bounds.Dy() / newHeight
			newImage.Set(x, y, m.At(srcX, srcY))
		}
	}

	fmt.Printf("newImage.Bounds(): %v\n", newImage.Bounds())

	var res string
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			r, g, b, _ := newImage.At(x, y).RGBA()
			fmt.Printf("r: %v\n", r)

			brightness := (r + g + b) / 3
			step := (whitepoint - blackpoint) / uint32(len(characters)-1)
			index := (brightness - blackpoint) / step
			if index >= uint32(len(characters)) {
				index = uint32(len(characters)) - 1
			}
			res += characters[index]
		}
		res += "\n"
	}

	fmt.Printf("res: %v\n", res)
}
