package goprimesum

import (
	"fmt"
	"math"
)

type Prime struct{}

func (p *Prime) countTo100() {
	a := make([]int, 100)
	for i, _ := range a {
		a[i] = i + 1
	}
	for _, ai := range a {
		p1 := Prime{}
		if p1.isPrime(ai) {
			fmt.Printf("\nPrime: %v", ai)
		}
	}
}

func (p *Prime) isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	nsq := math.Sqrt(float64(n))

	for i := 2; float64(i) <= nsq; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
