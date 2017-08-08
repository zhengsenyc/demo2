package main

import (
	"fmt"
	"unsafe"
)

func main(){
	s := "hello world"
	fmt.Println(len(s))

	//a := 'a'
	fmt.Println(unsafe.Sizeof(s))
}
