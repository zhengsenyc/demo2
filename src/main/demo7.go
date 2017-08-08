package main

import "fmt"

type Person1 interface {
	getName() string
	Connector1
}

type Connector1 interface {
	getPhone()
}

type Teacher struct {
	name string
	age int
}

func (t Teacher) getName() string{
	return t.name
}

func (t Teacher) getPhone(){
	fmt.Println("年龄：",t.age)
}

func main(){
	t := Teacher{
		age : 12,
	}
	a := Connector1(t)
	a.getPhone()
}
