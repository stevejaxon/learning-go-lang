package kata_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
)
var _ = Describe("Basic Tests", func() {
   It("Testing for Points(['1:0','2:0','3:0','4:0','2:1','3:1','4:1','3:2','4:2','4:3'])", func() {
     Expect(Points([]string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"})).To(Equal(30))
   })   
   It("Testing for points(['1:1','2:2','3:3','4:4','2:2','3:3','4:4','3:3','4:4','4:4'])", func() {
     Expect(Points([]string{"1:1","2:2","3:3","4:4","2:2","3:3","4:4","3:3","4:4","4:4"})).To(Equal(10))
   })
   It("Testing for points(['0:1','0:2','0:3','0:4','1:2','1:3','1:4','2:3','2:4','3:4'])", func() {
     Expect(Points([]string{"0:1","0:2","0:3","0:4","1:2","1:3","1:4","2:3","2:4","3:4"})).To(Equal(0))
   })
   It("Testing for points(['1:0','2:0','3:0','4:0','2:1','1:3','1:4','2:3','2:4','3:4'])", func() {
     Expect(Points([]string{"1:0","2:0","3:0","4:0","2:1","1:3","1:4","2:3","2:4","3:4"})).To(Equal(15))
   })
   It("Testing for points(['1:0','2:0','3:0','4:4','2:2','3:3','1:4','2:3','2:4','3:4'])", func() {
     Expect(Points([]string{"1:0","2:0","3:0","4:4","2:2","3:3","1:4","2:3","2:4","3:4"})).To(Equal(12))
   })
})