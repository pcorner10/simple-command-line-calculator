package div

import "testing"

func TestDivision(t *testing.T) {
	var tests = []struct {
		name  string
		input []float64
		want  float64
	}{
		{"Should be 1.0", []float64{24, 24}, 1.0},
		{"Should be 2.0", []float64{24, 12}, 2.0},
		{"Should be 5.0", []float64{20, 2, 2}, 5.0},
		{"Should be 0.5", []float64{-10, 20}, -0.5},
		{"Should be an error", []float64{-10, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := Division(tt.input)

			if err != nil {
				t.Errorf("Error occured while dividing the values: %v", err)
			}

			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}
