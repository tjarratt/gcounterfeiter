package invocations_test

import (
	"github.com/tjarratt/gcounterfeiter/fixtures/fixturesfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/gcounterfeiter/invocations"
)

var _ = Describe("counting invocations", func() {
	var fake *fixturesfakes.FakeExample

	BeforeEach(func() {
		fake = new(fixturesfakes.FakeExample)
	})

	It("should be zero initially", func() {
		Expect(CountTotalInvocations(fake.Invocations())).To(Equal(0))
	})

	Describe("when a single function is invoked", func() {
		BeforeEach(func() {
			fake.Something()
		})

		It("should be one", func() {
			Expect(CountTotalInvocations(fake.Invocations())).To(Equal(1))
		})

		Context("when the function is invoked again", func() {
			BeforeEach(func() {
				fake.Something()
			})

			It("should be two", func() {
				Expect(CountTotalInvocations(fake.Invocations())).To(Equal(2))
			})
		})

		Context("when a completely different function is invoked", func() {
			BeforeEach(func() {
				fake.TakesAParameter("")
			})

			It("should be two", func() {
				Expect(CountTotalInvocations(fake.Invocations())).To(Equal(2))
			})
		})
	})
})
