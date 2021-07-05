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
