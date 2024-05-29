package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	_fontPath = "Roboto-Bold.ttf"
)

var (
	_textColor       = color.RGBA{255, 255, 255, 255}
	_backgroundColor = color.RGBA{46, 26, 71, 255}
)

type ParamsCreateImage struct {
	PathFont  string
	PathImage string

	PixelWidth  uint
	PixelHeight uint

	ColorText       color.RGBA
	ColorBackground color.RGBA
}

func loadFont(fontPath string) (*truetype.Font, error) {
	fontBytes, errReadFile := os.ReadFile(fontPath)
	if errReadFile != nil {
		return nil,
			errReadFile
	}

	return freetype.ParseFont(fontBytes)
}

func CreateImage(params *ParamsCreateImage) error {
	font, errLoadFont := loadFont(params.PathFont)
	if errLoadFont != nil {
		return errLoadFont
	}

	img := image.NewRGBA(
		image.Rect(
			0,
			0,
			int(params.PixelWidth),
			int(params.PixelHeight),
		),
	)

	draw.Draw(
		img,
		img.Bounds(),
		&image.Uniform{_backgroundColor},
		image.Point{},
		draw.Src,
	)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(48)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(_textColor))

	text := fmt.Sprintf(
		"%dx%d",
		params.PixelWidth,
		params.PixelHeight,
	)

	textExtent, errDrawTextExtent := c.DrawString(
		text,
		freetype.Pt(0, 0),
	)
	if errDrawTextExtent != nil {
		return errDrawTextExtent
	}

	textWidth := int(textExtent.X >> 6)

	x := (int(params.PixelWidth) - textWidth) / 2
	y := (int(params.PixelHeight) + int(c.PointToFixed(48)>>6)) / 2

	if _, errDrawText := c.DrawString(text, freetype.Pt(x, y)); errDrawText != nil {
		return errDrawText
	}

	f, errCreateFile := os.Create(params.PathImage)
	if errCreateFile != nil {
		return errCreateFile
	}
	defer f.Close()

	return png.Encode(f, img)
}
