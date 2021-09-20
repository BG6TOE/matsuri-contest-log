package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	broadcastChans map[int]*websocket.Conn
	chanLock       sync.RWMutex
	ctr            int
)

func init() {
	broadcastChans = make(map[int]*websocket.Conn)
}

func insertConnection(c *websocket.Conn) int {
	chanLock.Lock()
	defer chanLock.Unlock()
	ctr++
	broadcastChans[ctr] = c
	return ctr
}

func removeConnection(id int) {
	chanLock.Lock()
	defer chanLock.Unlock()
	if c, ok := broadcastChans[id]; ok {
		c.Close()
		delete(broadcastChans, id)
	}
}

func HandleWebsocketConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	logrus.Infof("Incoming websocket connection: %s (from %s)", r.RemoteAddr, r.Header.Get("Origin"))

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Infof("Failed to establish websocket connection: %v", err)
		return
	}

	id := insertConnection(ws)

	defer removeConnection(id)

	// Ignore all messages
	ws.SetReadLimit(512)
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

type BroadcastMessage struct {
	Class   string
	Message interface{}
}

func Broadcast(msg *BroadcastMessage) {
	go func(msg *BroadcastMessage) {
		logrus.Debugf("Broadcasting %v message to all clients", msg.Class)
		chanLock.Lock()
		defer chanLock.Unlock()
		for _, v := range broadcastChans {
			err := v.WriteJSON(msg)
			if err != nil {
				logrus.Errorf("Failed to send broadcast to %s: %v", v.RemoteAddr().String(), err)
			}
		}
	}(msg)
}
