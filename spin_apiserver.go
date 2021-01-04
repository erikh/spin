package spin

import (
	"log"

	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
)

// spin-apiserver service example implementation.
// The example methods log the requests and return zero values.
type spinApiserversrvc struct {
	logger *log.Logger
}

// NewSpinApiserver returns the spin-apiserver service implementation.
func NewSpinApiserver(logger *log.Logger) spinapiserver.Service {
	return &spinApiserversrvc{logger}
}
