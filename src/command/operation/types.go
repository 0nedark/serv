package operation

import (
	"sync"

	"github.com/drupsys/serv/src/load"
)

// Handler defines the signature of operation
type Handler = func(load.Service, *sync.WaitGroup)
