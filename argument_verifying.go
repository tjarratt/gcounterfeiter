package gcounterfeiter

import (
	"fmt"

	"github.com/onsi/gomega/types"
	"github.com/tjarratt/gcounterfeiter/invocations"
)

type argumentVerifyingMatcher struct {
	functionToMatch string
	baseMatcher     types.GomegaMatcher
	argMatchers     []types.GomegaMatcher

	expected invocations.Recorder

	failedArgIndex       int
	failedMatcherMessage string
}

func NewArgumentVerifyingMatcher(baseMatcher types.GomegaMatcher, functionToMatch string, argMatcher types.GomegaMatcher) *argumentVerifyingMatcher {
	return &argumentVerifyingMatcher{
		baseMatcher:     baseMatcher,
		functionToMatch: functionToMatch,
		argMatchers:     []types.GomegaMatcher{argMatcher},
	}
}

func (m *argumentVerifyingMatcher) Match(expected interface{}) (bool, error) {
	// FIXME :: this should probably combine matchers with AND
	ok, err := m.baseMatcher.Match(expected)
	if !ok || err != nil {
		return ok, err
	}

	fake, ok := expected.(invocations.Recorder)
	if !ok {
		return false, expectedDoesNotImplementInterfaceError(expected)
	}

	m.expected = fake

	invocations := fake.Invocations()[m.functionToMatch]
	for _, invocation := range invocations {
		if len(invocation) > len(m.argMatchers) {
			return false, fmt.Errorf("Too few arguments provided for '%s'. Expected %d but received %d", m.functionToMatch, len(invocation), len(m.argMatchers))
		}
		if len(invocation) < len(m.argMatchers) {
			return false, fmt.Errorf("Too many arguments provided for '%s'. Expected %d but received %d", m.functionToMatch, len(invocation), len(m.argMatchers))
		}

		for i, arg := range invocation {
			matcher := m.argMatchers[i]

			ok, err := matcher.Match(arg)
			if err != nil || !ok {
				m.failedArgIndex = i + 1
				m.failedMatcherMessage = matcher.FailureMessage(expected)
				return false, err
			}
		}
	}

	return true, nil
}

func (m *argumentVerifyingMatcher) FailureMessage(interface{}) string {
	return fmt.Sprintf(`Expected to receive '%s' (and it did!) but the %d argument failed to match:\n\t'%s'`, m.functionToMatch, m.failedArgIndex, m.failedMatcherMessage)
}

func (m *argumentVerifyingMatcher) NegatedFailureMessage(interface{}) string {
	return fmt.Sprintf("Expected to not receive '%s' (with exact argument matching)", m.functionToMatch)
}

func (m *argumentVerifyingMatcher) With(matcherOrValue interface{}) HaveReceivableMatcher {
	argumentMatcher := matcherOrWrapValueWithEqual(matcherOrValue)
	m.argMatchers = append(m.argMatchers, argumentMatcher)
	return m
}

func (m *argumentVerifyingMatcher) AndWith(matcherOrValue interface{}) HaveReceivableMatcher {
	argumentMatcher := matcherOrWrapValueWithEqual(matcherOrValue)
	m.argMatchers = append(m.argMatchers, argumentMatcher)
	return m
}
