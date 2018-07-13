package main

// Repeat takes a string and repeats it
func Repeat(input string, repeatCount int) (repeated string) {
	for index := 0; index < repeatCount; index++ {
		repeated += input
	}

	return
}

func main() {}
