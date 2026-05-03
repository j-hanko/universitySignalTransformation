package utils

func ASCII_to_bit(ascii string) (bit []int) {
	for _, char := range ascii {
		code := int(char)

		if code < 32 || code > 127 {
			panic("Dozwolone są tylko znaki ASCII od 32 do 127")
		}
		for i := 6; i >= 0; i-- {
			bit = append(bit, (code>>i)&1)
		}
	}
	return bit
}
