package main

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
	"golang.org/x/image/draw"

	"go-it8951/all"
	"go-it8951/bmp"
)

type Panel struct {
	Width   uint16
	Height  uint16
	Address uint32
}

func (p Panel) DrawImage(out image.Image, x uint16, y uint16, w uint16, h uint16) {
	flipped := imaging.FlipV(out)

	left := imaging.CropAnchor(flipped, leftWidth, int(p.Height), imaging.TopLeft)
	right := imaging.CropAnchor(flipped, int(p.Width)-leftWidth, int(p.Height), imaging.TopRight)

	img := image.NewGray(image.Rect(0, 0, int(p.Width), int(p.Height)))

	draw.Draw(img, right.Bounds(), right, image.Point{}, draw.Src)
	draw.Draw(img, image.Rect(right.Bounds().Dx(), 0, img.Bounds().Dx(), img.Bounds().Dy()), left, image.Point{}, draw.Src)

	buff := new(bytes.Buffer)
	if err := bmp.Encode(buff, img); err != nil {
		panic(err)
	}

	all.EPD_IT8951_4bp_Refresh(buff.Bytes(), x, y, w, h, false, p.Address, false)
}
