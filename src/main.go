package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Standard     = "STANDARD"
	Special      = "SPECIAL"
	Rejected     = "REJECTED"
	MinDimension = 1
	MinMass      = 0
	MaxDimension = 150
	MaxMass      = 20
	MaxVolume    = 1000000
)

// should run the tests
func main() {
	// Arguments
	if len(os.Args) > 1 {
		fmt.Println("First user argument:", os.Args[1])
	} else {
		fmt.Println("No user arguments provided.")
	}

	if len(os.Args) != 5 {
		log.Fatal("Usage: go run main.go <width> <height> <length> <mass>")
	}

	w := os.Args[1]
	h := os.Args[2]
	l := os.Args[3]
	m := os.Args[4]

	width, err := strconv.Atoi(w)
	if err != nil {
		log.Fatal("Invalid width:", w)
	}
	height, err := strconv.Atoi(h)
	if err != nil {
		log.Fatal("Invalid height:", h)
	}
	length, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal("Invalid length:", l)
	}
	mass, err := strconv.Atoi(m)
	if err != nil {
		log.Fatal("Invalid mass:", m)
	}

	result, err := sort(width, height, length, mass)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(result)
}

/*
Rules:
Sort the packages using the following criteria:

- A package is **bulky** if its volume (Width x Height x Length) is greater than or equal to 1,000,000 cm³ or when one of its dimensions is greater or equal to 150 cm.
- A package is **heavy** when its mass is greater or equal to 20 kg.

You must dispatch the packages in the following stacks:

- **STANDARD**: standard packages (those that are not bulky or heavy) can be handled normally.
- **SPECIAL**: packages that are either heavy or bulky can't be handled automatically.
- **REJECTED**: packages that are **both** heavy and bulky are rejected.

Questions:
Is it really okay to have a length of 1000000 cm?
I am presuming that the mass is rounded so the mass can be 0. Is that okay?
I am presuming that no package would have a dimension of 0. Is that okay?
*/
func sort(width, height, length, mass int) (string, error) {

	if width < MinDimension || height < MinDimension || length < MinDimension || mass < MinMass {
		return "", fmt.Errorf("Dimensions must be positive integers and mass must be non-negative.")
	}

	bulky := (width > MaxDimension || height > MaxDimension || length > MaxDimension || width*height*length > MaxVolume)
	heavy := mass > MaxMass

	if !bulky && !heavy {
		return Standard, nil
	}

	if bulky && heavy {
		return Rejected, nil
	}

	// Either bulky or heavy, so Special
	return Special, nil
}
