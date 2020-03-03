package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var chessboard = [8][8]int{
	{0, 1, 0, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 1, 0},
}

// размер картинки в пикселях
const picSize = 300

func main() {
	cBlack := color.RGBA{0, 0, 0, 255}
	cWhite := color.RGBA{255, 255, 255, 255}

	rectImg := image.NewRGBA(image.Rect(0, 0, picSize, picSize))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{cWhite}, image.ZP,
		draw.Src)

	// Рисуем доску
	for x := 0; x < picSize; x++ {
		for y := 0; y < picSize; y++ {
			i := int(x * 8 / picSize)
			j := int(y * 8 / picSize)

			if chessboard[i][j] == 1 {
				rectImg.Set(x, y, cBlack)
			}

		}
	}

	file, err := os.Create("chessboard.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	png.Encode(file, rectImg)
}
