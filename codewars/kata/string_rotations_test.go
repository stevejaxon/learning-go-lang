package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

func doRottest(strng string, arr []string, exp bool) {
    ans := ContainAllRots(strng, arr)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests ContainAllRots", func() {

    It("should handle basic cases", func() {
        doRottest("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}, true)
        doRottest("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}, false)
        doRottest("", []string{}, true)
    })
    
})
