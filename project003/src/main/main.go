// 方法与属性的获取

package main

import (
	"fmt"
	"reflect"
)

func main() {
	neo := per{
		Age:  10,
		Name: "ca",
		Mark: 702,
		Sex:  true,
	}
	test(neo)

}

// 别忘了把方法属性公有化

type per struct {
	Age  int
	Name string
	Mark int `json:"score"`
	Sex  bool
}

func (p per) Say()  {
	fmt.Printf("i'm %s,%d years old",p.Name,p.Age)
}

func (p per) Iq(w int)int  {
	res := w * w + p.Age - w
	return res
}

func test(x interface{})  {

	v := reflect.ValueOf(x)
	k := v.Kind()
	t := reflect.TypeOf(x)
	path := t.PkgPath()
	fmt.Println("from: ",path,"\n")

	fmt.Println("-------------------")

	if k!=reflect.Struct{
		return
	}

	numfield := v.NumField()
	// 字段数
	for i:=0;i<numfield;i++{

		field := t.Field(i)
		// 获取属性
		fieldtype := field.Type
		// 获取属性的要素

		fieldname := field.Name
		tag := field.Tag.Get("json")
		if tag!=""{
			fieldname = tag
		}
		// 获取tag

		value := v.Field(i)
		// 注意，是value类型
		// 获取字段值
		fmt.Println("name: ",fieldname)
		fmt.Println("type: ",fieldtype)
		fmt.Println("value: ",value)
		fmt.Println()

	}

	fmt.Println("-------------------")

	nummethod := v.NumMethod()
	fmt.Println(nummethod)

	met1 := v.Method(0)
	// 获取method
	// method按照名字排序

	params1 := []reflect.Value{reflect.ValueOf(6)}
	// 参数传入用value切片

	result1 := met1.Call(params1)
	fmt.Println(result1[0])
	// 结果传出用value切片


	met2 := v.Method(1)
	met2.Call(nil)




}




