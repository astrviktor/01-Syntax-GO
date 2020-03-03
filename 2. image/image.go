package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP,
		draw.Src)

	// Горизонтальные линии
	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y = y + 20 {
			rectImg.Set(x, y, color.Black)
		}
	}

	// Вертикальные линии
	for y := 0; y < 200; y++ {
		for x := 0; x < 200; x = x + 20 {
			rectImg.Set(x, y, color.Black)
		}
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	png.Encode(file, rectImg)
}
