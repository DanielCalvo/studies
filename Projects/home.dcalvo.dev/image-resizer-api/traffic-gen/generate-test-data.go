package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

type fixture struct {
	width  int
	height int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: generate-test-data <directory>")
		os.Exit(2)
	}

	if err := generate(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generate(directory string) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create test-data directory: %w", err)
	}

	fixtures := []fixture{
		{width: 640, height: 480},
		{width: 800, height: 1_200},
		{width: 1_200, height: 1_200},
		{width: 1_600, height: 1_200},
		{width: 2_000, height: 1_500},
		{width: 2_400, height: 1_800},
		{width: 2_000, height: 3_000},
		{width: 3_200, height: 2_400},
		{width: 4_000, height: 3_000},
		{width: 5_000, height: 5_000},
	}
	for _, item := range fixtures {
		name := fmt.Sprintf("%dx%d.jpg", item.width, item.height)
		if err := generateJPEG(filepath.Join(directory, name), item.width, item.height); err != nil {
			return err
		}
	}
	return nil
}

func generateJPEG(path string, width, height int) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			offset := y*img.Stride + x*4
			img.Pix[offset] = uint8(x * 255 / width)
			img.Pix[offset+1] = uint8(y * 255 / height)
			img.Pix[offset+2] = uint8((x/64 + y/64) * 17)
			img.Pix[offset+3] = 255
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 85}); err != nil {
		_ = file.Close()
		return fmt.Errorf("encode %s: %w", path, err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close %s: %w", path, err)
	}
	return nil
}
