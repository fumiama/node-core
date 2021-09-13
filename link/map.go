package link

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	connmap = make(map[uint64]*websocket.Conn)
	connmu  sync.RWMutex
)

func saveMap(wsip uint64, conn *websocket.Conn) {
	if wsip > 0 {
		connmu.Lock()
		connmap[wsip] = conn
		connmu.Unlock()
	}
}

func delMap(wsip uint64) {
	_, ok := connmap[wsip]
	if ok {
		connmu.Lock()
		delete(connmap, wsip)
		connmu.Unlock()
	}
}
