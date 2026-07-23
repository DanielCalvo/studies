package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const oversizedUploadBytes = (10 << 20) + 1

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: fixture-tool generate <directory> | dimensions <jpeg>")
		os.Exit(2)
	}

	var err error
	switch os.Args[1] {
	case "generate":
		err = generateFixtures(os.Args[2])
	case "dimensions":
		err = printJPEGDimensions(os.Args[2])
	default:
		err = fmt.Errorf("unknown command %q", os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func generateFixtures(directory string) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create test-data directory: %w", err)
	}

	source := image.NewRGBA(image.Rect(0, 0, 120, 80))
	for y := range 80 {
		for x := range 120 {
			source.SetRGBA(x, y, color.RGBA{
				R: uint8(x * 2),
				G: uint8(y * 3),
				B: 140,
				A: 255,
			})
		}
	}

	var jpegData bytes.Buffer
	if err := jpeg.Encode(&jpegData, source, &jpeg.Options{Quality: 90}); err != nil {
		return fmt.Errorf("encode JPEG fixture: %w", err)
	}
	if err := writeFixture(directory, "valid.jpg", jpegData.Bytes()); err != nil {
		return err
	}

	var pngData bytes.Buffer
	if err := png.Encode(&pngData, source); err != nil {
		return fmt.Errorf("encode PNG fixture: %w", err)
	}
	if err := writeFixture(directory, "valid.png", pngData.Bytes()); err != nil {
		return err
	}

	tooManyPixels := append([]byte(nil), jpegData.Bytes()...)
	if err := replaceJPEGDimensions(tooManyPixels, 7_000, 6_000); err != nil {
		return err
	}
	if err := writeFixture(directory, "too-many-pixels.jpg", tooManyPixels); err != nil {
		return err
	}

	if err := writeFixture(directory, "corrupt.jpg", []byte{0xff, 0xd8, 0xff, 0xdb, 0x00}); err != nil {
		return err
	}
	if err := writeFixture(directory, "plain-text.txt", []byte("this is not an image\n")); err != nil {
		return err
	}
	if err := writeFixture(directory, "video.mp4", minimalMP4()); err != nil {
		return err
	}

	oversizedPath := filepath.Join(directory, "oversized.jpg")
	oversized, err := os.Create(oversizedPath)
	if err != nil {
		return fmt.Errorf("create %s: %w", oversizedPath, err)
	}
	if _, err := oversized.Write(jpegData.Bytes()); err != nil {
		_ = oversized.Close()
		return fmt.Errorf("write %s: %w", oversizedPath, err)
	}
	if err := oversized.Truncate(oversizedUploadBytes); err != nil {
		_ = oversized.Close()
		return fmt.Errorf("resize %s: %w", oversizedPath, err)
	}
	if err := oversized.Close(); err != nil {
		return fmt.Errorf("close %s: %w", oversizedPath, err)
	}

	return nil
}

func replaceJPEGDimensions(data []byte, width, height uint16) error {
	for index := 0; index+8 < len(data); index++ {
		if data[index] != 0xff || (data[index+1] != 0xc0 && data[index+1] != 0xc2) {
			continue
		}
		binary.BigEndian.PutUint16(data[index+5:index+7], height)
		binary.BigEndian.PutUint16(data[index+7:index+9], width)
		return nil
	}
	return fmt.Errorf("JPEG fixture has no supported start-of-frame marker")
}

func minimalMP4() []byte {
	return []byte{
		0x00, 0x00, 0x00, 0x18,
		'f', 't', 'y', 'p',
		'i', 's', 'o', 'm',
		0x00, 0x00, 0x02, 0x00,
		'i', 's', 'o', 'm',
		'i', 's', 'o', '2',
	}
}

func writeFixture(directory, name string, data []byte) error {
	path := filepath.Join(directory, name)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func printJPEGDimensions(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()

	config, err := jpeg.DecodeConfig(file)
	if err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	fmt.Printf("%dx%d\n", config.Width, config.Height)
	return nil
}
