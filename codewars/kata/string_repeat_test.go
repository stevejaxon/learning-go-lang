package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Example Tests", func() {
	It("should repeat correctly", func() {
		Expect(RepeatStr(4, "a")).To(Equal("aaaa"))
		Expect(RepeatStr(3, "hello ")).To(Equal("hello hello hello "))
		Expect(RepeatStr(2, "abc")).To(Equal("abcabc"))
	})
})
