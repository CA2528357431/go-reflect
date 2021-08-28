// 通过反射改原值

package main

import (
	"fmt"
	"reflect"
)

func main() {

	neoint := 10
	inttest(&neoint)
	fmt.Println(neoint)

	fmt.Println()

	neostruct := per{
		age:  10,
		name: "ca",
	}
	structtest(&neostruct)
	fmt.Println(neostruct)

	
}

func inttest(x interface{})  {
	p := reflect.ValueOf(x)
	v := p.Elem()
	v.SetInt(99999)

}



func structtest(x interface{})  {
	p := reflect.ValueOf(x)
	v := p.Elem()

	neo := per{
		age:  999,
		name: "shit",
	}
	neov := reflect.ValueOf(neo)

	v.Set(neov)
}

type per struct {
	age int
	name string
}
