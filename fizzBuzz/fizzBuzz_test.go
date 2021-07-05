package fizzBuzz_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.io/fizz-buzz-rest-server/fizzBuzz"
)

type args struct {
	int1  uint
	int2  uint
	limit uint
	str1  string
	str2  string
}

func TestFizzBuzz_WhenLimitIsLessThanOne_ThenReturnEmptySlice(t *testing.T) {
	result := fizzBuzz.FizzBuzz(0, 0, 0, "", "")
	assert.Empty(t, result, fmt.Sprintf("Expected empty slice - got %v", fmt.Sprintf("%v", result)))
}

func TestFizzBuzz_WhenLimitIsNotLessThanOne_AndOtherArgumentsAreValid_ThenDoNotReturnEmptySlice(t *testing.T) {
	result := fizzBuzz.FizzBuzz(1, 1, 1, "", "")
	assert.NotEmpty(t, result, fmt.Sprintf("Expected non-empty slice - got %v", fmt.Sprintf("%v", result)))
}

func TestFizzBuzz_WhenInt1orInt2IsZero_ThenReturnEmptySlice(t *testing.T) {
	var result []string

	result = fizzBuzz.FizzBuzz(0, 0, 1, "", "")
	assert.Empty(t, result, fmt.Sprintf("Expected empty slice - got %v", fmt.Sprintf("%v", result)))

	result = fizzBuzz.FizzBuzz(1, 0, 1, "", "")
	assert.Empty(t, result, fmt.Sprintf("Expected empty slice - got %v", fmt.Sprintf("%v", result)))

	result = fizzBuzz.FizzBuzz(0, 1, 1, "", "")
	assert.Empty(t, result, fmt.Sprintf("Expected empty slice - got %v", fmt.Sprintf("%v", result)))
}

func TestFizzBuzz_ReplaceMultiplesOfInt1WithStr1(t *testing.T) {
	tests := []struct {
		args     args
		expected string
	}{
		{args: args{
			int1:  1,
			int2:  1,
			limit: 1,
			str1:  "-1",
			str2:  "",
		}, expected: "[-1]"},
		{args: args{
			int1:  1,
			int2:  1,
			limit: 10,
			str1:  "-1",
			str2:  "",
		}, expected: "[-1 -1 -1 -1 -1 -1 -1 -1 -1 -1]"},
	}

	for i, tt := range tests {
		name := fmt.Sprintf("Test %d", i)
		t.Run(name, func(t *testing.T) {

			result := fizzBuzz.FizzBuzz(tt.args.int1, tt.args.int2, tt.args.limit, tt.args.str1, tt.args.str2)

			resultString := fmt.Sprintf("%v", result)
			assert.Equalf(t, tt.expected, resultString, "Expected %s - got %s", tt.expected, resultString)

		})
	}
}
