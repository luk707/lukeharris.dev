package r30 /* import "go.lukeharris.dev/r30" */

func shiftRight(data []byte) []byte {
	carry := data[len(data)-1] & 0b1 << 7
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = (data[i] >> 1) + carry
		carry = data[i] & 0b1 << 7
	}
	return result
}

func shiftLeft(data []byte) []byte {
	carry := data[0] & 0b10000000 >> 7
	result := make([]byte, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		result[i] = (data[i] << 1) + carry
		carry = data[i] & 0b10000000 >> 7
	}
	return result
}

// Computes the next iteration of the rule 30 cellular automata
func Step(data []byte) []byte {
	result := make([]byte, len(data))

	right := shiftLeft(data)
	left := shiftRight(data)

	for i := range data {
		result[i] = left[i] ^ (data[i] | right[i])
	}

	return result
}
