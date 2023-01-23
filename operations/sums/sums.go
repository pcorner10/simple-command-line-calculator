package sums

func Sums(values []float64) (float64, error) {

	var outPut float64

	for _, v := range values {

		outPut = outPut + v
	}

	return outPut, nil

}
