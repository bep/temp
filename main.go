package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("sunset.jpg")
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}

	toFile("sunset-resized.jpg", resize(src))
	toFile("sunset-fit.jpg", fit(src))
	toFile("sunset-fill.jpg", fill(src))
	toFile("sunset-fill-bottom-left.jpg", fillBottomLeft(src))
	toFile("sunset-fill-bottom-right.jpg", fillBottomRight(src))

}

func resize(src image.Image) *image.NRGBA {
	defer timer(time.Now(), "resize")
	return imaging.Resize(src, 500, 0, imaging.Box)
}

func fit(src image.Image) *image.NRGBA {
	defer timer(time.Now(), "fit")
	return imaging.Fit(src, 500, 500, imaging.Box)
}

func fill(src image.Image) *image.NRGBA {
	defer timer(time.Now(), "fill")
	return imaging.Fill(src, 500, 500, imaging.Center, imaging.Box)
}

func fillBottomLeft(src image.Image) *image.NRGBA {
	defer timer(time.Now(), "fill-bottom-left")
	return imaging.Fill(src, 200, 200, imaging.BottomLeft, imaging.Box)
}

func fillBottomRight(src image.Image) *image.NRGBA {
	defer timer(time.Now(), "fill-bottom-right")
	return imaging.Fill(src, 100, 200, imaging.BottomRight, imaging.Box)
}

func timer(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func toFile(name string, img *image.NRGBA) {
	err := imaging.Save(img, name)
	if err != nil {
		log.Fatalf("Save failed: %v", err)
	}
}
