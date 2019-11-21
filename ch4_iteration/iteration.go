package iteration

// Repeat repeats a character
func Repeat(character string, length int) string {

	// just declare, but don't initialize
	var rep string

	for i := 0; i < length; i++ {
		rep += character
	}

	return rep
}
