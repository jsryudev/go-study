package main

import (
	"math"
	"strconv"
)


func main() {
  //var integer32 int = 2147483648
  //println(integer32)

	//var integer uint = -10
	//println(integer)

	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	println(float32, float64)

	println(math.MaxFloat32, math.MaxFloat64)

	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	//var featureFlag bool = true

	var firstName string = "John"
	lastName := "Doe"
	println(firstName, lastName)
	fullName := "John Doe \t(alias \"Foo\")\n"

	println(fullName)

	var integer16 int16 = 127
	var integer32 int32 = 32767
	println(int32(integer16) + integer32)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	println(i, s)
}