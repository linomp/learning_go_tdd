package hello

import "fmt"

// Domain code
const spanishCode = "Spanish"
const frenchCode = "French"

// Hello greets a person or the world! (public function)
func Hello(name string, langCode string) string {
	if name == "" {
		name = "world"
	}

	return getPrefix(langCode) + name
}

// declaring a named return value, gets displayed in Go doc
// lowercase: private function
func getPrefix(langCode string) (greetingPrefix string) {

	switch langCode {
	case spanishCode:
		greetingPrefix = "Hola, "
	case frenchCode:
		greetingPrefix = "Bonjour, "
	default:
		greetingPrefix = "Hello, "
	}

	// will return whatever the named return value is set to
	// (gets initialized at corresponding zero value)
	return
}

// Outside world, side effects
func main() {
	fmt.Println(Hello("world", ""))
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
