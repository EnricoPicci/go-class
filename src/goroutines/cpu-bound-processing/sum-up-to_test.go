package main

import (
	"testing"
)

func TestSumUpTo_3(t *testing.T) {
	const upTo = 3
	const concurrent = 1
	const expected = 6
	result := sumUpTo(upTo, concurrent)

	// test the result
	if result != expected {
		t.Errorf("The result is %v instead of %v", result, expected)
	}
}

func TestSumUpTo_1000(t *testing.T) {
	const upTo = 1000
	const concurrent = 10
	const expected = 500500
	result := sumUpTo(upTo, concurrent)

	// test the result
	if result != expected {
		t.Errorf("The result is %v instead of %v", result, expected)
	}
}

func TestSumUpTo_1000000(t *testing.T) {
	const upTo = 1000000
	const concurrent = 1
	const expected = 500000500000
	result := sumUpTo(upTo, concurrent)

	// test the result
	if result != expected {
		t.Errorf("The result is %v instead of %v", result, expected)
	}
}

func TestCalcInterval_NoRemainder(t *testing.T) {
	const upTo = 30
	const concurrent = 10
	const expected = 3
	result := calcInterval(upTo, concurrent)

	// test the result
	if result != expected {
		t.Errorf("The result is %v instead of %v", result, expected)
	}
}

func TestCalcInterval_WithRemainder(t *testing.T) {
	const upTo = 30
	const concurrent = 11
	const expected = 2
	result := calcInterval(upTo, concurrent)

	// test the result
	if result != expected {
		t.Errorf("The result is %v instead of %v", result, expected)
	}
}

func TestCalcRanges_NoRemainder(t *testing.T) {
	const upTo = 30
	const concurrent = 10
	var expected = [][]int{{1, 3}, {4, 6}, {7, 9}, {10, 12}, {13, 15}, {16, 18}, {19, 21}, {22, 24}, {25, 27}, {28, 30}}
	result := calcRanges(upTo, concurrent)

	// test the result
	if len(result) != concurrent {
		t.Errorf("The result is %v instead of %v", len(result), concurrent)
	}
	for i, r := range result {
		if r[0] != expected[i][0] || r[1] != expected[i][1] {
			t.Errorf("The range %v is %v instead of %v", i, result[i], expected[i])
		}
	}
}

// TestCalcRanges_NoRemainder tests that the ranges are calculated correctly when dividing the range by the concurrent count
// returns a remainder.
func TestCalcRanges_WithRemainder_1(t *testing.T) {
	const upTo = 30
	const concurrent = 11
	var expected = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}, {13, 14}, {15, 16}, {17, 18}, {19, 20}, {21, 30}}
	result := calcRanges(upTo, concurrent)

	// test the result
	if len(result) != concurrent {
		t.Errorf("The result is %v instead of %v", len(result), concurrent)
	}
	for i, r := range result {
		if r[0] != expected[i][0] || r[1] != expected[i][1] {
			t.Errorf("The range %v is %v instead of %v", i, result[i], expected[i])
		}
	}
}

// TestCalcRanges_NoRemainder tests that the ranges are calculated correctly when dividing the range by the concurrent count
// returns a remainder.
func TestCalcRanges_WithRemainder_2(t *testing.T) {
	const upTo = 31
	const concurrent = 10
	var expected = [][]int{{1, 3}, {4, 6}, {7, 9}, {10, 12}, {13, 15}, {16, 18}, {19, 21}, {22, 24}, {25, 27}, {28, 31}}
	result := calcRanges(upTo, concurrent)

	// test the result
	if len(result) != concurrent {
		t.Errorf("The result is %v instead of %v", len(result), concurrent)
	}
	for i, r := range result {
		if r[0] != expected[i][0] || r[1] != expected[i][1] {
			t.Errorf("The range %v is %v instead of %v", i, result[i], expected[i])
		}
	}
}

func TestCalcRanges_1000000(t *testing.T) {
	const upTo = 1000000
	const concurrent = 10
	var expected = [][]int{
		{1, 100000},
		{100001, 200000},
		{200001, 300000},
		{300001, 400000},
		{400001, 500000},
		{500001, 600000},
		{600001, 700000},
		{700001, 800000},
		{800001, 900000},
		{900001, 1000000}}
	result := calcRanges(upTo, concurrent)

	// test the result
	if len(result) != concurrent {
		t.Errorf("The result is %v instead of %v", len(result), concurrent)
	}
	for i, r := range result {
		if r[0] != expected[i][0] || r[1] != expected[i][1] {
			t.Errorf("The range %v is %v instead of %v", i, result[i], expected[i])
		}
	}
}

func TestSumFromTo(t *testing.T) {
	const from = 1
	const to = 4
	var expected = 10

	results := make(chan int)

	go sumFromTo(from, to, results)

	// test the result
	result := <-results
	if result != expected {
		t.Errorf("The result is %v instead of %v", result, expected)
	}
}
