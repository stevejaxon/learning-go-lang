package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
  "fmt"
)

func singleTest(str string, res bool) {
    It(fmt.Sprintf("should return %v for \"%v\"",res,str), func() {
      Expect(ValidBraces(str)).To(Equal(res))
    })
}

var _ = Describe("Valid Braces", func() { 
    singleTest("(){}[]",true)
    singleTest("([{}])",true)
    singleTest("(}",false)
    singleTest("[(])",false)
	singleTest("[({)](]",false)
	singleTest("{()",false)
	singleTest("]",false)
	singleTest("[]]",false)
})