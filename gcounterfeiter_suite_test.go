package gcounterfeiter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGcounterfeiter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gcounterfeiter Suite")
}
