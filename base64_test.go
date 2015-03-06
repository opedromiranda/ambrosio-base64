package ambrosio_base64

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var tests = []struct {
		s    []string
		want string
	}{
		{[]string{"encode inzayn", "encode", "inzayn"}, "aW56YXlu"},
	}

	for _, c := range tests {
		got, _ := Encoder.Handler(c.s)
		if got != c.want {
			t.Errorf("Encoder.Handler(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

func TestDecode(t *testing.T) {
    var tests = []struct {
		s    []string
		want string
	}{
		{[]string{"decode aW56YXlu", "decode", "aW56YXlu"}, "inzayn"},
	}

	for _, c := range tests {
		got, _ := Encoder.Handler(c.s)
		if got != c.want {
			t.Errorf("Encoder.Handler(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
