package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var from, to string
	flag.StringVar(&from, "from", "", "Unit of input")
	flag.StringVar(&to, "to", "", "Unit of output")
	flag.Parse()

	fmt.Println("os.Args:", os.Args)

	fmt.Println("flag.Args():", flag.Args())
	if from == "" || to == "" {
		fmt.Printf("Usage: go run main.go -from=feet -to=meters 25")
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Printf("Usage: go run main.go -from=feet -to=meters 25")
		return
	}

	var input, output float64
	// flag.Args(): return non flag arugments
	numScanned, err := fmt.Sscanf(flag.Args()[0], "%f", &input)
	if numScanned != 1 || err != nil {
		fmt.Println("Cannot scan the value to convert. flag.Args() is:", flag.Args())
		return
	}

	switch {
	case from == "meters" && to == "feet":
		output = input / 0.3048
	case from == "feet" && to == "meters":
		output = 0.3048 * input
	case from == "meters" && to == "inches":
		output = input / 0.0254
	case from == "inches" && to == "meters":
		output = 0.00254 * input
	}

	fmt.Printf("%f %s are %f %s\n", input, from, output, to)
}
