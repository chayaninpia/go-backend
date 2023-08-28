package tox

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: &[]string{"test"},
			},
			want: []string{"test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice[string](tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestStruct struct {
	Name    string
	Number  int
	Boolean bool
}

func TestSliceOther(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want []TestStruct
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: &[]TestStruct{
					{
						Name:    "123",
						Number:  123,
						Boolean: true,
					},
				},
			},
			want: []TestStruct{
				{
					Name:    "123",
					Number:  123,
					Boolean: true,
				},
			},
		},
		{
			name: "test nil",
			args: args{
				s: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice[TestStruct](tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
