package ttal_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	ttal "test"
)

func TestGetMinRides(t *testing.T) {
	type args struct {
		P []int
		S []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 4, 1}, []int{1, 5, 1}}, 2},
		{"", args{[]int{4, 4, 2, 4}, []int{5, 5, 2, 5}}, 3},
		{"", args{[]int{2, 3, 4, 2}, []int{2, 5, 7, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ttal.GetMinRides(tt.args.P, tt.args.S)
			require.Equal(t, tt.want, got)
		})
	}
}
