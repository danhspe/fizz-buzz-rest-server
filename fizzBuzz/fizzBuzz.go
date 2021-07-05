package fizzBuzz

// FizzBuzz returns a list of strings with numbers from 1 to limit, where:
// all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2,
// all multiples of int1 and int2 are replaced by str1str2.
func FizzBuzz(int1 uint, int2 uint, limit uint, str1 string, str2 string) []string {
	var result []string

	if limit == 0 || int1 == 0 || int2 == 0 {
		return result
	}

	for i := uint(1); i <= limit; i++ {
		if i%int1 == 0 {
			result = append(result, str1)
		}
	}

	return result
}
