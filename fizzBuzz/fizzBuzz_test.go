package fizzBuzz_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.io/fizz-buzz-rest-server/fizzBuzz"
)

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
	var expectedResult = fmt.Sprintf("%v", []string{"-1"})

	result := fizzBuzz.FizzBuzz(1, 1, 1, "-1", "")

	resultString := fmt.Sprintf("%v", result)
	assert.Equalf(t, expectedResult, resultString, "Expected %s - got %s", expectedResult, resultString)
}
