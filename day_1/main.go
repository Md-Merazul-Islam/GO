package main

import (
	"fmt"
	"mypractice/newpg"
	"mypractice/utils"
)

func main() {
	fmt.Println("Bismillah hir Rahmanir Rahim")
	utils.PrintSomething("hey there!")
	utils.PrintAnotherThing(("how are you?"))

	newpg.Summation(10.5, 5)
}
