package funk

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 2, 3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NextPermutation(tt.args.nums); (err != nil) != tt.wantErr {
				t.Errorf("NextPermutation() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				fmt.Println(tt.args.nums)
			}
		})
	}
}
