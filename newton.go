package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, g float64) float64 {
	var ans float64
	for i := 0; i < 10; i++ {
		ans = g - ((g*g)-x)/(2*g)
		g = ans
		//fmt.Println(g)
	}
	return ans
}

func main() {
	fmt.Println(Sqrt(612, 10))
	fmt.Println(math.Sqrt(612))
}
