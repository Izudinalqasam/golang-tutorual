package testing

import (
	"fmt"
	"math"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return math.Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val := math.Sqrt(2)

	//if err != nil {
	//	t.Fatalf("error in calculation - %s", err)
	//}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad Value - %f", val)
	}
}

// Table driven test (test multiple input)
type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out := math.Sqrt(tc.value)

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}
