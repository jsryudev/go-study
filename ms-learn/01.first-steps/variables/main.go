var firstName string
var firstName, lastName string
var firstName, lastName string
var age int

var (
	firstName, lastName string
	age int
)

var (
	firstName string = "John"
	lastName string = "Doe"
	age int = 32
)

var (
	firstName = "John"
	lastName = "Doe"
	age = 32
)

var (
	firstName, lastName, age = "John", "Doe", 32
)

const HTTPStatusOK = 200

const (
	StatusOK = 0
	StatusConnectionReset = 1
	StatusOtherError = 2
)

func main() {
	firstName, lastName := "John", "Doe"
	age := 32
	println(firstName, lastName, age)
}