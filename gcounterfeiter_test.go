package gcounterfeiter_test

import (
	"github.com/onsi/gomega/types"
	"github.com/tjarratt/gcounterfeiter/fixtures/fixturesfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tjarratt/gcounterfeiter"
)

var _ = Describe("HaveReceived", func() {
	var fake *fixturesfakes.FakeExample

	BeforeEach(func() {
		fake = new(fixturesfakes.FakeExample)
	})

	It("initially has no recorded invocations", func() {
		Expect(fake).ToNot(HaveReceived("Something"))
		Expect(fake).ToNot(HaveReceived("TakesAParameter"))
	})

	It("can verify that no function calls were made", func() {
		Expect(fake).ToNot(HaveReceived())
	})

	Context("after a function is called", func() {
		BeforeEach(func() {
			fake.Something()
		})

		It("should have recorded the invocation", func() {
			Expect(fake).To(HaveReceived("Something"))
		})

		It("should no longer report that no function calls were made", func() {
			Expect(fake).To(HaveReceived())
		})
	})

	Describe("failure messages", func() {
		var subject types.GomegaMatcher

		Context("when the user specifies a function by name", func() {
			BeforeEach(func() {
				subject = HaveReceived("Something")
			})

			It("should tell the user when it was not invoked", func() {
				subject.Match(fake)
				Expect(subject.FailureMessage(fake)).To(ContainSubstring("to have received 'Something'"))
			})
		})

		Context("when the user specifies no function", func() {
			BeforeEach(func() {
				subject = HaveReceived()
			})

			It("should tell the user when nothing was expected, but a function was invoked", func() {
				fake.Something()
				subject.Match(fake)
				Expect(subject.FailureMessage(fake)).To(ContainSubstring("to have received nothing, but it received 1 invocations"))
			})
		})
	})

	Describe("negated failure messages", func() {
		var subject types.GomegaMatcher

		Context("when the user specifies a function by name", func() {
			BeforeEach(func() {
				subject = HaveReceived("Something")
			})

			It("should tell the user when nothing was expected, but a function was invoked", func() {
				fake.Something()
				subject.Match(fake)
				Expect(subject.NegatedFailureMessage(fake)).To(ContainSubstring("to not have received 'Something', but it was invoked 1 time"))
			})
		})

		Context("when the user specifies no function", func() {
			BeforeEach(func() {
				subject = HaveReceived()
			})

			It("should tell the user when a function was expected, but none were invoked", func() {
				subject.Match(fake)
				Expect(subject.NegatedFailureMessage(fake)).To(ContainSubstring("to have received at least one invocation, but there were none"))
			})
		})
	})

	Context("when the user accidentally specifies too many arguments", func() {
		var matcher types.GomegaMatcher

		BeforeEach(func() {
			matcher = HaveReceived("whoops", "I", "accidentally")
		})

		It("should fail to match", func() {
			matches, err := matcher.Match(nil)
			Expect(matches).To(BeFalse())
			Expect(err.Error()).To(ContainSubstring("too many arguments"))
		})
	})

	Context("when the user accidentally writes an assertion against an invalid type", func() {
		var matcher types.GomegaMatcher

		BeforeEach(func() {
			matcher = HaveReceived("Something")
		})

		It("should fail to match and tell the user they goofed", func() {
			matches, err := matcher.Match(struct{}{})
			Expect(matches).To(BeFalse())
			Expect(err.Error()).To(ContainSubstring("does not conform to the 'InvocationRecorder' interface"))
		})
	})
})
