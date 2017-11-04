package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	filepath := os.Args[1]
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
		return
	}

	alphacolor := img.At(0, 0)
	imgrect := img.Bounds()

	outimage := image.NewRGBA(imgrect)

	for h := imgrect.Min.Y; h < imgrect.Max.Y; h++ {
		for w := imgrect.Min.X; w < imgrect.Max.X; w++ {
			c := img.At(w, h)

			if c == alphacolor {
				c = image.Transparent
			}

			outimage.Set(w, h, c)
		}
	}

	filename := strings.Split(filepath, ".")
	if len(filename) != 2 {
		return
	}
	newfilename := filename[0] + "-alpha." + filename[1]

	output, _ := os.Create(newfilename)
	defer output.Close()

	png.Encode(output, outimage)

}
