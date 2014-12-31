package slip

var (
	end    byte = 0xC0
	esc    byte = 0xDB
	escEnd byte = 0xDC
	escEsc byte = 0xDD
)

// Decode Decodes a slip packet
// Note that it may not be the full packet if end is not hit
// Returns decoded byte slice and error (currently does not error check)
func Decode(data []byte) (out []byte, err error) {
	for i := 0; i < len(data); i++ {
		// If it's an escaped character
		if data[i] == esc && i+1 < len(data) {
			i++
			if data[i] == escEnd {
				out = append(out, end)
			} else if data[i] == escEsc {
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

// Encode Encodes data into slip format
// Returns encoded byte slice and err (currently does not error check)
func Encode(data []byte) (out []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == end {
			out = append(out, esc)
			out = append(out, escEnd)
		} else if data[i] == esc {
			out = append(out, esc)
			out = append(out, escEsc)
		} else {
			out = append(out, data[i])
		}
	}
	out = append(out, end)
	return out, err
}
