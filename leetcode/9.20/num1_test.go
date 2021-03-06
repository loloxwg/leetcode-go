package __20

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			"在排序数组中查找元素的第一个和最后一个位置",
			args{
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			[]int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"在排序数组中查找元素的第一个和最后一个位置",
			args{
				[]int{5, 7, 7, 8, 8, 10},
				8,
			},
			[]int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange2() = %v, want %v", got, tt.want)
			}
		})
	}
}
