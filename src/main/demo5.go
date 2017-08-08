package main

import "fmt"

type A int

func main(){
	var a A
	a.run5(100)
	fmt.Println(a)

}

func (a *A) run5(num int){
	*a += A(num)
}
