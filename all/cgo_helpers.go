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
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocIT8951_Load_Img_InfoMemory allocates memory for type C.IT8951_Load_Img_Info in C.
// The caller is responsible for freeing the this memory via C.free.
func allocIT8951_Load_Img_InfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfIT8951_Load_Img_InfoValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfIT8951_Load_Img_InfoValue = unsafe.Sizeof([1]C.IT8951_Load_Img_Info{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *IT8951_Load_Img_Info) Ref() *C.IT8951_Load_Img_Info {
	if x == nil {
		return nil
	}
	return x.ref66739e64
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *IT8951_Load_Img_Info) Free() {
	if x != nil && x.allocs66739e64 != nil {
		x.allocs66739e64.(*cgoAllocMap).Free()
		x.ref66739e64 = nil
	}
}

// NewIT8951_Load_Img_InfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewIT8951_Load_Img_InfoRef(ref unsafe.Pointer) *IT8951_Load_Img_Info {
	if ref == nil {
		return nil
	}
	obj := new(IT8951_Load_Img_Info)
	obj.ref66739e64 = (*C.IT8951_Load_Img_Info)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *IT8951_Load_Img_Info) PassRef() (*C.IT8951_Load_Img_Info, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref66739e64 != nil {
		return x.ref66739e64, nil
	}
	mem66739e64 := allocIT8951_Load_Img_InfoMemory(1)
	ref66739e64 := (*C.IT8951_Load_Img_Info)(mem66739e64)
	allocs66739e64 := new(cgoAllocMap)
	allocs66739e64.Add(mem66739e64)

	var cEndian_Type_allocs *cgoAllocMap
	ref66739e64.Endian_Type, cEndian_Type_allocs = (C.uint16_t)(x.Endian_Type), cgoAllocsUnknown
	allocs66739e64.Borrow(cEndian_Type_allocs)

	var cPixel_Format_allocs *cgoAllocMap
	ref66739e64.Pixel_Format, cPixel_Format_allocs = (C.uint16_t)(x.Pixel_Format), cgoAllocsUnknown
	allocs66739e64.Borrow(cPixel_Format_allocs)

	var cRotate_allocs *cgoAllocMap
	ref66739e64.Rotate, cRotate_allocs = (C.uint16_t)(x.Rotate), cgoAllocsUnknown
	allocs66739e64.Borrow(cRotate_allocs)

	var cSource_Buffer_Addr_allocs *cgoAllocMap
	ref66739e64.Source_Buffer_Addr, cSource_Buffer_Addr_allocs = (C.uint32_t)(x.Source_Buffer_Addr), cgoAllocsUnknown
	allocs66739e64.Borrow(cSource_Buffer_Addr_allocs)

	var cTarget_Memory_Addr_allocs *cgoAllocMap
	ref66739e64.Target_Memory_Addr, cTarget_Memory_Addr_allocs = (C.uint32_t)(x.Target_Memory_Addr), cgoAllocsUnknown
	allocs66739e64.Borrow(cTarget_Memory_Addr_allocs)

	x.ref66739e64 = ref66739e64
	x.allocs66739e64 = allocs66739e64
	return ref66739e64, allocs66739e64

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x IT8951_Load_Img_Info) PassValue() (C.IT8951_Load_Img_Info, *cgoAllocMap) {
	if x.ref66739e64 != nil {
		return *x.ref66739e64, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *IT8951_Load_Img_Info) Deref() {
	if x.ref66739e64 == nil {
		return
	}
	x.Endian_Type = (uint16)(x.ref66739e64.Endian_Type)
	x.Pixel_Format = (uint16)(x.ref66739e64.Pixel_Format)
	x.Rotate = (uint16)(x.ref66739e64.Rotate)
	x.Source_Buffer_Addr = (uint32)(x.ref66739e64.Source_Buffer_Addr)
	x.Target_Memory_Addr = (uint32)(x.ref66739e64.Target_Memory_Addr)
}

// allocIT8951_Area_Img_InfoMemory allocates memory for type C.IT8951_Area_Img_Info in C.
// The caller is responsible for freeing the this memory via C.free.
func allocIT8951_Area_Img_InfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfIT8951_Area_Img_InfoValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfIT8951_Area_Img_InfoValue = unsafe.Sizeof([1]C.IT8951_Area_Img_Info{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *IT8951_Area_Img_Info) Ref() *C.IT8951_Area_Img_Info {
	if x == nil {
		return nil
	}
	return x.refe518aace
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *IT8951_Area_Img_Info) Free() {
	if x != nil && x.allocse518aace != nil {
		x.allocse518aace.(*cgoAllocMap).Free()
		x.refe518aace = nil
	}
}

// NewIT8951_Area_Img_InfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewIT8951_Area_Img_InfoRef(ref unsafe.Pointer) *IT8951_Area_Img_Info {
	if ref == nil {
		return nil
	}
	obj := new(IT8951_Area_Img_Info)
	obj.refe518aace = (*C.IT8951_Area_Img_Info)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *IT8951_Area_Img_Info) PassRef() (*C.IT8951_Area_Img_Info, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe518aace != nil {
		return x.refe518aace, nil
	}
	meme518aace := allocIT8951_Area_Img_InfoMemory(1)
	refe518aace := (*C.IT8951_Area_Img_Info)(meme518aace)
	allocse518aace := new(cgoAllocMap)
	allocse518aace.Add(meme518aace)

	var cArea_X_allocs *cgoAllocMap
	refe518aace.Area_X, cArea_X_allocs = (C.uint16_t)(x.Area_X), cgoAllocsUnknown
	allocse518aace.Borrow(cArea_X_allocs)

	var cArea_Y_allocs *cgoAllocMap
	refe518aace.Area_Y, cArea_Y_allocs = (C.uint16_t)(x.Area_Y), cgoAllocsUnknown
	allocse518aace.Borrow(cArea_Y_allocs)

	var cArea_W_allocs *cgoAllocMap
	refe518aace.Area_W, cArea_W_allocs = (C.uint16_t)(x.Area_W), cgoAllocsUnknown
	allocse518aace.Borrow(cArea_W_allocs)

	var cArea_H_allocs *cgoAllocMap
	refe518aace.Area_H, cArea_H_allocs = (C.uint16_t)(x.Area_H), cgoAllocsUnknown
	allocse518aace.Borrow(cArea_H_allocs)

	x.refe518aace = refe518aace
	x.allocse518aace = allocse518aace
	return refe518aace, allocse518aace

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x IT8951_Area_Img_Info) PassValue() (C.IT8951_Area_Img_Info, *cgoAllocMap) {
	if x.refe518aace != nil {
		return *x.refe518aace, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *IT8951_Area_Img_Info) Deref() {
	if x.refe518aace == nil {
		return
	}
	x.Area_X = (uint16)(x.refe518aace.Area_X)
	x.Area_Y = (uint16)(x.refe518aace.Area_Y)
	x.Area_W = (uint16)(x.refe518aace.Area_W)
	x.Area_H = (uint16)(x.refe518aace.Area_H)
}

// allocIT8951_Dev_InfoMemory allocates memory for type C.IT8951_Dev_Info in C.
// The caller is responsible for freeing the this memory via C.free.
func allocIT8951_Dev_InfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfIT8951_Dev_InfoValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfIT8951_Dev_InfoValue = unsafe.Sizeof([1]C.IT8951_Dev_Info{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *IT8951_Dev_Info) Ref() *C.IT8951_Dev_Info {
	if x == nil {
		return nil
	}
	return x.refab7ec1bd
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *IT8951_Dev_Info) Free() {
	if x != nil && x.allocsab7ec1bd != nil {
		x.allocsab7ec1bd.(*cgoAllocMap).Free()
		x.refab7ec1bd = nil
	}
}

// NewIT8951_Dev_InfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewIT8951_Dev_InfoRef(ref unsafe.Pointer) *IT8951_Dev_Info {
	if ref == nil {
		return nil
	}
	obj := new(IT8951_Dev_Info)
	obj.refab7ec1bd = (*C.IT8951_Dev_Info)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *IT8951_Dev_Info) PassRef() (*C.IT8951_Dev_Info, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refab7ec1bd != nil {
		return x.refab7ec1bd, nil
	}
	memab7ec1bd := allocIT8951_Dev_InfoMemory(1)
	refab7ec1bd := (*C.IT8951_Dev_Info)(memab7ec1bd)
	allocsab7ec1bd := new(cgoAllocMap)
	allocsab7ec1bd.Add(memab7ec1bd)

	var cPanel_W_allocs *cgoAllocMap
	refab7ec1bd.Panel_W, cPanel_W_allocs = (C.uint16_t)(x.Panel_W), cgoAllocsUnknown
	allocsab7ec1bd.Borrow(cPanel_W_allocs)

	var cPanel_H_allocs *cgoAllocMap
	refab7ec1bd.Panel_H, cPanel_H_allocs = (C.uint16_t)(x.Panel_H), cgoAllocsUnknown
	allocsab7ec1bd.Borrow(cPanel_H_allocs)

	var cMemory_Addr_L_allocs *cgoAllocMap
	refab7ec1bd.Memory_Addr_L, cMemory_Addr_L_allocs = (C.uint16_t)(x.Memory_Addr_L), cgoAllocsUnknown
	allocsab7ec1bd.Borrow(cMemory_Addr_L_allocs)

	var cMemory_Addr_H_allocs *cgoAllocMap
	refab7ec1bd.Memory_Addr_H, cMemory_Addr_H_allocs = (C.uint16_t)(x.Memory_Addr_H), cgoAllocsUnknown
	allocsab7ec1bd.Borrow(cMemory_Addr_H_allocs)

	var cFW_Version_allocs *cgoAllocMap
	refab7ec1bd.FW_Version, cFW_Version_allocs = *(*[8]C.uint16_t)(unsafe.Pointer(&x.FW_Version)), cgoAllocsUnknown
	allocsab7ec1bd.Borrow(cFW_Version_allocs)

	var cLUT_Version_allocs *cgoAllocMap
	refab7ec1bd.LUT_Version, cLUT_Version_allocs = *(*[8]C.uint16_t)(unsafe.Pointer(&x.LUT_Version)), cgoAllocsUnknown
	allocsab7ec1bd.Borrow(cLUT_Version_allocs)

	x.refab7ec1bd = refab7ec1bd
	x.allocsab7ec1bd = allocsab7ec1bd
	return refab7ec1bd, allocsab7ec1bd

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x IT8951_Dev_Info) PassValue() (C.IT8951_Dev_Info, *cgoAllocMap) {
	if x.refab7ec1bd != nil {
		return *x.refab7ec1bd, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *IT8951_Dev_Info) Deref() {
	if x.refab7ec1bd == nil {
		return
	}
	x.Panel_W = (uint16)(x.refab7ec1bd.Panel_W)
	x.Panel_H = (uint16)(x.refab7ec1bd.Panel_H)
	x.Memory_Addr_L = (uint16)(x.refab7ec1bd.Memory_Addr_L)
	x.Memory_Addr_H = (uint16)(x.refab7ec1bd.Memory_Addr_H)
	x.FW_Version = *(*[8]uint16)(unsafe.Pointer(&x.refab7ec1bd.FW_Version))
	x.LUT_Version = *(*[8]uint16)(unsafe.Pointer(&x.refab7ec1bd.LUT_Version))
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// copyPUint32_tBytes copies the data from Go slice as *C.uint32_t.
func copyPUint32_tBytes(slice *sliceHeader) (*C.uint32_t, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfUint32_tValue) * slice.Len,
		Cap:  int(sizeOfUint32_tValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.uint32_t)(mem0), allocs
}

// allocUint32_tMemory allocates memory for type C.uint32_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUint32_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUint32_tValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfUint32_tValue = unsafe.Sizeof([1]C.uint32_t{})

// copyPCharBytes copies the data from Go slice as *C.char.
func copyPCharBytes(slice *sliceHeader) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfCharValue) * slice.Len,
		Cap:  int(sizeOfCharValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.char)(mem0), allocs
}

// allocCharMemory allocates memory for type C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfCharValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfCharValue = unsafe.Sizeof([1]C.char{})

// unpackPCharString copies the data from Go string as *C.char.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CString(str))
	allocs.Add(mem0)
	return (*C.char)(mem0), allocs
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// copyPUint8_tBytes copies the data from Go slice as *C.uint8_t.
func copyPUint8_tBytes(slice *sliceHeader) (*C.uint8_t, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	mem0 := unsafe.Pointer(C.CBytes(*(*[]byte)(unsafe.Pointer(&sliceHeader{
		Data: slice.Data,
		Len:  int(sizeOfUint8_tValue) * slice.Len,
		Cap:  int(sizeOfUint8_tValue) * slice.Len,
	}))))
	allocs.Add(mem0)

	return (*C.uint8_t)(mem0), allocs
}

// allocUint8_tMemory allocates memory for type C.uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUint8_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUint8_tValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfUint8_tValue = unsafe.Sizeof([1]C.uint8_t{})