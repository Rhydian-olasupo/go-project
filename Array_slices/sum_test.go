package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Sum(numbers)
	want := 45

	if got != want {
		t.Errorf("got %v want %v, given %v ", got, want, numbers)
	}
}

func TestSlice(t *testing.T) {
	t.Run("collection of any size (slices)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %v want %v , given %v", got, want, numbers)
		}
	})

	t.Run("Summation of different slices", func(t *testing.T) {
		myslices1, myslices2 := []int{1, 2, 3}, []int{1, 2, 3, 4}
		got := SumAll(myslices1, myslices2)
		want := []int{6, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v , given %v and %v", got, want, myslices1, myslices2)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumTails([]int{1, 2, 6, 6}, []int{1, 2, 9, 10})
	want := []int{14, 21}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
