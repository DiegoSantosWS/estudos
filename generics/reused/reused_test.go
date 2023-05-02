package reused_test

import (
	"generics/reused"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestMaxWithoutGenerics(t *testing.T) {
	casesTests := []struct {
		Name     string
		Value    []int
		Expected int
	}{
		{
			Name:     "case 2: without generics",
			Value:    []int{1, 23, 41, 78, 36},
			Expected: 78,
		},
		{
			Name:     "case 2: without generics",
			Value:    []int{1, 23, 41, 78, 36},
			Expected: 78,
		},
	}

	for _, tc := range casesTests {
		t.Run(tc.Name, func(t *testing.T) {
			got := reused.MaxWithoutGenerics(tc.Value)
			if got != tc.Expected {
				t.Error("[ TESTS GO!! ] the expected value don't has match with the got")
				return
			}
		})
	}
}

type ValuesToTests[T constraints.Integer] struct {
	Value    []T
	ExpValue T
}

func TestMaxWithGenerics(t *testing.T) {
	casesTests := []struct {
		Name     string
		Value    ValuesToTests[int64]
		Expected ValuesToTests[int64]
	}{
		{
			Name: "case 2: without generics",
			Value: ValuesToTests[int64]{
				Value: []int64{1, 23, 41, 78, 36},
			},
			Expected: ValuesToTests[int64]{
				ExpValue: 78,
			},
		},
		{
			Name: "case 2: without generics",
			Value: ValuesToTests[int64]{
				Value: []int64{1, 23, 41, 36},
			},
			Expected: ValuesToTests[int64]{
				ExpValue: 41,
			},
		},
	}

	for _, tc := range casesTests {
		t.Run(tc.Name, func(t *testing.T) {
			got := reused.MaxWithGenerics(tc.Value.Value)
			assert.Equal(t, tc.Expected.ExpValue, got, tc.Name)
		})
	}
}
