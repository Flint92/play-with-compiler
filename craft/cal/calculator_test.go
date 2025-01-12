package cal

import "testing"

func TestEval(t *testing.T) {
	tests := []struct {
		script string
		want   int64
	}{
		{"1+2", 3},
		{"1+2+3", 6},
		{"1+2+3*4", 15},
		{"1+4/2", 3},
		{"4/2-1", 1},
		{"4%2+(1+2)", 3},
	}

	for _, tt := range tests {
		_, got := Eval(tt.script)
		if got != tt.want {
			t.Errorf("Eval(%s) got = %d, want %d", tt.script, got, tt.want)
		}
	}

	err, _ := Eval("1+")
	if err == nil {
		t.Errorf("Eval(1+) got = nil, want error")
	}
}
