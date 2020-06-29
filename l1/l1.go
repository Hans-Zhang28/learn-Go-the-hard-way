package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func AddPhtoFrame() {
	//TODO:user the res/gophergala.jpg to generate a image and write to res/m.jpg which is similar to like the logo in the README.md
	abs, err := filepath.Abs("./res/gophergala.jpg")
	if err != nil {
		log.Fatal(err)
	}

	reader, err := os.Open(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	// config, _, err := image.DecodeConfig(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// draw ractangle on the image
	red := color.RGBA{255, 0, 0, 1}
	green := color.RGBA{0, 255, 0, 1}
	blue := color.RGBA{0, 0, 255, 1}
	yellow := color.RGBA{255, 255, 0, 1}

	b := img.Bounds()
	newImg := image.NewRGBA(b)
	println(b.Dx())
	println(b.Dy())
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)
	// draw red rectangle at the bottom
	draw.Draw(newImg, image.Rect(100, b.Dy()-100, b.Dx()-100, b.Dy()), &image.Uniform{red}, image.ZP, draw.Src)
	// draw the yallow rectangle at the top
	draw.Draw(newImg, image.Rect(100, 0, b.Dx()-100, 100), &image.Uniform{yellow}, image.ZP, draw.Src)
	// draw the blue rectangle at the left
	draw.Draw(newImg, image.Rect(0, 0, 100, b.Dy()), &image.Uniform{blue}, image.ZP, draw.Src)
	// draw the blue rectangle at the right
	draw.Draw(newImg, image.Rect(b.Dx()-100, 0, b.Dx(), b.Dy()), &image.Uniform{green}, image.ZP, draw.Src)

	abs, err = filepath.Abs("./res/my_res.jpg")
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	opt := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(outfile, newImg, &opt)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	AddPhtoFrame()
}
