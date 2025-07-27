package main

import (
	"fmt"
	errorHandle "mymodule/errorhandle"
)

func main() {
	// fmt.Println("Hello World")
	// var a, b int
	// fmt.Print("Enter two numbers: ")
	// fmt.Scan(&a, &b)
	// ans := sum.Summation(a, b)
	// fmt.Print("Sum of ", a, " and ", b, " is ", ans)

	// loop.Looping(10)
	ans, _ := errorHandle.ErrorHandle(10, 9)
	fmt.Println(ans)

}
