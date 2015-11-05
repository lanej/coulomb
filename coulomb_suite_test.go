package coulomb_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCoulomb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Coulomb Suite")
}
