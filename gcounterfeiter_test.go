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
		Expect(fake).ToNot(HaveReceived("TakesAParameter").With(Equal("anything-at-all")))
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

	Describe("argument matching", func() {
		BeforeEach(func() {
			fake.TakesAParameter("you-bet-it-does")
			fake.TakesAnInt(0)
			fake.TakesAUint64(0)
		})

		It("allows you to verify that the correct arguments were passed in", func() {
			Expect(fake).To(HaveReceived("TakesAParameter").With(Equal("you-bet-it-does")))
			Expect(fake).ToNot(HaveReceived("TakesAParameter").With(Equal("whoops")))
		})

		It("defaults to Equal() when no matcher is provided", func() {
			Expect(fake).To(HaveReceived("TakesAnInt").With(0))

			Expect(fake).ToNot(HaveReceived("TakesAUint64").With(0))
			Expect(fake).To(HaveReceived("TakesAUint64").With(uint64(0)))
		})

		It("can match when the function was invoked multiple times", func() {
			fake.TakesAParameter("but-will-it-blend?")
			Expect(fake).To(HaveReceived("TakesAParameter").With("but-will-it-blend?"))
			Expect(fake).To(HaveReceived("TakesAParameter").With("you-bet-it-does"))
			Expect(fake).ToNot(HaveReceived("TakesAParameter").With("surely-you're-joking"))
		})

		Context("when too many arguments are provided", func() {
			var subject types.GomegaMatcher

			BeforeEach(func() {
				subject = HaveReceived("TakesAParameter").With(Equal("you-bet-it-does")).AndWith(Equal("whoops"))
			})

			It("should tell the user they goofed", func() {
				ok, err := subject.Match(fake)
				Expect(ok).To(BeFalse())
				Expect(err.Error()).To(ContainSubstring("Too many arguments provided for 'TakesAParameter'. Expected 1 but received 2"))
			})
		})

		Context("when too few arguments are provided", func() {
			var subject types.GomegaMatcher

			BeforeEach(func() {
				subject = HaveReceived("TakesThreeParameters").With(Equal("whoops"))
			})

			It("should tell the user they goofed", func() {
				fake.TakesThreeParameters("", "", "")

				ok, err := subject.Match(fake)
				Expect(ok).To(BeFalse())
				Expect(err.Error()).To(ContainSubstring("Too few arguments provided for 'TakesThreeParameters'. Expected 3 but received 1"))
			})
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
				matched, err := subject.Match(fake)
				Expect(matched).To(BeFalse())
				Expect(err).ToNot(HaveOccurred())
				Expect(subject.FailureMessage(fake)).To(ContainSubstring("to have received at least one invocation, but it received 0"))
			})
		})

		Context("when the user specifies parameters to match", func() {
			BeforeEach(func() {
				subject = HaveReceived("TakesThreeParameters").
					With(Equal("a")).
					AndWith(Equal("b")).
					AndWith(Equal("c"))
			})

			It("should tell the user which parameter failed to match", func() {
				fake.TakesThreeParameters("a", "b", "whoops")
				subject.Match(fake)

				Expect(subject.FailureMessage(fake)).To(ContainSubstring("Expected to receive 'TakesThreeParameters' (and it did!) but the 3 argument failed to match"))
				Expect(subject.FailureMessage(fake)).To(ContainSubstring("<string>: whoops"))
				Expect(subject.FailureMessage(fake)).To(ContainSubstring("to equal"))
				Expect(subject.FailureMessage(fake)).To(ContainSubstring("<string>: c"))
			})

			It("should tell the user when it wasn't invoked at all", func() {
				subject.Match(fake)

				Expect(subject.FailureMessage(fake)).To(
					ContainSubstring("Expected to have received 'TakesThreeParameters', but it was not invoked"),
				)
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

			It("should tell the user when no invocations were expected, but some were invoked", func() {
				fake.Something()
				fake.Something()
				fake.Something()
				subject.Match(fake)

				Expect(subject.NegatedFailureMessage(fake)).To(ContainSubstring("to have received nothing, but there were 3 invocations"))
			})
		})

		Context("when the user specifies parameters to match", func() {
			BeforeEach(func() {
				subject = HaveReceived("TakesThreeParameters").
					With(Equal("a")).
					AndWith(Equal("b")).
					AndWith(Equal("c"))
			})

			It("should tell the user that the invocation occurred, despite their stating it should not occur", func() {
				fake.TakesThreeParameters("a", "b", "c")
				subject.Match(fake)
				Expect(subject.NegatedFailureMessage(fake)).To(ContainSubstring("to not receive 'TakesThreeParameters' (with exact argument matching)"))
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

	Context("when the user accidentally combines the no-argument HaveReceived() with argument matching", func() {
		var matcher types.GomegaMatcher

		BeforeEach(func() {
			matcher = HaveReceived().With(Equal("dang")).AndWith(Equal("You done goofed"))
		})

		It("should fail to match and tell the user they goofed", func() {
			matches, err := matcher.Match(fake)
			Expect(matches).To(BeFalse())
			Expect(err.Error()).To(ContainSubstring("You cannot combine HaveReceived() with argument matching"))
		})
	})
})
