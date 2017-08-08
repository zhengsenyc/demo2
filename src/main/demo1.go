package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	x := "hello!"
	for i:=0; i<len(x);i++{
		x := x[i]
		if x!='!'{
			x := x + 'A' - 'a'
			fmt.Printf("%c",x)
		}
	}

	fmt.Println("===============================")
	run1()
}

var cwd string

func run1(){
	cwd,err := os.Getwd()
	if err != nil {
		fmt.Println("------------------")
		log.Fatalf("报错了  %v",err)
	}

	log.Printf("hehaha   %s",cwd)

	var a int8 = 1
	fmt.Println(a)
	a = a << 3
	fmt.Println(a)
	a = a >> 3
	fmt.Println(a)
	a = -128
	fmt.Println(a)
	a = a>>2
	fmt.Println(a)


}
