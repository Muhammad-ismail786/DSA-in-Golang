
func reverseBits(n int) int {
	var result uint32

	for i := 0; i < 32; i++ {
		result = (result << 1) | (uint32(n) & 1)
		n >>= 1
	}

	return int(result)
}
