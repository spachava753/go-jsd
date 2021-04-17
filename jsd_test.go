package go_jsd

import "testing"

func TestCalculateJSD(t *testing.T) {
	type args struct {
		wfd WFD
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "simple",
			args: args{
				WFD{
					A: map[string]float64{
						"hi":    0.5,
						"there": 0.5,
					},
					B: map[string]float64{
						"hi":    0.5,
						"out":   0.25,
						"there": 0.25,
					},
				},
			},
			want: 0.3945111687006673,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateJSD(tt.args.wfd); got != tt.want {
				t.Errorf("CalculateJSD() = %v, want %v", got, tt.want)
			}
		})
	}
}
