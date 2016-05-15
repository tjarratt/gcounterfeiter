package gcounterfeiter

import (
	"fmt"

	"github.com/tjarratt/gcounterfeiter/invocations"
)

type haveReceivedNothingMatcher struct {
	functionWasInvokedCount int
}

func (m *haveReceivedNothingMatcher) Match(expected interface{}) (bool, error) {
	fake, ok := expected.(invocations.Recorder)
	if !ok {
		return false, expectedDoesNotImplementInterfaceError(expected)
	}

	m.functionWasInvokedCount = len(fake.Invocations())
	return m.functionWasInvokedCount != 0, nil
}

func (m *haveReceivedNothingMatcher) FailureMessage(expected interface{}) string {
	return fmt.Sprintf("Expected to have received nothing, but it received %d invocations", m.functionWasInvokedCount)
}

func (m *haveReceivedNothingMatcher) NegatedFailureMessage(expected interface{}) string {
	return "Expected to have received at least one invocation, but there were none"
}
