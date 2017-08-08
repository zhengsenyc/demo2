package main

import "fmt"
/*这是一个空接口，这就相当于Java中的Object类*/
type Empty interface {

}

type Connector interface {
	PhoneCall() string
	PhoneNumber
}

type PhoneNumber interface {
	GetPhoneNumber()
}

type Person struct {
	name string
}

func (p Person) PhoneCall() string{
	return p.name
}

func (p Person) GetPhoneNumber(){
	fmt.Println("phoneNumber=",110)
}

func main(){
	a := Person{"hehe"}
	a.GetPhoneNumber()
	fmt.Println(a.PhoneCall())
	run6(a)
}

/*这个方法是用来验证a是否实现了接口*/
func run6(con interface{}){
	/*因为使用了空接口实现的结构体是太多了，那么直接这么判断效率太低*/
	/*if pc,ok := con.(Person);ok { //判断person实现了这个接口么？
		fmt.Println("实现了",pc.name)
		return
	}*/
	/*采用switch语句*/
	switch a:=con.(type) {
	case Person:
		fmt.Println("实现了",a.name)
	default:
		fmt.Println("很抱歉，没有实现！")
	}
}
