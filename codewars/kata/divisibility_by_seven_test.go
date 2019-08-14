package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"  
)

func testSeven(n int64, exp []int) {
    var ans = Seven(n)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Seven", func() {

    It("should handle basic cases", func() {
		testSeven(0, []int{0,0})
		testSeven(91, []int{7, 1})
		testSeven(371, []int{35, 1})
        testSeven(477557101, []int{28, 7})
        testSeven(1889584453156367, []int{0, 0})
        testSeven(1603, []int{7, 2})
        testSeven(372, []int{0, 0})
        testSeven(2099061, []int{0, 0})
    })   
})