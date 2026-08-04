package main

import (
	"image"

	"golang.org/x/image/draw"
)

func resizeWidth(source image.Image, width int) image.Image {
	bounds := source.Bounds()
	height := scaledHeight(bounds.Dx(), bounds.Dy(), width)
	destination := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.ApproxBiLinear.Scale(destination, destination.Bounds(), source, bounds, draw.Src, nil)
	return destination
}

func scaledHeight(sourceWidth, sourceHeight, targetWidth int) int {
	height := (int64(sourceHeight)*int64(targetWidth) + int64(sourceWidth)/2) / int64(sourceWidth)
	if height < 1 {
		return 1
	}
	return int(height)
}
