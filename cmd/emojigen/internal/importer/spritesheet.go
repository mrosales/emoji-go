package importer

import (
	"fmt"
	"image"
	"image/png"
	"io"
)

// SpriteSheet represents the image data for a sprite sheet.
type SpriteSheet struct {
	itemWidth int
	rgbaImage *image.NRGBA
}

// NewSpriteSheet parses a PNG sprite of fixed-width sprites
func NewSpriteSheet(reader io.Reader, width int) (*SpriteSheet, error) {
	pngImage, err := png.Decode(reader)
	if err != nil {
		return nil, fmt.Errorf("failed decoding spite png: %w", err)
	}
	rgbaImage, ok := pngImage.(*image.NRGBA)
	if !ok {
		return nil, fmt.Errorf("image was not a valid RGBA png: %w, %T", err, pngImage)
	}
	return &SpriteSheet{
		itemWidth: width,
		rgbaImage: rgbaImage,
	}, nil
}

// Get retrieves an image from the sprite sheet.
func (s SpriteSheet) Get(x, y int) *image.NRGBA {
	// 1px padding on each side and initial offset of 1
	xOffset := 1 + (1+s.itemWidth+1)*x
	yOffset := 1 + (1+s.itemWidth+1)*y

	rect := image.Rect(xOffset, yOffset, xOffset+s.itemWidth, yOffset+s.itemWidth)
	return s.rgbaImage.SubImage(rect).(*image.NRGBA)
}
