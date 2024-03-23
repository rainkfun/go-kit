package hash

import (
	"testing"
)

func TestXXH64Sum(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want uint64
	}{
		{
			name: "empty",
			in:   []byte{},
			want: 17241709254077376921,
		},
		{
			name: "abc",
			in:   []byte("abc"),
			want: 4952883123889572249,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := XXH64Sum(tt.in); err != nil {
				t.Errorf("XXH64Sum() error = %v", err)
			} else if got != tt.want {
				t.Errorf("XXH64Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXXH64SumString(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want string
	}{
		{
			name: "empty",
			in:   []byte{},
			want: "ef46db3751d8e999",
		},
		{
			name: "abc",
			in:   []byte("abc"),
			want: "44bc2cf5ad770999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := XXH64SumString(tt.in); err != nil {
				t.Errorf("XXH64SumString() error = %v", err)
			} else if got != tt.want {
				t.Errorf("XXH64SumString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXXH64SumBytes(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want []byte
	}{
		{
			name: "empty",
			in:   []byte{},
			want: []byte("ef46db3751d8e999"),
		},
		{
			name: "abc",
			in:   []byte("abc"),
			want: []byte("44bc2cf5ad770999"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := XXH64SumBytes(tt.in); err != nil {
				t.Errorf("XXH64SumBytes() error = %v", err)
			} else if string(got) != string(tt.want) {
				t.Errorf("XXH64SumBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
