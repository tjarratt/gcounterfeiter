package gcounterfeiter

import (
	"fmt"

	"github.com/tjarratt/gcounterfeiter/invocations"
)

type haveReceivedNothingMatcher struct {
	expected                invocations.Recorder
	functionWasInvokedCount int
}

func (m *haveReceivedNothingMatcher) Match(expected interface{}) (bool, error) {
	fake, ok := expected.(invocations.Recorder)
	if !ok {
		return false, expectedDoesNotImplementInterfaceError(expected)
	}

	m.expected = fake
	return len(fake.Invocations()) != 0, nil
}

func (m *haveReceivedNothingMatcher) FailureMessage(interface{}) string {
	return fmt.Sprintf("Expected to have received nothing, but it received %d invocations", invocations.CountTotalInvocations(m.expected.Invocations()))
}

func (m *haveReceivedNothingMatcher) NegatedFailureMessage(interface{}) string {
	return "Expected to have received at least one invocation, but there were none"
}
