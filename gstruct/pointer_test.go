package gstruct_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"
	. "github.com/rshbintech/gomega/gstruct"
)

var _ = Describe("PointTo", func() {
	It("should fail when passed nil", func() {
		var p *struct{}
		Expect(p).Should(BeNil())
	})

	It("should succeed when passed non-nil pointer", func() {
		var s struct{}
		Expect(&s).Should(PointTo(Ignore()))
	})

	It("should unwrap the pointee value", func() {
		i := 1
		Expect(&i).Should(PointTo(Equal(1)))
		Expect(&i).ShouldNot(PointTo(Equal(2)))
	})

	It("should work with nested pointers", func() {
		i := 1
		ip := &i
		ipp := &ip
		Expect(ipp).Should(PointTo(PointTo(Equal(1))))
		Expect(ipp).ShouldNot(PointTo(PointTo(Equal(2))))
	})
})
