package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"
	. "github.com/velarii/gomega/matchers"
)

var _ = Describe("BeEmpty", func() {
	When("passed a supported type", func() {
		It("should do the right thing", func() {
			Expect("").Should(BeEmpty())
			Expect(" ").ShouldNot(BeEmpty())

			Expect([0]int{}).Should(BeEmpty())
			Expect([1]int{1}).ShouldNot(BeEmpty())

			Expect([]int{}).Should(BeEmpty())
			Expect([]int{1}).ShouldNot(BeEmpty())

			Expect(map[string]int{}).Should(BeEmpty())
			Expect(map[string]int{"a": 1}).ShouldNot(BeEmpty())

			c := make(chan bool, 1)
			Expect(c).Should(BeEmpty())
			c <- true
			Expect(c).ShouldNot(BeEmpty())
		})
	})

	When("passed a correctly typed nil", func() {
		It("should be true", func() {
			var nilSlice []int
			Expect(nilSlice).Should(BeEmpty())

			var nilMap map[int]string
			Expect(nilMap).Should(BeEmpty())
		})
	})

	When("passed an unsupported type", func() {
		It("should error", func() {
			success, err := (&BeEmptyMatcher{}).Match(0)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())

			success, err = (&BeEmptyMatcher{}).Match(nil)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
		})
	})
})
