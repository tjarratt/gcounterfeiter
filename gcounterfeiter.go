package gcounterfeiter

import (
	"github.com/onsi/gomega/types"
)

func HaveReceived(args ...string) types.GomegaMatcher {
	switch len(args) {
	case 0:
		return &haveReceivedNothingMatcher{}
	case 1:
		return &haveReceivedMatcher{functionToMatch: args[0]}
	default:
		return &userDoneGoofedMatcher{howMany: len(args)}
	}
}

type InvocationRecorder interface {
	Invocations() map[string][][]interface{}
}
