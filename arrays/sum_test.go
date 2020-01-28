package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum of 5 numbers", func(t *testing.T) {
		numbers := []int{0, 1, 2, 3, 4}
		want := 10
		got := Sum(numbers)
		if want != got {
			t.Errorf("want: %d got: %d given: %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 1}, []int{4, 2})
	want := []int{2, 6}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %d got: %d", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	tailsAssertion := func(t *testing.T, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got %v", want, got)
		}
	}
	t.Run("Run without empty sums", func(t *testing.T) {
		got := SumAllTails([]int{4, 9}, []int{6, 2})
		want := []int{9, 2}
		tailsAssertion(t, want, got)
	})
	t.Run("Run with empty sums", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{6, 2, 9})
		want := []int{0, 11}
		tailsAssertion(t, want, got)
	})
}
