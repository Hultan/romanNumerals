package roman

import (
	"testing"
)

func TestNumberToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"14 = XIV", args{14}, "XIV"},
		{"79 = LXXIX", args{79}, "LXXIX"},
		{"225 = CCXXV", args{225}, "CCXXV"},
		{"845 = DCCCXLV", args{845}, "DCCCXLV"},
		{"2022 = MMXXII", args{2022}, "MMXXII"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRoman(tt.args.num); got != tt.want {
				t.Errorf("ToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
