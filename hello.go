package main

import "fmt"

// Domain code
const helloPrefix = "Hello, "

// Hello greets a person or the world!
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return helloPrefix + name
}

// Outside world, side effects
func main() {
	fmt.Println(Hello("world"))
}

// go build <file>
// go run <file>

/*
TDD:

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful (important because it's hard to work when tests fail without clear explanation)
- Write enough code to make the test pass
- Refactor


*/
