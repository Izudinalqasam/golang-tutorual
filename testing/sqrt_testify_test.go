package testing

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSimpleTestify(t *testing.T) {
	val := math.Sqrt(2)
	require.InDelta(t, 1.414214, val, 0.001)
}
