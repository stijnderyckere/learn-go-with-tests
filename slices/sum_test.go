package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 2 slices", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{1, 9})
		want := []int{3, 10}

		//Go does not let you use equality operators with slices
		//reflect.DeepEqual is not type safe!!
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sum of 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
