package div

import (
	"errors"
)

func Division(values []float64) (float64, error) {

	var outPut float64

	for i, v := range values {

		if v == 0 {
			return 0, errors.New("División por zero")
		}

		if i == 0 {
			outPut = v
		} else {
			outPut = outPut / v
		}
	
	}

	return outPut, nil

}
