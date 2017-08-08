package zhengsen

import "fmt"

type Person interface {
	getName() string
	Connector
}

type Connector interface {
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
	a := Connector(t)
	a.getPhone()
}
