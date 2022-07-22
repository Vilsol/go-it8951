package main

import (
	"bytes"
	"image"
	"image/color"

	"golang.org/x/image/draw"

	"github.com/disintegration/imaging"

	"go-it8951/all"
	"go-it8951/bmp"
)

const leftWidth = 660

func Display_ColorPalette_Example(devInfo all.IT8951_Dev_Info, Panel_Width uint16, Panel_Height uint16, Init_Target_Memory_Addr uint32) {
	all.EPD_IT8951_Clear_Refresh(devInfo, Init_Target_Memory_Addr, 2) // GC16_Mode

	//In_4bp_Refresh_Area_Width := Panel_Width
	//In_4bp_Refresh_Area_Height := Panel_Height / 16
	//
	//ctx := gg.NewContext(int(Panel_Width), int(Panel_Height))
	//for i := uint16(0); i < 16; i++ {
	//	var c = int(255 - uint8(i*16))
	//	ctx.SetRGB255(c, c, c)
	//	ctx.DrawRectangle(0, float64(i*In_4bp_Refresh_Area_Height), float64(In_4bp_Refresh_Area_Width), float64(In_4bp_Refresh_Area_Height))
	//	ctx.Fill()
	//}
	//
	//if err := ctx.LoadFontFace("./impact.ttf", 96); err != nil {
	//	panic(err)
	//}
	//
	//ctx.DrawStringAnchored("Hello, world!", float64(Panel_Width/2), float64(Panel_Height/2), 0.5, 0.5)
	//
	//panel := Panel{
	//	Width:   Panel_Width,
	//	Height:  Panel_Height,
	//	Address: Init_Target_Memory_Addr,
	//}
	//
	//panel.DrawImage(ctx.Image(), 0, 0, Panel_Width, Panel_Height)

	sizes := []int{25, 50, 75, 100, 145, 200}

	offset := 0
	for i := uint16(0); i < uint16(len(sizes)); i++ {
		square := imaging.New(sizes[i], sizes[i], color.Black)
		img := image.NewGray(square.Rect)
		draw.Draw(img, img.Bounds(), square, image.Point{}, draw.Src)

		buff := new(bytes.Buffer)
		if err := bmp.Encode(buff, img); err != nil {
			panic(err)
		}

		maxFit := (Panel_Width - 100) / uint16(sizes[i]+20)

		for j := uint16(0); j < maxFit; j++ {
			all.EPD_IT8951_4bp_Refresh(buff.Bytes(), 50+(j*(uint16(img.Bounds().Dx())+20)), uint16(50+offset), uint16(img.Bounds().Dx()), uint16(img.Bounds().Dy()), false, Init_Target_Memory_Addr, false)
		}

		offset += sizes[i] + 50
	}

	all.EPD_IT8951_Clear_Refresh(devInfo, Init_Target_Memory_Addr, 2) // GC16_Mode

	sizes = []int{25, 50, 75, 100, 145, 200, 250}

	offset = 0
	for i := uint16(0); i < uint16(len(sizes)); i++ {
		square := imaging.New(sizes[i], sizes[i], color.Black)
		img := image.NewGray(square.Rect)
		draw.Draw(img, img.Bounds(), square, image.Point{}, draw.Src)

		buff := new(bytes.Buffer)
		if err := bmp.Encode(buff, img); err != nil {
			panic(err)
		}

		maxFit := (Panel_Height - 100) / uint16(sizes[i]+20)

		for j := uint16(0); j < maxFit; j++ {
			all.EPD_IT8951_4bp_Refresh(buff.Bytes(), uint16(50+offset), 50+(j*(uint16(img.Bounds().Dx())+20)), uint16(img.Bounds().Dx()), uint16(img.Bounds().Dy()), false, Init_Target_Memory_Addr, false)
		}

		offset += sizes[i] + 50
	}
}
