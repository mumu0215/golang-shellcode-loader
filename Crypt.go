package main

import (
	"flag"
	"fmt"
	"os"
)
var filename=flag.String("f","","fileName")
var data=flag.String("s","","input string")
var	encodeing=flag.Bool("e",false,"encode")
var decodeing=flag.Bool("d",false,"decode")

var key=[]byte{97,97,114,111,110}
func encode(s []byte) string {
	i:=0
	for k,v:=range s{
		if i!=4{
			s[k]=v+key[i]
			i++
		}else {
			i=0
		}
	}
	return string(s)
}
func decode(s []byte) string {
	i:=0
	for k,v:=range s{
		if i!=4{
			s[k]=v-key[i]
			i++
		}else {
			i=0
		}
	}
	return string(s)
}
func main() {
	flag.Parse()
	fileName:=data
	f,err:=os.OpenFile(*fileName,os.O_RDONLY,0666)
	if err!=nil{
		fmt.Println("fail to open file")
		return
	}
	buffer:=make([]byte,4096)
	_,_=f.Read(buffer)
	fmt.Println(decode())
}
