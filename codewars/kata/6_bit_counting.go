package kata

func CountBits(input uint) (bits int) {
	for input > 0 {
		if input & 1 != 0 {
			bits++
		}
		input = input >> 1
	}
	return
}
