package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Basic tests", func() {
	It("should return the correct values", func() {
		Expect(Century(int(1990))).To(Equal(20))
		Expect(Century(int(1705))).To(Equal(18))
		Expect(Century(int(1900))).To(Equal(19))
		Expect(Century(int(1601))).To(Equal(17))
		Expect(Century(int(2000))).To(Equal(20))
		Expect(Century(int(89))).To(Equal(1))
		Expect(Century(int(1))).To(Equal(1))
		Expect(Century(int(100))).To(Equal(1))
		Expect(Century(int(101))).To(Equal(2))
		Expect(Century(int(10000))).To(Equal(100))
		Expect(Century(int(10001))).To(Equal(101))
	})
})