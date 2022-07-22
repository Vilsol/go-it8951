package bmp

import (
	"encoding/binary"
	"errors"
	"image"
	"io"
)

type header struct {
	sigBM         [2]byte
	fileSize      uint32
	pixOffset     uint32
	dibHeaderSize uint32
	width         uint32
	height        uint32
	colorPlane    uint16
	bpp           uint16
	imageSize     uint32
}

func encodePalletted(w io.Writer, pix []uint8, dx, dy, stride, step int) error {
	var padding []byte
	if dx < step {
		padding = make([]byte, step-dx)
	}
	for y := dy - 1; y >= 0; y-- {
		min := y*stride + 0
		max := y*stride + dx
		if _, err := w.Write(to4bpp(pix[min:max])); err != nil {
			return err
		}
		if padding != nil {
			if _, err := w.Write(padding); err != nil {
				return err
			}
		}
	}
	return nil
}

func to4bpp(in []byte) []byte {
	out := make([]byte, len(in)/2)
	for i := 0; i < len(in)/2; i++ {
		left := (in[i*2+1] / 16) << 4
		right := in[i*2] / 16
		out[i] = left | right
	}
	return out
}

func Encode(w io.Writer, m image.Image) error {
	d := m.Bounds().Size()
	if d.X < 0 || d.Y < 0 {
		return errors.New("bmp: negative bounds")
	}
	h := &header{
		sigBM:         [2]byte{'B', 'M'},
		fileSize:      14 + 40,
		pixOffset:     14 + 40,
		dibHeaderSize: 40,
		width:         uint32(d.X),
		height:        uint32(d.Y),
		colorPlane:    1,
	}

	var step int
	var palette []byte
	switch m.(type) {
	case *image.Gray:
		step = (d.X + 3) &^ 3
		palette = make([]byte, 1024)
		for i := 0; i < 256; i++ {
			palette[i*4+0] = uint8(i)
			palette[i*4+1] = uint8(i)
			palette[i*4+2] = uint8(i)
			palette[i*4+3] = 0xFF
		}
		h.imageSize = uint32(d.Y * step)
		h.fileSize += uint32(len(palette)) + h.imageSize
		h.pixOffset += uint32(len(palette))
		h.bpp = 4
	default:
		panic("unknown image type")
	}

	if err := binary.Write(w, binary.LittleEndian, h); err != nil {
		return err
	}
	if palette != nil {
		if err := binary.Write(w, binary.LittleEndian, palette); err != nil {
			return err
		}
	}

	if d.X == 0 || d.Y == 0 {
		return nil
	}

	switch m := m.(type) {
	case *image.Gray:
		return encodePalletted(w, m.Pix, d.X, d.Y, m.Stride, step)
	}

	panic("unknown image type")
}
