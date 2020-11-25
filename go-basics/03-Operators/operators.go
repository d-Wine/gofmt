package main

import "fmt"

func main()  {
	var(
		i int = 10
		j int = 20
		k int = 30
	)
	//arithmetic operators
	fmt.Println("+ operator 	=> i + j",i+j)
	fmt.Println("- operator 	=> i - j",i-j)
	fmt.Println("* operator 	=> i * j",i*j)
	fmt.Println("/ operator 	=> i / j",i/j)
	fmt.Println("mod operator 	=> i mod j",i%j)
	//comparison operators
	fmt.Println("== operator	=> i == k",i == k)
	fmt.Println("!= operator	=> i == k",i != k)
	fmt.Println("> operator	=> i == k",i > k)
	fmt.Println(">= operator	=> i == k",i >= k)
	fmt.Println("< operator	=> i == k",i < k)
	fmt.Println("<= operator	=> i == k",i <= k)
	//logical operators
	fmt.Println("&& operator	=> true && true	 ",true && true)
	fmt.Println("&& operator	=> true && false ",true && false)
	fmt.Println("&& operator	=> false && true ",false && true)
	fmt.Println("&& operator	=> false && false",false && false)
	fmt.Println("|| operator	=> true || true	 ",true || true)
	fmt.Println("|| operator	=> true || false ",true || false)
	fmt.Println("|| operator	=> false || true ",false || true)
	fmt.Println("|| operator	=> false || false",false || false)
	//assignment operator
	x := 100
	y := 25
	x +=y
	fmt.Println(x)
	x = 100
	y = 25
	x -= y
	fmt.Println(x)
	x = 100
	y = 25
	x = y
	fmt.Println(x)
	x = 100
	y = 25
	x *= y
	fmt.Println(x)
	x = 100
	y = 25
	x /= y
	fmt.Println(x)
	x = 101
	y = 25
	x %= y
	fmt.Println()

}