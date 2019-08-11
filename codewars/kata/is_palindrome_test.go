package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Test Example", func() {
  It("tests basic strings", func() {
	Expect(IsPalindrome("")).To(Equal(false))
	Expect(IsPalindrome("  ")).To(Equal(false))
	Expect(IsPalindrome("a")).To(Equal(true))
	Expect(IsPalindrome("aa")).To(Equal(true))
    Expect(IsPalindrome("aba")).To(Equal(true))
    Expect(IsPalindrome("Abba")).To(Equal(true))
	Expect(IsPalindrome("he")).To(Equal(false))
	Expect(IsPalindrome("hello")).To(Equal(false))
	Expect(IsPalindrome("racecar")).To(Equal(true))
	Expect(IsPalindrome("Bob")).To(Equal(true))
    Expect(IsPalindrome("Madam")).To(Equal(true))
    Expect(IsPalindrome("AbBa")).To(Equal(true))
  })
})
