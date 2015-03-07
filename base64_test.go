package ambrosio_base64

import (
	"testing"
	"regexp"
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

func TestPattern(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"encode 1", true},
		{"encode balili", true},
		{"decode this", true},
		{"decode that as well", true},
		{"decode \"#$%&/()\"", true},

		{"decodeABC", false},
		{"encode123", false},
	}

	for _, c := range tests {
		got, _ := regexp.MatchString(Encoder.Pattern, c.s)
		if got != c.want {
			t.Errorf("regexp.MatchString(%q, %q) == %t, want %t", Encoder.Pattern, c.s, got, c.want)
		}
	}
}
