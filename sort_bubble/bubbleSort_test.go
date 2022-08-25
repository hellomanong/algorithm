package sort_bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "",
			args: args{
				arr: []int64{1,2,3,6,5,4},
			},
			want: []int64{1,2,3,4,5,6},
		},
		{
			name: "",
			args: args{
				arr: []int64{6,5,4,3,2,1},
			},
			want: []int64{1,2,3,4,5,6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.arr)

			ok := reflect.DeepEqual(tt.args.arr, tt.want)
			if !ok {
				t.Errorf("---------")
			}

		})
	}
}
