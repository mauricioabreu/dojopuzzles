package main

import (
	"reflect"
	"testing"
)

func TestWithOneInterval(t *testing.T) {
	result := GetIntervals([]int{100, 101})
	if !reflect.DeepEqual(result, []string{"100-101"}) {
		t.Errorf("Should be ['100-101'] but got %v", result)
	}
}

func TestWithTwoIntervals(t *testing.T) {
	result := GetIntervals([]int{100, 101, 102, 103, 104, 105})
	if !reflect.DeepEqual(result, []string{"100-105"}) {
		t.Errorf("Should be ['100-105'] but got %v", result)
	}
}

func TestWithoutValues(t *testing.T) {
	result := GetIntervals([]int{})
	if !reflect.DeepEqual(result, []string{}) {
		t.Errorf("Should be [] but got %v", result)
	}
}

func TestWithMultipleIntervals(t *testing.T) {
	result := GetIntervals([]int{100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150})
	if !reflect.DeepEqual(result, []string{"100-105", "110-111", "113-115", "150"}) {
		t.Errorf("Should be [] but got %v", result)
	}
}
