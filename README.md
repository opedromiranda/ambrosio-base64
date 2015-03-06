# Ambrosio Base64

Base64 encoding/decoding plugin for Ambrosio.

[![Build Status](https://travis-ci.org/opedromiranda/ambrosio-base64.svg)](https://travis-ci.org/opedromiranda/ambrosio-base64)


### Version
0.1

### Usage

Add the behaviour to Ambrosio:

```golang
import (
    "github.com/opedromiranda/ambrosio"
    "github.com/opedromiranda/ambrosio-base64"
)

func main() {
	steve := ambrosio.NewAmbrosio("Steve")

    steve.NewBehaviour(ambrosio_base64.Encoder)

	steve.Listen(3000)
}

```

License
----

MIT
