package gstruct_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gstruct Suite")
}
