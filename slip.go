package slip

var (
	end     byte = 0xC0
	esc     byte = 0xDB
	esc_end byte = 0xDC
	esc_esc byte = 0xDD
)

func Decode(data []byte) (out []byte, err error) {
	for i := 0; i < len(data); i++ {
		// If it's an escaped character
		if data[i] == esc && i+1 < len(data) {
			i++
			if data[i] == esc_end {
				out = append(out, end)
			} else if data[i] == esc_esc {
				out = append(out, esc)
			}
		} else if data[i] == end {
			break
		} else {
			out = append(out, data[i])
		}
	}
	return out, err
}

func Encode(data []byte) (out []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == end {
			out = append(out, esc)
			out = append(out, esc_end)
		} else if data[i] == esc {
			out = append(out, esc)
			out = append(out, esc_esc)
		} else {
			out = append(out, data[i])
		}
	}
	out = append(out, end)
	return out, err
}
