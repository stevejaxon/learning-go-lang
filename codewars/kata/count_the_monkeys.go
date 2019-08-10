package kata

func MonkeyCount(n int) []int {
	output := make([]int, n)
	for i := 1; i <= n; i++ {
		output[i-1] = i
	}
	return output
}
