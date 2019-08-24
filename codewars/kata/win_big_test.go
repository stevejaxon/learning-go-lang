package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/stevejaxon/learning-go-lang/codewars/kata"
  "math/big"
  "math/rand"
)

var _ = Describe("Example tests", func() {
	It("Example test cases", func() {
		Expect(Play("100000000", "1000000")).To(Equal([2]string{"101000000", "0"}))
		Expect(Play("99991808", "1000000")).To(Equal([2]string{"97991808", "3000000"}))
		Expect(Play("16777215", "4000000")).To(Equal([2]string{"20777215", "0"}))
		Expect(Play("100000000", "100000000")).To(Equal([2]string{"100000000", "100000000"}))
		Expect(Play("429461545327902976", "592817564534231")).To(Equal([2]string{"428275910198834514", "1778452693602693"}))
		Expect(Play("4611686018427387903", "937836257282654172")).To(Equal([2]string{"5549522275710042075", "0"}))
		Expect(Play("19223372036854775808", "578721384725774464")).To(Equal([2]string{"18065929267403226880", "1736164154177323392"}))
		Expect(Play("9444732965739290427392", "1")).To(Equal([2]string{"9444732965739290427393", "0"}))
	})

	It("Edge case tests", func() {
		Expect(Play("0", "0")).To(Equal([2]string{"0", "0"}))
		Expect(Play("1", "0")).To(Equal([2]string{"1", "0"}))
		Expect(Play("0", "1")).To(Equal([2]string{"0", "1"}))
		Expect(Play("1", "1")).To(Equal([2]string{"1", "1"}))
		Expect(Play("3", "1")).To(Equal([2]string{"4", "0"}))
		Expect(Play("2", "1")).To(Equal([2]string{"3", "0"}))
		Expect(Play("5", "2")).To(Equal([2]string{"1", "6"}))
		Expect(Play("100000000", "50000001")).To(Equal([2]string{"100000000", "50000001"}))
		Expect(Play("9223372036854775807", "1")).To(Equal([2]string{"9223372036854775808", "0"}))
		Expect(Play("9223372034707292159", "1")).To(Equal([2]string{"9223372034707292157", "3"}))
	})

	It("Valid random test cases", func() {
		rnd := rand.New(rand.NewSource(42))
		for i := 0; i < 10; i++ {
			jackpot := getVeryLargeRandom(rnd)
			jackpotString := jackpot.String()
			gamble := getValidGambleAmount(rnd, jackpot).String()
			Expect(Play(jackpotString, gamble)).To(Equal(solution(jackpotString, gamble)))
		}
	})
})

func getVeryLargeRandom(rnd *rand.Rand) *big.Int {
	max, _ := new(big.Int).SetString("100000000000000000000000000", 10)
	return new(big.Int).Rand(rnd, max)
}

func getValidGambleAmount(rnd *rand.Rand, jackpot *big.Int) *big.Int {
	max := new(big.Int).Div(jackpot, big.NewInt(int64(2)))
	return new(big.Int).Rand(rnd, max)
}

func solution(reserves, gambled string) [2]string {
	amountInMachine, _ := new(big.Int).SetString(reserves, 10)
	amountGambled, _ := new(big.Int).SetString(gambled, 10)
	jackpot := new(big.Int).Mul(amountGambled, big.NewInt(int64(2)))
	if jackpot.Cmp(amountInMachine) > 0 {
		return [2]string{reserves, gambled}
	}
	if hasWon(amountInMachine) {
		remaining := new(big.Int).Sub(amountInMachine, jackpot)
		payout := new(big.Int).Add(amountGambled, jackpot)
		return [2]string{remaining.String(), payout.String()}
	}
	remaining := new(big.Int).Add(amountInMachine, amountGambled)
	return [2]string{remaining.String(), "0"}
}

func hasWon(amountInMachine *big.Int) bool {
	numBits := amountInMachine.BitLen()
	if numBits & 1 == 1 {
		return amountInMachine.Bit(numBits/2) == uint(0)
	}
	return false
}
