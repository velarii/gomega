package matchers_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"
	. "github.com/rshbintech/gomega/matchers"
)

type CustomError struct {
}

func (c CustomError) Error() string {
	return "an error"
}

var _ = Describe("MatchErrorMatcher", func() {
	Context("When asserting against an error", func() {
		When("passed an error", func() {
			It("should succeed when errors are deeply equal", func() {
				err := errors.New("an error")
				fmtErr := fmt.Errorf("an error")
				customErr := CustomError{}

				Expect(err).Should(MatchError(errors.New("an error")))
				Expect(err).ShouldNot(MatchError(errors.New("another error")))

				Expect(fmtErr).Should(MatchError(errors.New("an error")))
				Expect(customErr).Should(MatchError(CustomError{}))
			})

			It("should succeed when any error in the chain matches the passed error", func() {
				innerErr := errors.New("inner error")
				outerErr := fmt.Errorf("outer error wrapping: %w", innerErr)

				Expect(outerErr).Should(MatchError(innerErr))
			})
		})

		When("actual an expected are both pointers to an error", func() {
			It("should succeed when errors are deeply equal", func() {
				err := CustomError{}
				Expect(&err).To(MatchError(&err))
			})
		})

		It("should succeed when matching with a string", func() {
			err := errors.New("an error")
			fmtErr := fmt.Errorf("an error")
			customErr := CustomError{}

			Expect(err).Should(MatchError("an error"))
			Expect(err).ShouldNot(MatchError("another error"))

			Expect(fmtErr).Should(MatchError("an error"))
			Expect(customErr).Should(MatchError("an error"))
		})

		When("passed a matcher", func() {
			It("should pass if the matcher passes against the error string", func() {
				err := errors.New("error 123 abc")

				Expect(err).Should(MatchError(MatchRegexp(`\d{3}`)))
			})

			It("should fail if the matcher fails against the error string", func() {
				err := errors.New("no digits")
				Expect(err).ShouldNot(MatchError(MatchRegexp(`\d`)))
			})
		})

		It("should fail when passed anything else", func() {
			actualErr := errors.New("an error")
			_, err := (&MatchErrorMatcher{
				Expected: []byte("an error"),
			}).Match(actualErr)
			Expect(err).Should(HaveOccurred())

			_, err = (&MatchErrorMatcher{
				Expected: 3,
			}).Match(actualErr)
			Expect(err).Should(HaveOccurred())
		})
	})

	When("passed nil", func() {
		It("should fail", func() {
			_, err := (&MatchErrorMatcher{
				Expected: "an error",
			}).Match(nil)
			Expect(err).Should(HaveOccurred())
		})
	})

	When("passed a non-error", func() {
		It("should fail", func() {
			_, err := (&MatchErrorMatcher{
				Expected: "an error",
			}).Match("an error")
			Expect(err).Should(HaveOccurred())

			_, err = (&MatchErrorMatcher{
				Expected: "an error",
			}).Match(3)
			Expect(err).Should(HaveOccurred())
		})
	})

	When("passed an error that is also a string", func() {
		It("should use it as an error", func() {
			var e mockErr = "mockErr"

			// this fails if the matcher casts e to a string before comparison
			Expect(e).Should(MatchError(e))
		})
	})

	It("shows failure message", func() {
		failuresMessages := InterceptGomegaFailures(func() {
			Expect(errors.New("foo")).To(MatchError("bar"))
		})
		Expect(failuresMessages[0]).To(ContainSubstring("{s: \"foo\"}\nto match error\n    <string>: bar"))
	})

	It("shows negated failure message", func() {
		failuresMessages := InterceptGomegaFailures(func() {
			Expect(errors.New("foo")).ToNot(MatchError("foo"))
		})
		Expect(failuresMessages[0]).To(ContainSubstring("{s: \"foo\"}\nnot to match error\n    <string>: foo"))
	})
})

type mockErr string

func (m mockErr) Error() string { return string(m) }
