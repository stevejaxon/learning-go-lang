package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Test Example", func() {
  It("returns the correct result on some examples", func() {
    Expect(HasUniqueChar("  nAa")).To(Equal(false))
    Expect(HasUniqueChar("abcde")).To(Equal(true))
    Expect(HasUniqueChar("++-")).To(Equal(false))
    Expect(HasUniqueChar("AaBbC")).To(Equal(true))
  })
})
