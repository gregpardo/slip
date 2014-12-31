package slip

import (
	"bytes"
	"testing"
)

var (
	encoded = []byte{0x01, esc, escEsc, 0x02, esc, escEnd, 0x3, end}
	decoded = []byte{0x01, esc, 0x02, end, 0x03}
)

func TestEncode(t *testing.T) {
	result, err := Encode(decoded)
	if bytes.Equal(result, encoded) == false || err != nil {
		t.Errorf("Encoding failed")
	}
}

func TestDecode(t *testing.T) {
	result, err := Decode(encoded)
	if bytes.Equal(result, decoded) == false || err != nil {
		t.Errorf("Decoding failed \nResult: %X", result)
	}
}
