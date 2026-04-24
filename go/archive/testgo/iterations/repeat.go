package iterations

const repeatCount = 5

// Repeat takes a character and returns a string with that character repeated 5 times
func Repeat(character string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}
