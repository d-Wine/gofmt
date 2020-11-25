package main

import "fmt"

func main(){
	//declare variables
	var i int
	var s string
	//initialize variables
	i = 20
	s = "Some String"
	fmt.Println(i)
	fmt.Println(s)
	//declare and initialize variables
	var k int = 34
	fmt.Println(k)
	//declare short variable
	l := 30
	j := "Some String"
	fmt.Println(l)
	fmt.Println(j)
	//declare multiple variables
	firstName,lastName := "Ar","Kar"
	fmt.Println(firstName+lastName)
	//declare variable blocks
	var(
		name = "Dy"
		age = 23
	)
	fmt.Println(name)
	fmt.Println(age)
}
