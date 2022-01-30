package ttal_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	ttal "test"
)

func TestTrimMessage(t *testing.T) {
	type args struct {
		message string
		K       int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"doesn't include part words", args{"Codility We test coders", 14}, "Codility We"},
		{"doesn't trim if don't need to", args{"Why not   ", 100}, "Why not"},
		{"Trims space", args{"The quick brown fox jumps over the lazy dog", 39}, "The quick brown fox jumps over the lazy"},
		{"crops", args{"To crop or not to crop", 21}, "To crop or not to"},
		{"includes spaces", args{"To crop or not      to crop", 21}, "To crop or not"},
		{"handles empty", args{"", 21}, ""},
		{"returns empty if K 0", args{"soemthing", 0}, ""},
		{"returns empty if K -1", args{"soemthing", -1}, ""},
		{"long word", args{"supercalifragilisticexpiolodocious", 10}, ""},
		{"long word", args{"a supercalifragilisticexpiolodocious", 10}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ttal.TrimMessage(tt.args.message, tt.args.K)
			require.Equal(t, tt.want, got)
		})
	}
}
