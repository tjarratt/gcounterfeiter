package gcounterfeiter

import (
	"fmt"
	"github.com/onsi/gomega/types"
	"github.com/tjarratt/gcounterfeiter/invocations"
)

type haveReceivedMatcher struct {
	functionToMatch string
	expected        invocations.Recorder
}

func (m *haveReceivedMatcher) Match(expected interface{}) (bool, error) {
	fake, ok := expected.(invocations.Recorder)
	if !ok {
		return false, expectedDoesNotImplementInterfaceError(expected)
	}

	m.expected = fake
	return len(fake.Invocations()[m.functionToMatch]) > 0, nil
}

func (m *haveReceivedMatcher) FailureMessage(interface{}) string {
	return fmt.Sprintf("Expected to have received '%s', but it was not invoked", m.functionToMatch)
}

func (m *haveReceivedMatcher) NegatedFailureMessage(interface{}) string {
	invocationCount := invocations.CountTotalInvocations(m.expected.Invocations())
	return fmt.Sprintf("Expected to not have received '%s', but it was invoked %d times", m.functionToMatch, invocationCount)
}

func (m *haveReceivedMatcher) With(argumentMatcher types.GomegaMatcher) HaveReceivableMatcher {
	return NewArgumentVerifyingMatcher(m, m.functionToMatch, argumentMatcher)
}

func (m *haveReceivedMatcher) AndWith(argumentMatcher types.GomegaMatcher) HaveReceivableMatcher {
	return NewArgumentVerifyingMatcher(m, m.functionToMatch, argumentMatcher)
}
