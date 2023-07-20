package goprimesum_test

import (
	"github.com/simonmittag/goprimesum"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoprimesum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goprimesum Suite")
}

var _ = Describe("PrimeCandidates", func() {
	var primeCandidate goprimesum.PrimeCandidate

	BeforeEach(func() {
		primeCandidate = goprimesum.PrimeCandidate{}
	})

	Describe("Less Than 10", func() {
		Describe("Special Cases", func() {
			Context("1", func() {
				It("is not a prime because it's not divisible by itself AND 1", func() {
					Expect(primeCandidate.IsPrime(1)).To(BeFalse())
				})
			})
			Context("3", func() {
				It("is a prime", func() {
					Expect(primeCandidate.IsPrime(1)).To(BeFalse())
				})
			})
		})
	})
	Describe("Greater Than 10", func() {
		Context("11", func() {
			It("is a prime", func() {
				Expect(primeCandidate.IsPrime(11)).To(BeTrue())
			})
		})
		Context("13", func() {
			It("is a prime", func() {
				Expect(primeCandidate.IsPrime(13)).To(BeTrue())
			})
		})
	})
})
