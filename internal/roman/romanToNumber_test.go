package roman

import (
	"testing"
)

func TestRomanToNumber(t *testing.T) {
	type args struct {
		roman string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"XIX = 14", args{"XIV"}, 14},
		{"LXXIX = 79", args{"LXXIX"}, 79},
		{"CCXXV = 225", args{"CCXXV"}, 225},
		{"DCCCXLV = 845", args{"DCCCXLV"}, 845},
		{"MMXXII = 2022", args{"MMXXII"}, 2022},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArabic(tt.args.roman); got != tt.want {
				t.Errorf("ToArabic() = %v, want %v", got, tt.want)
			}
		})
	}
}
