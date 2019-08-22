package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)
var _ = Describe("Test Example", func() {
  It("should test that the solution returns the correct value", func() {
	 Expect(Solution("a")).To(Equal([]string{"a_"}))
	 Expect(Solution("ab")).To(Equal([]string{"ab"}))
	 Expect(Solution("abcd")).To(Equal([]string{"ab", "cd"}))
     Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
   })
})