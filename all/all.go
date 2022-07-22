// WARNING: This file has automatically been generated on Fri, 22 Jul 2022 22:52:51 EEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package all

/*
#include "DEV_Config.h"
#include "bcm2835.h"
#include "EPD_IT8951.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// DEV_Digital_Write function as declared in all/DEV_Config.h:83
func DEV_Digital_Write(pin uint16, value byte) {
	cpin, cpinAllocMap := (C.uint16_t)(pin), cgoAllocsUnknown
	cvalue, cvalueAllocMap := (C.uint8_t)(value), cgoAllocsUnknown
	C.DEV_Digital_Write(cpin, cvalue)
	runtime.KeepAlive(cvalueAllocMap)
	runtime.KeepAlive(cpinAllocMap)
}

// DEV_Digital_Read function as declared in all/DEV_Config.h:84
func DEV_Digital_Read(pin uint16) byte {
	cpin, cpinAllocMap := (C.uint16_t)(pin), cgoAllocsUnknown
	__ret := C.DEV_Digital_Read(cpin)
	runtime.KeepAlive(cpinAllocMap)
	__v := (byte)(__ret)
	return __v
}

// DEV_SPI_WriteByte function as declared in all/DEV_Config.h:86
func DEV_SPI_WriteByte(value byte) {
	cvalue, cvalueAllocMap := (C.uint8_t)(value), cgoAllocsUnknown
	C.DEV_SPI_WriteByte(cvalue)
	runtime.KeepAlive(cvalueAllocMap)
}

// DEV_SPI_ReadByte function as declared in all/DEV_Config.h:87
func DEV_SPI_ReadByte() byte {
	__ret := C.DEV_SPI_ReadByte()
	__v := (byte)(__ret)
	return __v
}

// DEV_Delay_ms function as declared in all/DEV_Config.h:89
func DEV_Delay_ms(xms uint32) {
	cxms, cxmsAllocMap := (C.uint32_t)(xms), cgoAllocsUnknown
	C.DEV_Delay_ms(cxms)
	runtime.KeepAlive(cxmsAllocMap)
}

// DEV_Delay_us function as declared in all/DEV_Config.h:90
func DEV_Delay_us(xus uint32) {
	cxus, cxusAllocMap := (C.uint32_t)(xus), cgoAllocsUnknown
	C.DEV_Delay_us(cxus)
	runtime.KeepAlive(cxusAllocMap)
}

// DEV_Module_Init function as declared in all/DEV_Config.h:92
func DEV_Module_Init() byte {
	__ret := C.DEV_Module_Init()
	__v := (byte)(__ret)
	return __v
}

// DEV_Module_Exit function as declared in all/DEV_Config.h:93
func DEV_Module_Exit() {
	C.DEV_Module_Exit()
}

// bcm2835_init function as declared in all/bcm2835.h:1390
func bcm2835_init() int32 {
	__ret := C.bcm2835_init()
	__v := (int32)(__ret)
	return __v
}

// bcm2835_close function as declared in all/bcm2835.h:1395
func bcm2835_close() int32 {
	__ret := C.bcm2835_close()
	__v := (int32)(__ret)
	return __v
}

// bcm2835_set_debug function as declared in all/bcm2835.h:1404
func bcm2835_set_debug(debug byte) {
	cdebug, cdebugAllocMap := (C.uint8_t)(debug), cgoAllocsUnknown
	C.bcm2835_set_debug(cdebug)
	runtime.KeepAlive(cdebugAllocMap)
}

// bcm2835_version function as declared in all/bcm2835.h:1409
func bcm2835_version() uint32 {
	__ret := C.bcm2835_version()
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_regbase function as declared in all/bcm2835.h:1426
func bcm2835_regbase(regbase byte) *uint32 {
	cregbase, cregbaseAllocMap := (C.uint8_t)(regbase), cgoAllocsUnknown
	__ret := C.bcm2835_regbase(cregbase)
	runtime.KeepAlive(cregbaseAllocMap)
	__v := *(**uint32)(unsafe.Pointer(&__ret))
	return __v
}

// bcm2835_peri_read function as declared in all/bcm2835.h:1436
func bcm2835_peri_read(paddr []uint32) uint32 {
	cpaddr, cpaddrAllocMap := copyPUint32_tBytes((*sliceHeader)(unsafe.Pointer(&paddr)))
	__ret := C.bcm2835_peri_read(cpaddr)
	runtime.KeepAlive(cpaddrAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_peri_read_nb function as declared in all/bcm2835.h:1448
func bcm2835_peri_read_nb(paddr []uint32) uint32 {
	cpaddr, cpaddrAllocMap := copyPUint32_tBytes((*sliceHeader)(unsafe.Pointer(&paddr)))
	__ret := C.bcm2835_peri_read_nb(cpaddr)
	runtime.KeepAlive(cpaddrAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_peri_write function as declared in all/bcm2835.h:1462
func bcm2835_peri_write(paddr []uint32, value uint32) {
	cpaddr, cpaddrAllocMap := copyPUint32_tBytes((*sliceHeader)(unsafe.Pointer(&paddr)))
	cvalue, cvalueAllocMap := (C.uint32_t)(value), cgoAllocsUnknown
	C.bcm2835_peri_write(cpaddr, cvalue)
	runtime.KeepAlive(cvalueAllocMap)
	runtime.KeepAlive(cpaddrAllocMap)
}

// bcm2835_peri_write_nb function as declared in all/bcm2835.h:1476
func bcm2835_peri_write_nb(paddr []uint32, value uint32) {
	cpaddr, cpaddrAllocMap := copyPUint32_tBytes((*sliceHeader)(unsafe.Pointer(&paddr)))
	cvalue, cvalueAllocMap := (C.uint32_t)(value), cgoAllocsUnknown
	C.bcm2835_peri_write_nb(cpaddr, cvalue)
	runtime.KeepAlive(cvalueAllocMap)
	runtime.KeepAlive(cpaddrAllocMap)
}

// bcm2835_peri_set_bits function as declared in all/bcm2835.h:1490
func bcm2835_peri_set_bits(paddr []uint32, value uint32, mask uint32) {
	cpaddr, cpaddrAllocMap := copyPUint32_tBytes((*sliceHeader)(unsafe.Pointer(&paddr)))
	cvalue, cvalueAllocMap := (C.uint32_t)(value), cgoAllocsUnknown
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	C.bcm2835_peri_set_bits(cpaddr, cvalue, cmask)
	runtime.KeepAlive(cmaskAllocMap)
	runtime.KeepAlive(cvalueAllocMap)
	runtime.KeepAlive(cpaddrAllocMap)
}

// bcm2835_gpio_fsel function as declared in all/bcm2835.h:1504
func bcm2835_gpio_fsel(pin byte, mode byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	cmode, cmodeAllocMap := (C.uint8_t)(mode), cgoAllocsUnknown
	C.bcm2835_gpio_fsel(cpin, cmode)
	runtime.KeepAlive(cmodeAllocMap)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_set function as declared in all/bcm2835.h:1511
func bcm2835_gpio_set(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_set(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr function as declared in all/bcm2835.h:1518
func bcm2835_gpio_clr(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_set_multi function as declared in all/bcm2835.h:1525
func bcm2835_gpio_set_multi(mask uint32) {
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	C.bcm2835_gpio_set_multi(cmask)
	runtime.KeepAlive(cmaskAllocMap)
}

// bcm2835_gpio_clr_multi function as declared in all/bcm2835.h:1532
func bcm2835_gpio_clr_multi(mask uint32) {
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	C.bcm2835_gpio_clr_multi(cmask)
	runtime.KeepAlive(cmaskAllocMap)
}

// bcm2835_gpio_lev function as declared in all/bcm2835.h:1540
func bcm2835_gpio_lev(pin byte) byte {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	__ret := C.bcm2835_gpio_lev(cpin)
	runtime.KeepAlive(cpinAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_gpio_eds function as declared in all/bcm2835.h:1550
func bcm2835_gpio_eds(pin byte) byte {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	__ret := C.bcm2835_gpio_eds(cpin)
	runtime.KeepAlive(cpinAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_gpio_eds_multi function as declared in all/bcm2835.h:1557
func bcm2835_gpio_eds_multi(mask uint32) uint32 {
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	__ret := C.bcm2835_gpio_eds_multi(cmask)
	runtime.KeepAlive(cmaskAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_gpio_set_eds function as declared in all/bcm2835.h:1564
func bcm2835_gpio_set_eds(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_set_eds(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_set_eds_multi function as declared in all/bcm2835.h:1570
func bcm2835_gpio_set_eds_multi(mask uint32) {
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	C.bcm2835_gpio_set_eds_multi(cmask)
	runtime.KeepAlive(cmaskAllocMap)
}

// bcm2835_gpio_ren function as declared in all/bcm2835.h:1580
func bcm2835_gpio_ren(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_ren(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr_ren function as declared in all/bcm2835.h:1585
func bcm2835_gpio_clr_ren(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr_ren(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_fen function as declared in all/bcm2835.h:1595
func bcm2835_gpio_fen(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_fen(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr_fen function as declared in all/bcm2835.h:1600
func bcm2835_gpio_clr_fen(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr_fen(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_hen function as declared in all/bcm2835.h:1606
func bcm2835_gpio_hen(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_hen(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr_hen function as declared in all/bcm2835.h:1611
func bcm2835_gpio_clr_hen(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr_hen(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_len function as declared in all/bcm2835.h:1617
func bcm2835_gpio_len(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_len(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr_len function as declared in all/bcm2835.h:1622
func bcm2835_gpio_clr_len(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr_len(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_aren function as declared in all/bcm2835.h:1630
func bcm2835_gpio_aren(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_aren(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr_aren function as declared in all/bcm2835.h:1635
func bcm2835_gpio_clr_aren(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr_aren(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_afen function as declared in all/bcm2835.h:1643
func bcm2835_gpio_afen(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_afen(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_clr_afen function as declared in all/bcm2835.h:1648
func bcm2835_gpio_clr_afen(pin byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	C.bcm2835_gpio_clr_afen(cpin)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_pud function as declared in all/bcm2835.h:1658
func bcm2835_gpio_pud(pud byte) {
	cpud, cpudAllocMap := (C.uint8_t)(pud), cgoAllocsUnknown
	C.bcm2835_gpio_pud(cpud)
	runtime.KeepAlive(cpudAllocMap)
}

// bcm2835_gpio_pudclk function as declared in all/bcm2835.h:1670
func bcm2835_gpio_pudclk(pin byte, on byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	con, conAllocMap := (C.uint8_t)(on), cgoAllocsUnknown
	C.bcm2835_gpio_pudclk(cpin, con)
	runtime.KeepAlive(conAllocMap)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_pad function as declared in all/bcm2835.h:1677
func bcm2835_gpio_pad(group byte) uint32 {
	cgroup, cgroupAllocMap := (C.uint8_t)(group), cgoAllocsUnknown
	__ret := C.bcm2835_gpio_pad(cgroup)
	runtime.KeepAlive(cgroupAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_gpio_set_pad function as declared in all/bcm2835.h:1686
func bcm2835_gpio_set_pad(group byte, control uint32) {
	cgroup, cgroupAllocMap := (C.uint8_t)(group), cgoAllocsUnknown
	ccontrol, ccontrolAllocMap := (C.uint32_t)(control), cgoAllocsUnknown
	C.bcm2835_gpio_set_pad(cgroup, ccontrol)
	runtime.KeepAlive(ccontrolAllocMap)
	runtime.KeepAlive(cgroupAllocMap)
}

// bcm2835_delay function as declared in all/bcm2835.h:1698
func bcm2835_delay(millis uint32) {
	cmillis, cmillisAllocMap := (C.uint)(millis), cgoAllocsUnknown
	C.bcm2835_delay(cmillis)
	runtime.KeepAlive(cmillisAllocMap)
}

// bcm2835_delayMicroseconds function as declared in all/bcm2835.h:1713
func bcm2835_delayMicroseconds(micros uint32) {
	cmicros, cmicrosAllocMap := (C.uint64_t)(micros), cgoAllocsUnknown
	C.bcm2835_delayMicroseconds(cmicros)
	runtime.KeepAlive(cmicrosAllocMap)
}

// bcm2835_gpio_write function as declared in all/bcm2835.h:1719
func bcm2835_gpio_write(pin byte, on byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	con, conAllocMap := (C.uint8_t)(on), cgoAllocsUnknown
	C.bcm2835_gpio_write(cpin, con)
	runtime.KeepAlive(conAllocMap)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_write_multi function as declared in all/bcm2835.h:1725
func bcm2835_gpio_write_multi(mask uint32, on byte) {
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	con, conAllocMap := (C.uint8_t)(on), cgoAllocsUnknown
	C.bcm2835_gpio_write_multi(cmask, con)
	runtime.KeepAlive(conAllocMap)
	runtime.KeepAlive(cmaskAllocMap)
}

// bcm2835_gpio_write_mask function as declared in all/bcm2835.h:1731
func bcm2835_gpio_write_mask(value uint32, mask uint32) {
	cvalue, cvalueAllocMap := (C.uint32_t)(value), cgoAllocsUnknown
	cmask, cmaskAllocMap := (C.uint32_t)(mask), cgoAllocsUnknown
	C.bcm2835_gpio_write_mask(cvalue, cmask)
	runtime.KeepAlive(cmaskAllocMap)
	runtime.KeepAlive(cvalueAllocMap)
}

// bcm2835_gpio_set_pud function as declared in all/bcm2835.h:1738
func bcm2835_gpio_set_pud(pin byte, pud byte) {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	cpud, cpudAllocMap := (C.uint8_t)(pud), cgoAllocsUnknown
	C.bcm2835_gpio_set_pud(cpin, cpud)
	runtime.KeepAlive(cpudAllocMap)
	runtime.KeepAlive(cpinAllocMap)
}

// bcm2835_gpio_get_pud function as declared in all/bcm2835.h:1746
func bcm2835_gpio_get_pud(pin byte) byte {
	cpin, cpinAllocMap := (C.uint8_t)(pin), cgoAllocsUnknown
	__ret := C.bcm2835_gpio_get_pud(cpin)
	runtime.KeepAlive(cpinAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_spi_begin function as declared in all/bcm2835.h:1764
func bcm2835_spi_begin() int32 {
	__ret := C.bcm2835_spi_begin()
	__v := (int32)(__ret)
	return __v
}

// bcm2835_spi_end function as declared in all/bcm2835.h:1770
func bcm2835_spi_end() {
	C.bcm2835_spi_end()
}

// bcm2835_spi_setBitOrder function as declared in all/bcm2835.h:1779
func bcm2835_spi_setBitOrder(order byte) {
	corder, corderAllocMap := (C.uint8_t)(order), cgoAllocsUnknown
	C.bcm2835_spi_setBitOrder(corder)
	runtime.KeepAlive(corderAllocMap)
}

// bcm2835_spi_setClockDivider function as declared in all/bcm2835.h:1786
func bcm2835_spi_setClockDivider(divider uint16) {
	cdivider, cdividerAllocMap := (C.uint16_t)(divider), cgoAllocsUnknown
	C.bcm2835_spi_setClockDivider(cdivider)
	runtime.KeepAlive(cdividerAllocMap)
}

// bcm2835_spi_set_speed_hz function as declared in all/bcm2835.h:1792
func bcm2835_spi_set_speed_hz(speed_hz uint32) {
	cspeed_hz, cspeed_hzAllocMap := (C.uint32_t)(speed_hz), cgoAllocsUnknown
	C.bcm2835_spi_set_speed_hz(cspeed_hz)
	runtime.KeepAlive(cspeed_hzAllocMap)
}

// bcm2835_spi_setDataMode function as declared in all/bcm2835.h:1799
func bcm2835_spi_setDataMode(mode byte) {
	cmode, cmodeAllocMap := (C.uint8_t)(mode), cgoAllocsUnknown
	C.bcm2835_spi_setDataMode(cmode)
	runtime.KeepAlive(cmodeAllocMap)
}

// bcm2835_spi_chipSelect function as declared in all/bcm2835.h:1807
func bcm2835_spi_chipSelect(cs byte) {
	ccs, ccsAllocMap := (C.uint8_t)(cs), cgoAllocsUnknown
	C.bcm2835_spi_chipSelect(ccs)
	runtime.KeepAlive(ccsAllocMap)
}

// bcm2835_spi_setChipSelectPolarity function as declared in all/bcm2835.h:1817
func bcm2835_spi_setChipSelectPolarity(cs byte, active byte) {
	ccs, ccsAllocMap := (C.uint8_t)(cs), cgoAllocsUnknown
	cactive, cactiveAllocMap := (C.uint8_t)(active), cgoAllocsUnknown
	C.bcm2835_spi_setChipSelectPolarity(ccs, cactive)
	runtime.KeepAlive(cactiveAllocMap)
	runtime.KeepAlive(ccsAllocMap)
}

// bcm2835_spi_transfer function as declared in all/bcm2835.h:1829
func bcm2835_spi_transfer(value byte) byte {
	cvalue, cvalueAllocMap := (C.uint8_t)(value), cgoAllocsUnknown
	__ret := C.bcm2835_spi_transfer(cvalue)
	runtime.KeepAlive(cvalueAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_spi_transfernb function as declared in all/bcm2835.h:1842
func bcm2835_spi_transfernb(tbuf []byte, rbuf []byte, len uint32) {
	ctbuf, ctbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&tbuf)))
	crbuf, crbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&rbuf)))
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	C.bcm2835_spi_transfernb(ctbuf, crbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(crbufAllocMap)
	runtime.KeepAlive(ctbufAllocMap)
}

// bcm2835_spi_transfern function as declared in all/bcm2835.h:1851
func bcm2835_spi_transfern(buf []byte, len uint32) {
	cbuf, cbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&buf)))
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	C.bcm2835_spi_transfern(cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
}

// bcm2835_spi_writenb function as declared in all/bcm2835.h:1859
func bcm2835_spi_writenb(buf string, len uint32) {
	cbuf, cbufAllocMap := unpackPCharString(buf)
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	C.bcm2835_spi_writenb(cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
}

// bcm2835_spi_write function as declared in all/bcm2835.h:1869
func bcm2835_spi_write(data uint16) {
	cdata, cdataAllocMap := (C.uint16_t)(data), cgoAllocsUnknown
	C.bcm2835_spi_write(cdata)
	runtime.KeepAlive(cdataAllocMap)
}

// bcm2835_aux_spi_begin function as declared in all/bcm2835.h:1876
func bcm2835_aux_spi_begin() int32 {
	__ret := C.bcm2835_aux_spi_begin()
	__v := (int32)(__ret)
	return __v
}

// bcm2835_aux_spi_end function as declared in all/bcm2835.h:1882
func bcm2835_aux_spi_end() {
	C.bcm2835_aux_spi_end()
}

// bcm2835_aux_spi_setClockDivider function as declared in all/bcm2835.h:1887
func bcm2835_aux_spi_setClockDivider(divider uint16) {
	cdivider, cdividerAllocMap := (C.uint16_t)(divider), cgoAllocsUnknown
	C.bcm2835_aux_spi_setClockDivider(cdivider)
	runtime.KeepAlive(cdividerAllocMap)
}

// bcm2835_aux_spi_CalcClockDivider function as declared in all/bcm2835.h:1894
func bcm2835_aux_spi_CalcClockDivider(speed_hz uint32) uint16 {
	cspeed_hz, cspeed_hzAllocMap := (C.uint32_t)(speed_hz), cgoAllocsUnknown
	__ret := C.bcm2835_aux_spi_CalcClockDivider(cspeed_hz)
	runtime.KeepAlive(cspeed_hzAllocMap)
	__v := (uint16)(__ret)
	return __v
}

// bcm2835_aux_spi_write function as declared in all/bcm2835.h:1902
func bcm2835_aux_spi_write(data uint16) {
	cdata, cdataAllocMap := (C.uint16_t)(data), cgoAllocsUnknown
	C.bcm2835_aux_spi_write(cdata)
	runtime.KeepAlive(cdataAllocMap)
}

// bcm2835_aux_spi_writenb function as declared in all/bcm2835.h:1909
func bcm2835_aux_spi_writenb(buf string, len uint32) {
	cbuf, cbufAllocMap := unpackPCharString(buf)
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	C.bcm2835_aux_spi_writenb(cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
}

// bcm2835_aux_spi_transfern function as declared in all/bcm2835.h:1918
func bcm2835_aux_spi_transfern(buf []byte, len uint32) {
	cbuf, cbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&buf)))
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	C.bcm2835_aux_spi_transfern(cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
}

// bcm2835_aux_spi_transfernb function as declared in all/bcm2835.h:1928
func bcm2835_aux_spi_transfernb(tbuf string, rbuf []byte, len uint32) {
	ctbuf, ctbufAllocMap := unpackPCharString(tbuf)
	crbuf, crbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&rbuf)))
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	C.bcm2835_aux_spi_transfernb(ctbuf, crbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(crbufAllocMap)
	runtime.KeepAlive(ctbufAllocMap)
}

// bcm2835_aux_spi_transfer function as declared in all/bcm2835.h:1937
func bcm2835_aux_spi_transfer(value byte) byte {
	cvalue, cvalueAllocMap := (C.uint8_t)(value), cgoAllocsUnknown
	__ret := C.bcm2835_aux_spi_transfer(cvalue)
	runtime.KeepAlive(cvalueAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_i2c_begin function as declared in all/bcm2835.h:1955
func bcm2835_i2c_begin() int32 {
	__ret := C.bcm2835_i2c_begin()
	__v := (int32)(__ret)
	return __v
}

// bcm2835_i2c_end function as declared in all/bcm2835.h:1961
func bcm2835_i2c_end() {
	C.bcm2835_i2c_end()
}

// bcm2835_i2c_setSlaveAddress function as declared in all/bcm2835.h:1966
func bcm2835_i2c_setSlaveAddress(addr byte) {
	caddr, caddrAllocMap := (C.uint8_t)(addr), cgoAllocsUnknown
	C.bcm2835_i2c_setSlaveAddress(caddr)
	runtime.KeepAlive(caddrAllocMap)
}

// bcm2835_i2c_setClockDivider function as declared in all/bcm2835.h:1972
func bcm2835_i2c_setClockDivider(divider uint16) {
	cdivider, cdividerAllocMap := (C.uint16_t)(divider), cgoAllocsUnknown
	C.bcm2835_i2c_setClockDivider(cdivider)
	runtime.KeepAlive(cdividerAllocMap)
}

// bcm2835_i2c_set_baudrate function as declared in all/bcm2835.h:1980
func bcm2835_i2c_set_baudrate(baudrate uint32) {
	cbaudrate, cbaudrateAllocMap := (C.uint32_t)(baudrate), cgoAllocsUnknown
	C.bcm2835_i2c_set_baudrate(cbaudrate)
	runtime.KeepAlive(cbaudrateAllocMap)
}

// bcm2835_i2c_write function as declared in all/bcm2835.h:1988
func bcm2835_i2c_write(buf string, len uint32) byte {
	cbuf, cbufAllocMap := unpackPCharString(buf)
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	__ret := C.bcm2835_i2c_write(cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_i2c_read function as declared in all/bcm2835.h:1996
func bcm2835_i2c_read(buf []byte, len uint32) byte {
	cbuf, cbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&buf)))
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	__ret := C.bcm2835_i2c_read(cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_i2c_read_register_rs function as declared in all/bcm2835.h:2012
func bcm2835_i2c_read_register_rs(regaddr []byte, buf []byte, len uint32) byte {
	cregaddr, cregaddrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&regaddr)))
	cbuf, cbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&buf)))
	clen, clenAllocMap := (C.uint32_t)(len), cgoAllocsUnknown
	__ret := C.bcm2835_i2c_read_register_rs(cregaddr, cbuf, clen)
	runtime.KeepAlive(clenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
	runtime.KeepAlive(cregaddrAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_i2c_write_read_rs function as declared in all/bcm2835.h:2024
func bcm2835_i2c_write_read_rs(cmds []byte, cmds_len uint32, buf []byte, buf_len uint32) byte {
	ccmds, ccmdsAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&cmds)))
	ccmds_len, ccmds_lenAllocMap := (C.uint32_t)(cmds_len), cgoAllocsUnknown
	cbuf, cbufAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&buf)))
	cbuf_len, cbuf_lenAllocMap := (C.uint32_t)(buf_len), cgoAllocsUnknown
	__ret := C.bcm2835_i2c_write_read_rs(ccmds, ccmds_len, cbuf, cbuf_len)
	runtime.KeepAlive(cbuf_lenAllocMap)
	runtime.KeepAlive(cbufAllocMap)
	runtime.KeepAlive(ccmds_lenAllocMap)
	runtime.KeepAlive(ccmdsAllocMap)
	__v := (byte)(__ret)
	return __v
}

// bcm2835_smi_begin function as declared in all/bcm2835.h:2042
func bcm2835_smi_begin() int32 {
	__ret := C.bcm2835_smi_begin()
	__v := (int32)(__ret)
	return __v
}

// bcm2835_smi_end function as declared in all/bcm2835.h:2048
func bcm2835_smi_end() {
	C.bcm2835_smi_end()
}

// bcm2835_smi_set_timing function as declared in all/bcm2835.h:2063
func bcm2835_smi_set_timing(smichannel uint32, readchannel uint32, setupcycles uint32, strobecycles uint32, holdcycles uint32, pacecycles uint32) {
	csmichannel, csmichannelAllocMap := (C.uint32_t)(smichannel), cgoAllocsUnknown
	creadchannel, creadchannelAllocMap := (C.uint32_t)(readchannel), cgoAllocsUnknown
	csetupcycles, csetupcyclesAllocMap := (C.uint32_t)(setupcycles), cgoAllocsUnknown
	cstrobecycles, cstrobecyclesAllocMap := (C.uint32_t)(strobecycles), cgoAllocsUnknown
	choldcycles, choldcyclesAllocMap := (C.uint32_t)(holdcycles), cgoAllocsUnknown
	cpacecycles, cpacecyclesAllocMap := (C.uint32_t)(pacecycles), cgoAllocsUnknown
	C.bcm2835_smi_set_timing(csmichannel, creadchannel, csetupcycles, cstrobecycles, choldcycles, cpacecycles)
	runtime.KeepAlive(cpacecyclesAllocMap)
	runtime.KeepAlive(choldcyclesAllocMap)
	runtime.KeepAlive(cstrobecyclesAllocMap)
	runtime.KeepAlive(csetupcyclesAllocMap)
	runtime.KeepAlive(creadchannelAllocMap)
	runtime.KeepAlive(csmichannelAllocMap)
}

// bcm2835_smi_write function as declared in all/bcm2835.h:2074
func bcm2835_smi_write(smichannel uint32, data byte, address uint32) {
	csmichannel, csmichannelAllocMap := (C.uint32_t)(smichannel), cgoAllocsUnknown
	cdata, cdataAllocMap := (C.uint8_t)(data), cgoAllocsUnknown
	caddress, caddressAllocMap := (C.uint32_t)(address), cgoAllocsUnknown
	C.bcm2835_smi_write(csmichannel, cdata, caddress)
	runtime.KeepAlive(caddressAllocMap)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(csmichannelAllocMap)
}

// bcm2835_smi_read function as declared in all/bcm2835.h:2083
func bcm2835_smi_read(smichannel uint32, address uint32) uint32 {
	csmichannel, csmichannelAllocMap := (C.uint32_t)(smichannel), cgoAllocsUnknown
	caddress, caddressAllocMap := (C.uint32_t)(address), cgoAllocsUnknown
	__ret := C.bcm2835_smi_read(csmichannel, caddress)
	runtime.KeepAlive(caddressAllocMap)
	runtime.KeepAlive(csmichannelAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_st_read function as declared in all/bcm2835.h:2096
func bcm2835_st_read() uint32 {
	__ret := C.bcm2835_st_read()
	__v := (uint32)(__ret)
	return __v
}

// bcm2835_st_delay function as declared in all/bcm2835.h:2102
func bcm2835_st_delay(offset_micros uint32, micros uint32) {
	coffset_micros, coffset_microsAllocMap := (C.uint64_t)(offset_micros), cgoAllocsUnknown
	cmicros, cmicrosAllocMap := (C.uint64_t)(micros), cgoAllocsUnknown
	C.bcm2835_st_delay(coffset_micros, cmicros)
	runtime.KeepAlive(cmicrosAllocMap)
	runtime.KeepAlive(coffset_microsAllocMap)
}

// bcm2835_pwm_set_clock function as declared in all/bcm2835.h:2119
func bcm2835_pwm_set_clock(divisor uint32) {
	cdivisor, cdivisorAllocMap := (C.uint32_t)(divisor), cgoAllocsUnknown
	C.bcm2835_pwm_set_clock(cdivisor)
	runtime.KeepAlive(cdivisorAllocMap)
}

// bcm2835_pwm_set_mode function as declared in all/bcm2835.h:2127
func bcm2835_pwm_set_mode(channel byte, markspace byte, enabled byte) {
	cchannel, cchannelAllocMap := (C.uint8_t)(channel), cgoAllocsUnknown
	cmarkspace, cmarkspaceAllocMap := (C.uint8_t)(markspace), cgoAllocsUnknown
	cenabled, cenabledAllocMap := (C.uint8_t)(enabled), cgoAllocsUnknown
	C.bcm2835_pwm_set_mode(cchannel, cmarkspace, cenabled)
	runtime.KeepAlive(cenabledAllocMap)
	runtime.KeepAlive(cmarkspaceAllocMap)
	runtime.KeepAlive(cchannelAllocMap)
}

// bcm2835_pwm_set_range function as declared in all/bcm2835.h:2134
func bcm2835_pwm_set_range(channel byte, _range uint32) {
	cchannel, cchannelAllocMap := (C.uint8_t)(channel), cgoAllocsUnknown
	c_range, c_rangeAllocMap := (C.uint32_t)(_range), cgoAllocsUnknown
	C.bcm2835_pwm_set_range(cchannel, c_range)
	runtime.KeepAlive(c_rangeAllocMap)
	runtime.KeepAlive(cchannelAllocMap)
}

// bcm2835_pwm_set_data function as declared in all/bcm2835.h:2142
func bcm2835_pwm_set_data(channel byte, data uint32) {
	cchannel, cchannelAllocMap := (C.uint8_t)(channel), cgoAllocsUnknown
	cdata, cdataAllocMap := (C.uint32_t)(data), cgoAllocsUnknown
	C.bcm2835_pwm_set_data(cchannel, cdata)
	runtime.KeepAlive(cdataAllocMap)
	runtime.KeepAlive(cchannelAllocMap)
}

// EPD_IT8951_SystemRun function as declared in all/EPD_IT8951.h:187
func EPD_IT8951_SystemRun() {
	C.EPD_IT8951_SystemRun()
}

// EPD_IT8951_Standby function as declared in all/EPD_IT8951.h:189
func EPD_IT8951_Standby() {
	C.EPD_IT8951_Standby()
}

// EPD_IT8951_Sleep function as declared in all/EPD_IT8951.h:191
func EPD_IT8951_Sleep() {
	C.EPD_IT8951_Sleep()
}

// EPD_IT8951_Init function as declared in all/EPD_IT8951.h:193
func EPD_IT8951_Init(vCOM uint16) IT8951_Dev_Info {
	cvCOM, cvCOMAllocMap := (C.uint16_t)(vCOM), cgoAllocsUnknown
	__ret := C.EPD_IT8951_Init(cvCOM)
	runtime.KeepAlive(cvCOMAllocMap)
	__v := *NewIT8951_Dev_InfoRef(unsafe.Pointer(&__ret))
	return __v
}

// EPD_IT8951_Clear_Refresh function as declared in all/EPD_IT8951.h:195
func EPD_IT8951_Clear_Refresh(dev_Info IT8951_Dev_Info, target_Memory_Addr uint32, mode uint16) {
	cdev_Info, cdev_InfoAllocMap := dev_Info.PassValue()
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	cmode, cmodeAllocMap := (C.uint16_t)(mode), cgoAllocsUnknown
	C.EPD_IT8951_Clear_Refresh(cdev_Info, ctarget_Memory_Addr, cmode)
	runtime.KeepAlive(cmodeAllocMap)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(cdev_InfoAllocMap)
}

// EPD_IT8951_1bp_Refresh function as declared in all/EPD_IT8951.h:197
func EPD_IT8951_1bp_Refresh(frame_Buf []byte, x uint16, y uint16, w uint16, h uint16, mode byte, target_Memory_Addr uint32, packed_Write bool) {
	cframe_Buf, cframe_BufAllocMap := copyPUint8_tBytes((*sliceHeader)(unsafe.Pointer(&frame_Buf)))
	cx, cxAllocMap := (C.uint16_t)(x), cgoAllocsUnknown
	cy, cyAllocMap := (C.uint16_t)(y), cgoAllocsUnknown
	cw, cwAllocMap := (C.uint16_t)(w), cgoAllocsUnknown
	ch, chAllocMap := (C.uint16_t)(h), cgoAllocsUnknown
	cmode, cmodeAllocMap := (C.uint8_t)(mode), cgoAllocsUnknown
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	cpacked_Write, cpacked_WriteAllocMap := (C._Bool)(packed_Write), cgoAllocsUnknown
	C.EPD_IT8951_1bp_Refresh(cframe_Buf, cx, cy, cw, ch, cmode, ctarget_Memory_Addr, cpacked_Write)
	runtime.KeepAlive(cpacked_WriteAllocMap)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(cmodeAllocMap)
	runtime.KeepAlive(chAllocMap)
	runtime.KeepAlive(cwAllocMap)
	runtime.KeepAlive(cyAllocMap)
	runtime.KeepAlive(cxAllocMap)
	runtime.KeepAlive(cframe_BufAllocMap)
}

// EPD_IT8951_1bp_Multi_Frame_Write function as declared in all/EPD_IT8951.h:198
func EPD_IT8951_1bp_Multi_Frame_Write(frame_Buf []byte, x uint16, y uint16, w uint16, h uint16, target_Memory_Addr uint32, packed_Write bool) {
	cframe_Buf, cframe_BufAllocMap := copyPUint8_tBytes((*sliceHeader)(unsafe.Pointer(&frame_Buf)))
	cx, cxAllocMap := (C.uint16_t)(x), cgoAllocsUnknown
	cy, cyAllocMap := (C.uint16_t)(y), cgoAllocsUnknown
	cw, cwAllocMap := (C.uint16_t)(w), cgoAllocsUnknown
	ch, chAllocMap := (C.uint16_t)(h), cgoAllocsUnknown
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	cpacked_Write, cpacked_WriteAllocMap := (C._Bool)(packed_Write), cgoAllocsUnknown
	C.EPD_IT8951_1bp_Multi_Frame_Write(cframe_Buf, cx, cy, cw, ch, ctarget_Memory_Addr, cpacked_Write)
	runtime.KeepAlive(cpacked_WriteAllocMap)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(chAllocMap)
	runtime.KeepAlive(cwAllocMap)
	runtime.KeepAlive(cyAllocMap)
	runtime.KeepAlive(cxAllocMap)
	runtime.KeepAlive(cframe_BufAllocMap)
}

// EPD_IT8951_1bp_Multi_Frame_Refresh function as declared in all/EPD_IT8951.h:199
func EPD_IT8951_1bp_Multi_Frame_Refresh(x uint16, y uint16, w uint16, h uint16, target_Memory_Addr uint32) {
	cx, cxAllocMap := (C.uint16_t)(x), cgoAllocsUnknown
	cy, cyAllocMap := (C.uint16_t)(y), cgoAllocsUnknown
	cw, cwAllocMap := (C.uint16_t)(w), cgoAllocsUnknown
	ch, chAllocMap := (C.uint16_t)(h), cgoAllocsUnknown
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	C.EPD_IT8951_1bp_Multi_Frame_Refresh(cx, cy, cw, ch, ctarget_Memory_Addr)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(chAllocMap)
	runtime.KeepAlive(cwAllocMap)
	runtime.KeepAlive(cyAllocMap)
	runtime.KeepAlive(cxAllocMap)
}

// EPD_IT8951_2bp_Refresh function as declared in all/EPD_IT8951.h:201
func EPD_IT8951_2bp_Refresh(frame_Buf []byte, x uint16, y uint16, w uint16, h uint16, hold bool, target_Memory_Addr uint32, packed_Write bool) {
	cframe_Buf, cframe_BufAllocMap := copyPUint8_tBytes((*sliceHeader)(unsafe.Pointer(&frame_Buf)))
	cx, cxAllocMap := (C.uint16_t)(x), cgoAllocsUnknown
	cy, cyAllocMap := (C.uint16_t)(y), cgoAllocsUnknown
	cw, cwAllocMap := (C.uint16_t)(w), cgoAllocsUnknown
	ch, chAllocMap := (C.uint16_t)(h), cgoAllocsUnknown
	chold, choldAllocMap := (C._Bool)(hold), cgoAllocsUnknown
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	cpacked_Write, cpacked_WriteAllocMap := (C._Bool)(packed_Write), cgoAllocsUnknown
	C.EPD_IT8951_2bp_Refresh(cframe_Buf, cx, cy, cw, ch, chold, ctarget_Memory_Addr, cpacked_Write)
	runtime.KeepAlive(cpacked_WriteAllocMap)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(choldAllocMap)
	runtime.KeepAlive(chAllocMap)
	runtime.KeepAlive(cwAllocMap)
	runtime.KeepAlive(cyAllocMap)
	runtime.KeepAlive(cxAllocMap)
	runtime.KeepAlive(cframe_BufAllocMap)
}

// EPD_IT8951_4bp_Refresh function as declared in all/EPD_IT8951.h:203
func EPD_IT8951_4bp_Refresh(frame_Buf []byte, x uint16, y uint16, w uint16, h uint16, hold bool, target_Memory_Addr uint32, packed_Write bool) {
	cframe_Buf, cframe_BufAllocMap := copyPUint8_tBytes((*sliceHeader)(unsafe.Pointer(&frame_Buf)))
	cx, cxAllocMap := (C.uint16_t)(x), cgoAllocsUnknown
	cy, cyAllocMap := (C.uint16_t)(y), cgoAllocsUnknown
	cw, cwAllocMap := (C.uint16_t)(w), cgoAllocsUnknown
	ch, chAllocMap := (C.uint16_t)(h), cgoAllocsUnknown
	chold, choldAllocMap := (C._Bool)(hold), cgoAllocsUnknown
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	cpacked_Write, cpacked_WriteAllocMap := (C._Bool)(packed_Write), cgoAllocsUnknown
	C.EPD_IT8951_4bp_Refresh(cframe_Buf, cx, cy, cw, ch, chold, ctarget_Memory_Addr, cpacked_Write)
	runtime.KeepAlive(cpacked_WriteAllocMap)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(choldAllocMap)
	runtime.KeepAlive(chAllocMap)
	runtime.KeepAlive(cwAllocMap)
	runtime.KeepAlive(cyAllocMap)
	runtime.KeepAlive(cxAllocMap)
	runtime.KeepAlive(cframe_BufAllocMap)
}

// EPD_IT8951_8bp_Refresh function as declared in all/EPD_IT8951.h:205
func EPD_IT8951_8bp_Refresh(frame_Buf []byte, x uint16, y uint16, w uint16, h uint16, hold bool, target_Memory_Addr uint32) {
	cframe_Buf, cframe_BufAllocMap := copyPUint8_tBytes((*sliceHeader)(unsafe.Pointer(&frame_Buf)))
	cx, cxAllocMap := (C.uint16_t)(x), cgoAllocsUnknown
	cy, cyAllocMap := (C.uint16_t)(y), cgoAllocsUnknown
	cw, cwAllocMap := (C.uint16_t)(w), cgoAllocsUnknown
	ch, chAllocMap := (C.uint16_t)(h), cgoAllocsUnknown
	chold, choldAllocMap := (C._Bool)(hold), cgoAllocsUnknown
	ctarget_Memory_Addr, ctarget_Memory_AddrAllocMap := (C.uint32_t)(target_Memory_Addr), cgoAllocsUnknown
	C.EPD_IT8951_8bp_Refresh(cframe_Buf, cx, cy, cw, ch, chold, ctarget_Memory_Addr)
	runtime.KeepAlive(ctarget_Memory_AddrAllocMap)
	runtime.KeepAlive(choldAllocMap)
	runtime.KeepAlive(chAllocMap)
	runtime.KeepAlive(cwAllocMap)
	runtime.KeepAlive(cyAllocMap)
	runtime.KeepAlive(cxAllocMap)
	runtime.KeepAlive(cframe_BufAllocMap)
}