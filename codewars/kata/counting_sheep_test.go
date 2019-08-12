package kata_test

import (
  "strings"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Sample Test Cases", func() {
	It("The solution should returns the correct values", func() {
		Expect(CountSheep(-1)).To(Equal(""))
		Expect(CountSheep(0)).To(Equal(""))
		Expect(CountSheep(1)).To(Equal("1 sheep..."))
		Expect(CountSheep(2)).To(Equal("1 sheep...2 sheep..."))
		Expect(CountSheep(3)).To(Equal("1 sheep...2 sheep...3 sheep..."))
		Expect(strings.Contains(CountSheep(10), "10 sheep...")).To(Equal(true))
	})
})