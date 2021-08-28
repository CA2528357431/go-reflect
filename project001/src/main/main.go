// value type kind

package main

import (
	"fmt"
	"reflect"
)

func main() {

	neoint := 10
	inttest(neoint)

	fmt.Println()

	neostruct := per{
		age:  10,
		name: "ca",
	}
	structtest(neostruct)

	
}

func inttest(x interface{})  {
	v := reflect.ValueOf(x)
	fmt.Println("value:",v)

	t0 := reflect.TypeOf(x)
	t1 := v.Type()
	fmt.Println("type:",t0,t1)

	k0 := v.Kind()
	k1 := t0.Kind()
	fmt.Println("kind:",k0,k1)

	i := v.Interface()
	fmt.Println("interface:",i)

	n := v.Int()
	fmt.Println("num:",n)


}

// kind是struct
// type是per

func structtest(x interface{})  {
	v := reflect.ValueOf(x)
	fmt.Println("value:",v)

	t0 := reflect.TypeOf(x)
	t1 := v.Type()
	fmt.Println("type:",t0,t1)

	k0 := v.Kind()
	k1 := t0.Kind()
	fmt.Println("kind:",k0,k1)

	i := v.Interface()
	fmt.Println("interface:",i)

	obj := i.(per)
	fmt.Println("object:",obj)



}

type per struct {
	age int
	name string
}
