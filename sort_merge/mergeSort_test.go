package sort_merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				arr: []int{2,4,1,2,4,7,5,6,8,3,9},
			},
			want: []int{1,2,2,3,4,4,5,6,7,8,9},
		},
		{
			name: "",
			args: args{
				arr: []int{2,1},
			},
			want: []int{1,2},
		},
		{
			name: "",
			args: args{
				arr: []int{1,1},
			},
			want: []int{1,1},
		},
		{
			name: "",
			args: args{
				arr: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
