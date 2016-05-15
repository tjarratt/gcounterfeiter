package gcounterfeiter

import "fmt"

type userDoneGoofedMatcher struct {
	howMany int
}

func (m *userDoneGoofedMatcher) Match(interface{}) (bool, error) {
	return false, fmt.Errorf("You provided too many arguments. Expected 0 or 1, but you provided %d", m.howMany)
}

func (m *userDoneGoofedMatcher) FailureMessage(interface{}) string {
	return ""
}

func (m *userDoneGoofedMatcher) NegatedFailureMessage(interface{}) string {
	return ""
}
