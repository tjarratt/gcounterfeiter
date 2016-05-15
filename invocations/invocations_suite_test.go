package invocations_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInvocations(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Invocations Suite")
}
