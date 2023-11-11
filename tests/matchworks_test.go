package tests

import (
	"chethan0707/basic_number_filtering/mathworks"
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	output := mathworks.EvenNumbers(input)

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestOddNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{1, 3, 5, 7, 9}

	output := mathworks.OddNumbers(input)

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestPrimeNumbers(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{2, 3, 5, 7}

	output := mathworks.PrimeNumbers(input)

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestOddPrimeNumberFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{3, 5, 7}

	output := mathworks.OddPrimeNumbers(input)

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestEvenMultiplesOf5NumberFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{10, 20}

	output := mathworks.EvenMultiplesOf5NumberFilter(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}

}

func TestOddMultiplesOfThreeGreaterThan10NumberFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{15}

	output := mathworks.MultiplesOf3NumerFilter(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, got %v", expected, output)
	}

}

func TestOddMultiplesOf3GreaterThan10NumberFilter(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{15}

	odd := func(n int) bool { return n%2 != 0 }
	multiplesOf := func(n int) mathworks.Condition { return func(m int) bool { return m%n == 0 } }
	multiplesOf3 := multiplesOf(3)
	// multiplesOf3 := func(n int) bool { return n%3 == 0 }

	greaterThanN := func(n int) mathworks.Condition { return func(m int) bool { return m > n } }
	greaterThan10 := greaterThanN(10)
	//greaterThan10 := func(n int) bool { return n > 10 }

	// when
	output := mathworks.Filter(numRange, odd, multiplesOf3, greaterThan10)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestMatchAllConditionsNumberFilter(t *testing.T) {
	// given
	odd := func(n int) bool { return n%2 != 0 }
	even := func(n int) bool { return !odd(n) }
	greaterThanN := func(n int) mathworks.Condition { return func(m int) bool { return m > n } }
	greaterThan5 := greaterThanN(5)
	multiplesOf := func(n int) mathworks.Condition { return func(m int) bool { return m%n == 0 } }
	multiplesOf3 := multiplesOf(3)
	lessThanN := func(n int) mathworks.Condition { return func(m int) bool { return m < n } }
	lessThan15 := lessThanN(15)

	tt := []struct {
		nums  []int
		conds []mathworks.Condition
		want  []int
	}{
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []mathworks.Condition{odd, greaterThan5, multiplesOf3},
			want:  []int{9, 15},
		},
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []mathworks.Condition{even, lessThan15, multiplesOf3},
			want:  []int{6, 12},
		},
		{
			nums:  []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
			conds: []mathworks.Condition{even, greaterThan5, multiplesOf3},
			want:  []int{54, 60, 66, 72, 78, 84, 90, 96},
		},
	}

	// when
	for _, tc := range tt {
		got := mathworks.Filter(tc.nums, tc.conds...)
		// then
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestMatchAnyConditionsNumberFilter(t *testing.T) {
	// given
	prime := func(n int) bool { return mathworks.IsPrimeNume(n) }
	greaterThanN := func(n int) mathworks.Condition { return func(m int) bool { return m > n } }
	greaterThan15 := greaterThanN(15)
	multiplesOf := func(n int) mathworks.Condition { return func(m int) bool { return m%n == 0 } }
	multiplesOf5 := multiplesOf(5)
	multiplesOf3 := multiplesOf(3)
	lessThanN := func(n int) mathworks.Condition { return func(m int) bool { return m < n } }
	lessThan6 := lessThanN(6)
	tt := []struct {
		nums  []int
		conds []mathworks.Condition
		want  []int
	}{
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []mathworks.Condition{prime, greaterThan15, multiplesOf5},
			want:  []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20},
		},
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []mathworks.Condition{lessThan6, multiplesOf3}, // less than 6, multiple of 3
			want:  []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18},
		},
	}

	// when
	for _, tc := range tt {
		got := mathworks.FilterAny(tc.nums, tc.conds...)
		// then
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
