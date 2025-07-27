package newpg

import "fmt"

func Summation(a float64, b int) {
	fmt.Print("summation: ", a+float64(b), "\n")


	newvar:= a*float64(b)
	fmt.Print("newvar: ", newvar, "\n")
}