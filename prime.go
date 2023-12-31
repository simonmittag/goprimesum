package goprimesum

import (
	"fmt"
	"math"
)

type PrimeCandidate struct{}

func (p *PrimeCandidate) CountTo100() {
	a := make([]int, 100)
	for i, _ := range a {
		a[i] = i + 1
	}
	for _, ai := range a {
		p1 := PrimeCandidate{}
		if p1.IsPrime(ai) {
			fmt.Printf("\nPrime: %v", ai)
		}
	}
}

func (p *PrimeCandidate) IsPrime(n int) bool {
	// 1 is not a prime
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
