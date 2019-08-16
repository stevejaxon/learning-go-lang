package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

func doCarttest(x,y int, exp [][]int){
  var act = CartesianNeighbor(x,y)
  Expect(exp).To(Equal(act))
}

var _ = Describe("Tests", func() {     
   It("ExampleTest", func() {
	doCarttest(2,2, [][]int{[]int{1, 1}, []int{1, 2}, []int{1, 3}, []int{2, 1}, []int{2, 3}, []int{3, 1}, []int{3, 2}, []int{3, 3}})
	doCarttest(5,6, [][]int{[]int{4, 5}, []int{4, 6}, []int{4, 7}, []int{5, 5}, []int{5, 7}, []int{6, 5}, []int{6, 6}, []int{6, 7}})
   })
})