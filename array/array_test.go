package array_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/vinaycharlie01/godsapkg/array"
)

func Test_sumDivisibleByK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 2, 3, 3, 3, 3, 4},
			k:    2,
			want: 16,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: 0,
		},
		{
			name: "Example 3",
			nums: []int{4, 4, 4, 1, 2, 3},
			k:    3,
			want: 12,
		},
		{
			name: "empty slice",
			nums: []int{},
			k:    2,
			want: 0,
		},
		{
			name: "all frequencies divisible",
			nums: []int{2, 2, 4, 4, 6, 6},
			k:    2,
			want: 24, // 2*2 + 4*2 + 6*2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := array.SumDivisibleByK(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("sumDivisibleByK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "basic merge",
			nums1:    []int{1, 3, 5, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 4, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "nums1 empty",
			nums1:    []int{0, 0, 0},
			m:        0,
			nums2:    []int{2, 5, 7},
			n:        3,
			expected: []int{2, 5, 7},
		},
		{
			name:     "nums2 empty",
			nums1:    []int{1, 2, 3},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "interleaved elements",
			nums1:    []int{1, 4, 7, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 3, 8},
			n:        3,
			expected: []int{1, 2, 3, 4, 7, 8},
		},
		{
			name:     "duplicate elements",
			nums1:    []int{1, 2, 2, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 2, 3},
			n:        3,
			expected: []int{1, 2, 2, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := array.Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Merge() = %v, want %v", tt.nums1[:tt.m+tt.n], tt.expected)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		k        int
		expected []int
	}{
		{
			name:     "rotate by 0",
			items:    []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by 1",
			items:    []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			name:     "rotate by 2",
			items:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "rotate by length (no change)",
			items:    []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by more than length",
			items:    []int{1, 2, 3, 4, 5},
			k:        7,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "empty slice",
			items:    []int{},
			k:        3,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy so original test data isn’t mutated across cases
			input := slices.Clone(tt.items)
			array.Rotate(input, tt.k)

			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("Rotate(%v, %d) = %v; want %v", tt.items, tt.k, input, tt.expected)
			}
		})
	}
}
