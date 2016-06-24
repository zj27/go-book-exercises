package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount23(x uint64) int {
	var result byte = 0
	for i := uint(0); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

func PopCount24(x uint64) int {
	count := 0
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 == 1 {
			count++
		}
	}
	return count
}

func PopCount25(x uint64) int {
	count := 0
	for i := x; i > 0; {
		count++
		i = i & (i - 1)
	}
	return count
}
