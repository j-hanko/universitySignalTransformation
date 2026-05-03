package lab6

const (
	A1   = 0.5
	A2   = 1.0
	fs   = 2000
	step = 200
	fn   = 10
	Tc   = 1
)

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
