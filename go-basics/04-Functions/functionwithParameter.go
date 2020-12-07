package main

import "fmt"

func main()  {
	hello("Gofmt")
	addTwoNum(3,4)
}
func hello(x string){	// string parameter
	fmt.Println(x)
}
func addTwoNum(x,y int)  {	// int parameter
	fmt.Println(x+y)
}