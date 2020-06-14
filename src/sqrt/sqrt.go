package main

import (
	"fmt"
	"math"
)

func sqrt() string {
	x := 0.0001
	for i := 0; i < 10; i++ {
		x += math.Sqrt(x)
	}
	return "Code.education Rocks!"
}

func main() {
	fmt.Println(sqrt())
}
