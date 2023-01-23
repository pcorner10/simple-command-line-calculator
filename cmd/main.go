package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pcorner10/simple-command-line-calculator/operations/div"
	"github.com/pcorner10/simple-command-line-calculator/operations/multiply"
	"github.com/pcorner10/simple-command-line-calculator/operations/substract"
	"github.com/pcorner10/simple-command-line-calculator/operations/sums"
)

func main() {
	fmt.Println("Project running...")

	var result float64
	var err error
	var succes bool = true

	args := os.Args[1:]

	if len(args) < 2 {
		succes = false
		fmt.Println("Usage: go run cmd/main.go [operation] [list]")
	}

	operation := args[0]

	listNumbers := stringsToFloats(args[1:])

	switch operation {
	case "div":
		result, err = div.Division(listNumbers)
	case "sum":
		result, err = sums.Sums(listNumbers)
	case "subs":
		result, err = substract.Substract(listNumbers)
	case "multi":
		result, err = multiply.Multiply(listNumbers)
	default:
		succes = false
		fmt.Println("Invalid input: Operation unknown")
		fmt.Println("Can be: div, sum, subs or multi")
	}

	if err != nil {
		succes = false
		fmt.Println(err)
	}

	if succes {
		fmt.Printf("The result of: %s is %f", operation, result)
	}

}

func stringsToFloats(strings []string) []float64 {
	outPut := []float64{}

	for _, v := range strings {
		floatValue, _ := strconv.ParseFloat(v, 64)
		outPut = append(outPut, floatValue)
	}

	return outPut
}
