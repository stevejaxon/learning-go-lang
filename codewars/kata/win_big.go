package kata

import "math/big"

func Play(reserves, gambled string) [2]string {
	amountInMachine, _ := new(big.Int).SetString(reserves, 10)
	amountGambled, _ := new(big.Int).SetString(gambled, 10)
	// Check if there is enough money in the machine to pay out (in case of a win). If not then just return the amount gambled.
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