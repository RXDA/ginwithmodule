package sortfunc

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				[]int{6, 1, 2, 5, 4, 3, 9, 6, 10, 8},
			},
			[]int{1,2,3,4,5,6,6,8,9,10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quick(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksort() = %v, want %v", got, tt.want)
			}else{
				fmt.Println(got)
			}
		})
	}
}
