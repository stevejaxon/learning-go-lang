package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)
var _ = Describe("CountBits()", func() {
  It("basic tests", func() {
    Expect(CountBits(0)).To(Equal(0))
    Expect(CountBits(4)).To(Equal(1))
    Expect(CountBits(7)         ).To(Equal(3))
    Expect(CountBits(9)         ).To(Equal(2))
    Expect(CountBits(10)        ).To(Equal(2))
  })
})