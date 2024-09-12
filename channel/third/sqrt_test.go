package sqrt_test

import (
	sq "golang_learning/channel/third"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	// Test cases
	tests := []struct {
		input    float64
		expected float64
	}{
		{4, 2},
		{9, 3},
		{-4, 0},
		{16, 4},
	}

	for _, test := range tests {
		result := sq.Sqrt(test.input)
		if math.Abs(float64(result)-float64(test.expected)) > 1e-9 {
			t.Errorf("Expected %v for sqrt(%v), got %v", test.expected, test.input, result)
		}

	}
}
