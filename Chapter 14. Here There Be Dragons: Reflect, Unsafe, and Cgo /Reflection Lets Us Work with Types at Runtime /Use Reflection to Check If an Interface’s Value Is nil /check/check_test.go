package check

import (
	"testing"
)

func Test_hasNoValue(t *testing.T) {
	var i any
	var p *int
	var d int
	var f any = p
	type args struct {
		i any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty interface", args: args{i}, want: true},
		{name: "empty ptr", args: args{p}, want: true},
		{name: "empty value", args: args{d}, want: false},
		{name: "empty interface no value", args: args{f}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasNoValue(tt.args.i); got != tt.want {
				t.Errorf("hasNoValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
