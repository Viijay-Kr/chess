package room

import (
	player "chess/game/players"
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func CreateConnection() *socketio.Server {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("connection successful", so.ID())
		return nil
	})

	server.OnError("/", func(so socketio.Conn, err error) {
		log.Println("error:", err)
	})

	server.OnEvent("/room", "join-room", func(s socketio.Conn, name string) string {
		s.SetContext(name)
		fmt.Println("Trying to add player into database")
		pos, err := player.AddPlayer(name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Player joined successfully")
		return pos
	})
	return server
}
