package main

import (
	"math"

	"go-it8951/all"
)

const vcomRaw = -2.21
const epd_mode = 0

func main() {
	if all.DEV_Module_Init() != 0 {
		panic("failed initializing dev module")
	}

	VCOM := math.Abs(vcomRaw) * 1000

	devInfo := all.EPD_IT8951_Init(uint16(VCOM))
	devInfo.Deref()

	println("VCOM:", VCOM)
	println("Display mode:", epd_mode)
	println("W:", devInfo.Panel_W)
	println("H:", devInfo.Panel_H)

	println("FW Version:", string(uint16ToByteArray(devInfo.FW_Version[:])))
	println("LUT Version:", string(uint16ToByteArray(devInfo.LUT_Version[:])))

	println("Low", devInfo.Memory_Addr_L)
	println("High", devInfo.Memory_Addr_H)

	Init_Target_Memory_Addr := (uint32(devInfo.Memory_Addr_L) & 0x0000FFFF) | ((uint32(devInfo.Memory_Addr_H) << 16) & 0xFFFF0000)
	//Init_Target_Memory_Addr = 122480

	println("Memory Address:", Init_Target_Memory_Addr)

	all.EPD_IT8951_Clear_Refresh(devInfo, Init_Target_Memory_Addr, 0) // INIT_Mode

	Display_ColorPalette_Example(devInfo, devInfo.Panel_W, devInfo.Panel_H, Init_Target_Memory_Addr)
}

func uint16ToByteArray(arr []uint16) []byte {
	out := make([]byte, len(arr)*2)
	for i, u := range arr {
		out[i*2] = byte(u & 0x00FF)
		out[i*2+1] = byte(u >> 8)
	}
	return out
}
