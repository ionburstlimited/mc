package mc

//

import (
	"hash"
	"hash/fnv"
)

// Hasher - interface for hasher implemetationa
type Hasher interface {
	Update(servers []*Server)
	GetServerIndex(key string) (uint, error)
}

type moduloHasher struct {
	nServers uint
	h32      hash.Hash32
}

// NewModuloHasher - create instance of modulo hasher
func NewModuloHasher() Hasher {
	var h Hasher = &moduloHasher{h32: fnv.New32a()}
	return h
}

func (h *moduloHasher) Update(servers []*Server) {
	h.nServers = uint(len(servers))
}

func (h *moduloHasher) GetServerIndex(key string) (uint, error) {
	if h.nServers < 1 {
		return 0, &Error{StatusNetworkError, "No server available", nil}
	}

	h.h32.Write([]byte(key))
	defer h.h32.Reset()

	return uint(h.h32.Sum32()) % h.nServers, nil
}
