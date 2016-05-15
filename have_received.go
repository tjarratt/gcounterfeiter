package gcounterfeiter

import "fmt"

type haveReceivedMatcher struct {
	functionToMatch         string
	functionWasInvokedCount int
}

func (m *haveReceivedMatcher) Match(expected interface{}) (bool, error) {
	fake, ok := expected.(InvocationRecorder)
	if !ok {
		return false, expectedDoesNotImplementInterfaceError(expected)
	}

	m.functionWasInvokedCount = len(fake.Invocations()[m.functionToMatch])
	return m.functionWasInvokedCount > 0, nil
}

func (m *haveReceivedMatcher) FailureMessage(interface{}) string {
	return fmt.Sprintf("Expected to have received '%s', but it was not invoked", m.functionToMatch)
}

func (m *haveReceivedMatcher) NegatedFailureMessage(interface{}) string {
	return fmt.Sprintf("Expected to not have received '%s', but it was invoked %d times", m.functionToMatch, m.functionWasInvokedCount)
}
