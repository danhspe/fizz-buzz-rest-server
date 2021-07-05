package fizzBuzz

import "strconv"

// FizzBuzz returns a list of strings with numbers from 1 to limit, where:
// all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2,
// all multiples of int1 and int2 are replaced by str1str2.
func FizzBuzz(int1 uint, int2 uint, limit uint, str1 string, str2 string) []string {
	var result []string

	if limit == 0 || int1 == 0 || int2 == 0 {
		return result
	}

	for i := uint(1); i <= limit; i++ {
		replacement := strconv.Itoa(int(i))
		if numberIsDivisibleOnlyByInt1(i, int1, int2) {
			replacement = str1
		} else if numberIsDivisibleOnlyByInt2(i, int1, int2) {
			replacement = str2
		} else if numberIsDivisibleByInt1AndInt2(i, int1, int2) {
			replacement = str1 + str2
		}
		result = append(result, replacement)
	}

	return result
}

func numberIsDivisibleOnlyByInt1(number uint, int1 uint, int2 uint) bool {
	return number%int1 == 0 && number%int2 != 0
}

func numberIsDivisibleOnlyByInt2(number uint, int1 uint, int2 uint) bool {
	return number%int1 != 0 && number%int2 == 0
}

func numberIsDivisibleByInt1AndInt2(number uint, int1 uint, int2 uint) bool {
	return number%int1 == 0 && number%int2 == 0
}
