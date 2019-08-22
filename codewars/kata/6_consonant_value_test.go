package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

var _ = Describe("Test Example", func() {
   It("should test that the solution returns the correct value", func() {
	 Expect(Solve("z")).To(Equal(26))
	 Expect(Solve("a")).To(Equal(0))
	 Expect(Solve("az")).To(Equal(26))
	 Expect(Solve("za")).To(Equal(26))
	 Expect(Solve("aza")).To(Equal(26))
     Expect(Solve("baz")).To(Equal(26))
     Expect(Solve("aeiou")).To(Equal(0))
     Expect(Solve("uaoczei")).To(Equal(29))
     Expect(Solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")).To(Equal(143))
     Expect(Solve("codewars")).To(Equal(37))
     Expect(Solve("bup")).To(Equal(16))
   })
})