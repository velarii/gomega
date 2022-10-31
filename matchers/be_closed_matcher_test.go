package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"
	. "github.com/velarii/gomega/matchers"
)

var _ = Describe("BeClosedMatcher", func() {
	When("passed a channel", func() {
		It("should do the right thing", func() {
			openChannel := make(chan bool)
			Expect(openChannel).ShouldNot(BeClosed())

			var openReaderChannel <-chan bool
			openReaderChannel = openChannel
			Expect(openReaderChannel).ShouldNot(BeClosed())

			closedChannel := make(chan bool)
			close(closedChannel)

			Expect(closedChannel).Should(BeClosed())

			var closedReaderChannel <-chan bool
			closedReaderChannel = closedChannel
			Expect(closedReaderChannel).Should(BeClosed())
		})
	})

	When("passed a send-only channel", func() {
		It("should error", func() {
			openChannel := make(chan bool)
			var openWriterChannel chan<- bool
			openWriterChannel = openChannel

			success, err := (&BeClosedMatcher{}).Match(openWriterChannel)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())

			closedChannel := make(chan bool)
			close(closedChannel)

			var closedWriterChannel chan<- bool
			closedWriterChannel = closedChannel

			success, err = (&BeClosedMatcher{}).Match(closedWriterChannel)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())

		})
	})

	When("passed something else", func() {
		It("should error", func() {
			var nilChannel chan bool

			success, err := (&BeClosedMatcher{}).Match(nilChannel)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())

			success, err = (&BeClosedMatcher{}).Match(nil)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())

			success, err = (&BeClosedMatcher{}).Match(7)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
		})
	})
})
