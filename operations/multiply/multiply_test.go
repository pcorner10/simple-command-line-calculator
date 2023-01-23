package multiply

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct{
		name string
		input []float64
		want float64
	} {
		{"Should be 1", []float64{1.0, 1.0}, 1.0},
		{"Should be 0", []float64{1.0, 0.0}, 0.0},
		{"Should be minus 1", []float64{-1.0, 1.0}, -1.0},
		{"Should be one", []float64{-1.0, -1.0}, 1.0},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			ans, err := Multiply(tt.input)

			if err != nil {
				t.Errorf("Error occurr during the multiply: %s", err)
			}

			if ans != tt.want {
				t.Errorf("got %f; want %f", ans, tt.want)
			}
		})
		
	}
}
