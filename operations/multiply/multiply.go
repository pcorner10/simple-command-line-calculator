package multiply

import "errors"

func Multiply(values []float64) (float64, error) {

	var outPut float64 

	if len(values) == 0 {
		return 0, errors.New("Empty list")
	}

	for i, v := range values {
		if i == 0 {
			outPut = v
		} else {
			outPut = outPut * v
		}
		
	}

	return outPut, nil
}