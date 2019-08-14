package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"

)

func dotest(card, ticket int, perc float64, exp int) {
    var ans = Movie(card, ticket, perc)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Movie", func() {

    It("should handle basic cases", func() {
		dotest(500, 15, 0.9, 43)
		dotest(100, 10, 0.95, 24)
		dotest(0, 10, 0.95, 2)
		dotest(620, 20, 0.49, 33)
		dotest(1310, 37, 0.72, 39)
    })
})
