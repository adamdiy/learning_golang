package kata

import (
	"math"
)

func Step(desiredGap, beginPoint, endPoint int) []int {
    answer := []int{}
    
    for i:=beginPoint; i<endPoint; i++ {
      if isPrime(i) {
        if isPrime(i+desiredGap) {
          return append(answer, i, i+desiredGap)
        }
      }
    }
    return nil;
}

// https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
func isPrime(value int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}