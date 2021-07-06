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

		numberIsDivisibleByInt1 := isDivisibleBy(i, int1)
		numberIsDivisibleByInt2 := isDivisibleBy(i, int2)

		numberIsDivisibleOnlyByInt1 := numberIsDivisibleByInt1 && !numberIsDivisibleByInt2
		numberIsDivisibleOnlyByInt2 := !numberIsDivisibleByInt1 && numberIsDivisibleByInt2
		numberIsDivisibleByInt1AndInt2 := numberIsDivisibleByInt1 && numberIsDivisibleByInt2

		if numberIsDivisibleOnlyByInt1 {
			replacement = str1
		} else if numberIsDivisibleOnlyByInt2 {
			replacement = str2
		} else if numberIsDivisibleByInt1AndInt2 {
			replacement = str1 + str2
		}
		result = append(result, replacement)
	}

	return result
}

// Returns true if the remainder of dividend/divisor is zero, false otherwise.
func isDivisibleBy(dividend uint, divisor uint) bool {
	if divisor == 0 {
		return false
	}
	return dividend%divisor == 0
}
