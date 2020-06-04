package main

import (
	"flag"
	//"fmt"
	"os"
)
var filename=flag.String("f","","fileName")
var data=flag.String("s","","input string")
var	encodeing=flag.Bool("e",false,"encode")
var decodeing=flag.Bool("d",false,"decode")


func main() {
	flag.Parse()
	if *filename==""{
		return
	}
	fileName:=filename
	//buffer:=make([]byte,4096)
	//k,_:=f.Read(buffer)
	//fmt.Println(buffer[k-1])
//	t:=encode(buffer[:k])
//	fmt.Println(decode([]byte(t)))
	if *encodeing&&*data!=""{
		f,err:=os.OpenFile(*fileName,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
		if err!=nil{
			return
		}
		temp:=*data
		encodeTemp:=encode([]byte(temp))
		f.WriteString(encodeTemp)
		return
	}else if *decodeing{
		f,err:=os.OpenFile(*fileName,os.O_RDWR,0666)
		if err!=nil{
			return
		}
		buffer:=make([]byte,4096)
		size,_:=f.Read(buffer)
		temp:=buffer[:size]
		//os.Truncate(*filename,0)
		f.Truncate(0)
		f.Seek(0,0)
		f.WriteString(decode(temp))
		return
	}else{
		return}
}
