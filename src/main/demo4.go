package main

import (
	"fmt"
	"strconv"
)

func main(){
	a := "hello world"
	for i,v := range a {
		fmt.Println(i,v)
		fmt.Printf("%q\n",v)
	}

	fmt.Println("====================================")
	b := strconv.Itoa(123)
	fmt.Println(b)
}
