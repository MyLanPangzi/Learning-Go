package table

import "testing"

func TestDoMath(t *testing.T) {
	type args struct {
		x  int
		y  int
		op string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		//	{"addition", 2, 2, "+", 4, ""},
		//    {"subtraction", 2, 2, "-", 0, ""},
		//    {"multiplication", 2, 2, "*", 4, ""},
		//    {"division", 2, 2, "/", 1, ""},
		//    {"bad_division", 2, 0, "/", 0, `division by zero`},
		{"2+2=4", args{2, 2, "+"}, 4, false},
		{"2*2=4", args{2, 2, "*"}, 4, false},
		{"2/2=1", args{2, 2, "/"}, 1, false},
		{"2/0 error", args{2, 0, "/"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DoMath(tt.args.x, tt.args.y, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoMath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoMath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
