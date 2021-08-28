package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
			region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
			region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
			region, continent = "California", "USA"
	default:
			region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

func main() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero..")
	case 1:
		fmt.Print("one..")
	case 2:
		fmt.Print("two..")
	default:
		fmt.Print("no match..")
	}

	fmt.Println("ok")

	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)

	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())

	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}