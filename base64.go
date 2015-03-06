package ambrosio_base64

import (
	"encoding/base64"
	"github.com/opedromiranda/ambrosio"
)

var Encoder = ambrosio.Behaviour{
	"^(encode|decode) (.*)$",
	func(matches []string) (string, error) {
		var result string
		var error error

		operation := matches[1]
		input := matches[2]

		switch operation {
		case "encode":
			result = base64.StdEncoding.EncodeToString([]byte(input))
		case "decode":
			b, error := base64.StdEncoding.DecodeString(input)

			if error == nil {
				result = string(b)
			}
		}

		return result, error
	},
}
