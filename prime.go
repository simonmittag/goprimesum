package goprimesum

import "math"

type Prime struct{}

func (p *Prime) countTo100() {

}

func (p *Prime) isPrime(n int) bool {
	nsq := math.Sqrt(float64(n))
	if n == 1 {
		return false
	}
	for i := 2; float64(i) <= nsq; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
