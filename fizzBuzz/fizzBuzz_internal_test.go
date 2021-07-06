package fizzBuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	dividend uint
	divisor  uint
	result   bool
}

func TestFizzBuzz_isDivisibleBy(t *testing.T) {
	tests := []struct {
		args
	}{
		{args: args{1, 0, false}},
		{args: args{1, 1, true}},
		{args: args{2, 1, true}},
		{args: args{2, 2, true}},
		{args: args{15, 2, false}},
		{args: args{15, 3, true}},
		{args: args{15, 5, true}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			result := isDivisibleBy(tt.dividend, tt.divisor)
			assert.Equal(t, tt.result, result, "Expected %t - got %t", tt.result, result)
		})
	}
}
