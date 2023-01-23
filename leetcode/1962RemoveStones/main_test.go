package main

import "testing"

func Test_minStoneSum(t *testing.T) {
	type args struct {
		piles []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_1", args{[]int{5, 4, 9},2},12},
		{"test_2", args{[]int{4,3,6,7},3},12},
		{"test_3", args{[]int{1},1000},1},
		{"test_4", args{[]int{7481,7973,3635,5320,2721,4394,7861},10},13867},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStoneSum(tt.args.piles, tt.args.k); got != tt.want {
				t.Errorf("minStoneSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
