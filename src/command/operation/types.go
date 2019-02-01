package operation

import (
	"sync"

	"github.com/0nedark/serv/src/load"
)

// Handler defines the signature of operation
type Handler = func(load.Service, *sync.WaitGroup)
