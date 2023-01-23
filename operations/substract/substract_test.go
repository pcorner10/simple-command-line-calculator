package substract

import "testing"

func TestSubstract(t *testing.T) {
	tests := []struct{
		name string
		input []float64
		want float64
	}{
		{"Should be -10", []float64{5,2,3}, -10},
		{"Should be 0", []float64{0,0,0}, 0},
		{"Should be -2", []float64{-10,2,10}, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := Substract(tt.input)

			if err != nil {
				t.Errorf("Error ocurr during substract: %s", err)
			}

			if ans != tt.want {
				t.Errorf("got %f, want %f", ans, tt.want)
			}
		})
	}
}
