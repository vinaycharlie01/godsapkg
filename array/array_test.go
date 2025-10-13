package array_test

import (
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
