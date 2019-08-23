package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Example tests", func() {
	It("Edge case tests", func() {
		Expect(Play("0", "0")).To(Equal([2]string{"0", "0"}))
		Expect(Play("1", "0")).To(Equal([2]string{"1", "0"}))
		Expect(Play("0", "1")).To(Equal([2]string{"0", "1"}))
		Expect(Play("1", "1")).To(Equal([2]string{"1", "1"}))
		//Expect(Play("3", "1")).To(Equal([2]string{"1", "3"}))
		//Expect(Play("2", "1")).To(Equal([2]string{"3", "0"}))
		Expect(Play("9223372036854775807", "1")).To(Equal([2]string{"9223372036854775808", "0"}))
		Expect(Play("9223372034707292159", "1")).To(Equal([2]string{"9223372034707292157", "3"}))
	})
})