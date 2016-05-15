package gcounterfeiter

import (
	"errors"

	"github.com/onsi/gomega/types"
)

type userDoneGoofedMatcher struct {
	message string
}

func newUserDoneGoofedMatcher(message string) *userDoneGoofedMatcher {
	return &userDoneGoofedMatcher{message: message}
}

func (m *userDoneGoofedMatcher) Match(interface{}) (bool, error) {
	return false, errors.New(m.message)
}

func (m *userDoneGoofedMatcher) FailureMessage(interface{}) string {
	return ""
}

func (m *userDoneGoofedMatcher) NegatedFailureMessage(interface{}) string {
	return ""
}

func (m *userDoneGoofedMatcher) With(argumentMatcher types.GomegaMatcher) HaveReceivableMatcher {
	return m
}

func (m *userDoneGoofedMatcher) AndWith(argumentMatcher types.GomegaMatcher) HaveReceivableMatcher {
	return m
}
