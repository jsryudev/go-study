package main

import (
	"os"
	"strconv"
)

func main() {
	sum, _ := calc(os.Args[1], os.Args[2])
	println("Sum:", sum)
	//println("Mul:", mul)

	firstName := "John"
	updateName(&firstName)
	println(firstName)
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(name *string) {
	*name = "David"
}