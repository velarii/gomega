package matchers_test

import (
	"time"

	. "github.com/rshbintech/gomega/matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"
)

var _ = Describe("BeSent", func() {
	When("passed a channel and a matching type", func() {
		When("the channel is ready to receive", func() {
			It("should succeed and send the value down the channel", func() {
				c := make(chan string)
				d := make(chan string)
				go func() {
					val := <-c
					d <- val
				}()

				time.Sleep(10 * time.Millisecond)

				Expect(c).Should(BeSent("foo"))
				Eventually(d).Should(Receive(Equal("foo")))
			})

			It("should succeed (with a buffered channel)", func() {
				c := make(chan string, 1)
				Expect(c).Should(BeSent("foo"))
				Expect(<-c).Should(Equal("foo"))
			})
		})

		When("the channel is not ready to receive", func() {
			It("should fail and not send down the channel", func() {
				c := make(chan string)
				Expect(c).ShouldNot(BeSent("foo"))
				Consistently(c).ShouldNot(Receive())
			})
		})

		When("the channel is eventually ready to receive", func() {
			It("should succeed", func() {
				c := make(chan string)
				d := make(chan string)
				go func() {
					time.Sleep(30 * time.Millisecond)
					val := <-c
					d <- val
				}()

				Eventually(c).Should(BeSent("foo"))
				Eventually(d).Should(Receive(Equal("foo")))
			})
		})

		When("the channel is closed", func() {
			It("should error", func() {
				c := make(chan string)
				close(c)
				success, err := (&BeSentMatcher{Arg: "foo"}).Match(c)
				Expect(success).Should(BeFalse())
				Expect(err).Should(HaveOccurred())
			})

			It("should short-circuit Eventually", func() {
				c := make(chan string)
				close(c)

				t := time.Now()
				failures := InterceptGomegaFailures(func() {
					Eventually(c, 10.0).Should(BeSent("foo"))
				})
				Expect(failures).Should(HaveLen(1))
				Expect(time.Since(t)).Should(BeNumerically("<", time.Second))
			})
		})
	})

	When("passed a channel and a non-matching type", func() {
		It("should error", func() {
			success, err := (&BeSentMatcher{Arg: "foo"}).Match(make(chan int, 1))
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
		})
	})

	When("passed a receive-only channel", func() {
		It("should error", func() {
			var c <-chan string
			c = make(chan string, 1)
			success, err := (&BeSentMatcher{Arg: "foo"}).Match(c)
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
		})
	})

	When("passed a nonchannel", func() {
		It("should error", func() {
			success, err := (&BeSentMatcher{Arg: "foo"}).Match("bar")
			Expect(success).Should(BeFalse())
			Expect(err).Should(HaveOccurred())
		})
	})
})
