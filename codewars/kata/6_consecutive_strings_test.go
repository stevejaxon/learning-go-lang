package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)

func doConsecStringstest(strarr []string, k int, exp string) {
    var ans = LongestConsec(strarr, k)
    Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

    It("should handle basic cases", func() {
        doConsecStringstest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta")
        doConsecStringstest([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1, 
            "oocccffuucccjjjkkkjyyyeehh")
            doConsecStringstest([]string{}, 3, "")
        doConsecStringstest([]string{"itvayloxrp","wkppqsztdkmvcuwvereiupccauycnjutlv","vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2, 
            "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
    })
})