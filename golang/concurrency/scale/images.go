package main

import (
	"fmt"
	"image"

	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func scaleImage(srcPath, destPath string, size image.Rectangle) error {

	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	src, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("cannot decode into jpeg %s", err)
	}

	dest := image.NewRGBA(size)
	bounds := src.Bounds()
	draw.NearestNeighbor.Scale(dest, size, src, bounds, draw.Over, nil)

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer out.Close()

	if err := png.Encode(out, dest); err != nil {
		out.Close()
		os.Remove(destPath)
		return err
	}

	return nil
}
