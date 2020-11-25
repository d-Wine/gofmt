package main

import "fmt"

//declare constants
const PI float64 = 3.14159265359
const VALUE float64 = 1000
//declare multiple constants block
const(
	NAME = "NAME"
	AGE = 45
)

func main()  {
	fmt.Println(PI)
	fmt.Println(VALUE)
	fmt.Println(NAME)
	fmt.Println(AGE)
}