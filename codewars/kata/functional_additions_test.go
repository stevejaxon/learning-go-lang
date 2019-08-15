package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)
var _ = It("sample test", func() {
  Expect(Add(0)(1)).To(Equal(1))
  Expect(Add(1)(3)).To(Equal(4))
})
