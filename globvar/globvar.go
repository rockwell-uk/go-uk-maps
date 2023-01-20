package globvar

import (
	"os"
)

var (
	Stop = make(chan os.Signal, 1)
	Http = make(chan struct{}, 1)
)
