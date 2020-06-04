package main

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
