package luhna

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty string", s: "", want: false},
		{name: "alphanum string", s: "asdasd9898", want: false},
		{name: "valid", s: "4561261212345467", want: true},
		{name: "invalid", s: "4561261212345464", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.s); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDigitsString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty string", s: "", want: false},
		{name: "digits string", s: "000178", want: true},
		{name: "alfanum string", s: "dfsf789", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDigitsString(tt.s); got != tt.want {
				t.Errorf("IsDigitsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runeIsNotDigit(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{name: "zero", r: '0', want: false},
		{name: "nine", r: '0', want: false},
		{name: "below zero", r: '0' - 1, want: true},
		{name: "above zero", r: '9' + 1, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeIsNotDigit(tt.r); got != tt.want {
				t.Errorf("runeIsNotDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		prefix string
		l      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "odd digits", args: args{"", 16}, want: true},
		{name: "even digits", args: args{"", 15}, want: true},
		{name: "len 1", args: args{"", 1}, want: true},
		{name: "with prefix", args: args{"123", 11}, want: true},
		{name: "with alfanum prefix", args: args{"123a", 11}, want: false},
		{name: "invalid len no prefix", args: args{"", 0}, want: false},
		{name: "invalid len with prefix", args: args{"123", 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.args.prefix, tt.args.l)
			valid := Validate(got)
			assert.Equalf(t, tt.want, valid, "expect Validate(%s) == %t", got, tt.want)
			if valid {
				assert.Equalf(t, tt.args.l, len(got), "required length %d, got len(%s)=%d", tt.args.l, got, len(got))
				if tt.args.prefix != "" {
					assert.Equalf(
						t,
						tt.args.prefix,
						got[:len(tt.args.prefix)],
						"expected prefix %s, got %s (%s)",
						tt.args.prefix,
						got[:len(tt.args.prefix)],
						got,
					)
				}
			}
		})
	}
}
