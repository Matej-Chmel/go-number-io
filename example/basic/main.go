package main

import (
	"fmt"
	"os"

	nio "github.com/Matej-Chmel/go-number-io"
)

func main() {
	file, err := os.Open("data/int/3D.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	// data, err := nio.Read3D[int32](file) // Concrete approach
	data, err := nio.Read[[][][]int32](file) // Dynamic approach

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(data)
}
