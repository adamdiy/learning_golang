package kata

import (
"math"
"fmt"
)

func EquableTriangle(a, b, c int) bool {
    p := (a+b+c)/2
    area := math.Pow(float64(p*(p-a)*(p-b)*(p-c)), 0.5)
    fmt.Print(area)
    if area == float64(a+b+c) {
      return true
    } else {
      return false
    }
}