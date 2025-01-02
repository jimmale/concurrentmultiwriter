package concurrentmultiwriter

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestConcurrentMultiWriter_Write(t *testing.T) {
	type fields struct {
		writers []io.Writer
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmw := ConcurrentMultiWriter{
				writers: tt.fields.writers,
			}
			gotN, err := cmw.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestMultiWriter(t *testing.T) {
	type args struct {
		writers []io.Writer
	}
	tests := []struct {
		name string
		args args
		want ConcurrentMultiWriter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiWriter(tt.args.writers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intMin(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{
			name: "positive ints",
			args: args{ints: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "negative ints",
			args: args{ints: []int{-1, -2, -3, -4, -5}},
			want: -5,
		},
		{
			name: "positive ints with zero",
			args: args{ints: []int{0, 1, 2, 3, 4, 5}},
			want: 0,
		},
		{
			name: "negative ints with zero",
			args: args{ints: []int{0, -1, -2, -3, -4, -5}},
			want: -5,
		},
		{
			name: "just zero",
			args: args{ints: []int{0}},
			want: 0,
		},
		{
			name: "positive, negative, and zero",
			args: args{ints: []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5}},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intMin(tt.args.ints...); got != tt.want {
				t.Errorf("intMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wrapErrors(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "one nil error",
			args:    args{errs: []error{nil}},
			wantErr: false,
		},

		{
			name:    "four nil errors",
			args:    args{errs: []error{nil, nil, nil, nil}},
			wantErr: false,
		},

		{
			name:    "empty error slice",
			args:    args{errs: []error{}},
			wantErr: false,
		},

		{
			name:    "2 nil errors 2 actual errors",
			args:    args{errs: []error{nil, fmt.Errorf("one"), fmt.Errorf("two"), nil}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := wrapErrors(tt.args.errs...); (err != nil) != tt.wantErr {
				t.Errorf("wrapErrors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
