package race

import "testing"

func Test_getCounter(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "counter 5000", want: 5000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCounter(); got != tt.want {
				t.Errorf("getCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
