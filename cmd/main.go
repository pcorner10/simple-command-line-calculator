package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pcorner10/simple-command-line-calculator/operations/div"
)

func main() {
	fmt.Println("Project running...")

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: go run cmd/main.go [operation] [list]")
	}

	operation := args[0]

	listNumbers := stringsToFloats(args[1:])

	var result float64
	var err error

	switch operation {
	case "div":
		result, err = div.Division(listNumbers)
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The result of: %s is %f", operation, result)

}

func stringsToFloats(strings []string) []float64 {
	outPut := []float64{}

	for _, v := range strings {
		floatValue, _ := strconv.ParseFloat(v, 64)
		outPut = append(outPut, floatValue)
	}

	return outPut
}
