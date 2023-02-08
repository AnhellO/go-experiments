package main

import "testing"

func TestSumTableDriven(t *testing.T) {
	var intsTests = []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 10},
			25,
		},
	}

	for _, test := range intsTests {
		if output := SumSliceValues(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, output: {}", test.input, test.expected, output)
		}
	}

	var floatsTests = []struct {
		input    []float64
		expected float64
	}{
		{
			[]float64{11.2, 2.2, 3.3, 4.4, 5.6, 10},
			36.7,
		},
	}

	for _, test := range floatsTests {
		if output := SumSliceValues(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, output: {}", test.input, test.expected, output)
		}
	}
}
