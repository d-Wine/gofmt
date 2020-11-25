package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {
	//convert string to float
	s := "3.1415926535"
	f,_ := strconv.ParseFloat(s,8)
	fmt.Println(reflect.TypeOf(f))
	//convert string to bool
	str := "0"	// 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	b,_ := strconv.ParseBool(str)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)
	str = "1"
	b,_ = strconv.ParseBool(str)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)
	//convert float to int
	var f32 float32 = 3.1415926535
	fmt.Println(reflect.TypeOf(f32))
	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
}