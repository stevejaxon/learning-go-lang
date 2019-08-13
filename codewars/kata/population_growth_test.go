package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("NbYear", func() {
  It("fixed tests", func() {
    Expect(NbYear(1500, 5, 100, 5000)).To(Equal(15))
    Expect(NbYear(1500000, 2.5, 10000, 2000000)).To(Equal(10))
    Expect(NbYear(1500000, 0.25, 1000, 2000000)).To(Equal(94))
    Expect(NbYear(1500000, 0.25, -1000, 2000000)).To(Equal(151))
  })
})