package generics_test

import (
	"github.com/stretchr/testify/require"
	"goKnowledge/generics"
	"testing"
)

func TestSumIntegers(t *testing.T) {
	require.Equal(t,
		generics.SumIntegers(map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
		}),
		10,
	)
}

func TestSumFloats(t *testing.T) {
	require.Equal(t,
		generics.SumFloats(map[string]float64{
			"one":   1.0,
			"two":   2.0,
			"three": 3.0,
			"four":  4.0,
		}),
		10.0,
	)
}

func TestSumIntegersOrFloats(t *testing.T) {
	t.Run("sum integers", func(t *testing.T) {
		require.Equal(t,
			generics.SumIntegersOrFloats(map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
			}),
			10,
		)
	})

	t.Run("sum floats", func(t *testing.T) {
		require.Equal(t,
			generics.SumIntegersOrFloats(map[string]float64{
				"one":   1.0,
				"two":   2.0,
				"three": 3.0,
				"four":  4.0,
			}),
			10.0,
		)
	})
}

func TestSumNumbers(t *testing.T) {
	t.Run("sum integer numbers", func(t *testing.T) {
		require.Equal(t,
			generics.SumNumbers(map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
			}),
			10,
		)
	})

	t.Run("sum float numbers", func(t *testing.T) {
		require.Equal(t,
			generics.SumNumbers(map[string]float64{
				"one":   1.0,
				"two":   2.0,
				"three": 3.0,
				"four":  4.0,
			}),
			10.0,
		)
	})
}
