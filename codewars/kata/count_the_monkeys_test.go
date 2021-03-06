// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("MonkeyCount", func() {
	It("Should work for fixed tests", func() {
		Expect(MonkeyCount(1)).To(Equal([]int{1}))
		Expect(MonkeyCount(2)).To(Equal([]int{1, 2}))
		Expect(MonkeyCount(3)).To(Equal([]int{1, 2, 3}))
		Expect(MonkeyCount(4)).To(Equal([]int{1, 2, 3, 4}))
		Expect(MonkeyCount(5)).To(Equal([]int{1, 2, 3, 4, 5}))
		Expect(MonkeyCount(6)).To(Equal([]int{1, 2, 3, 4, 5, 6}))
		Expect(MonkeyCount(7)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7}))
		Expect(MonkeyCount(8)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8}))
		Expect(MonkeyCount(9)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
		Expect(MonkeyCount(10)).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	})
})
