package player

import (
	"chess/game/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Player struct {
	Name     string
	Position string
	Id       string
	RoomId   string
}

type Room struct {
	RoomId       string
	PlayersCount int
}

func AddPlayer(name string) (string, error) {
	var id = uuid.New().String()
	var roomId string
	var playersCount int = 1
	var pos string = "1"
	chessDB := db.Database()
	var room Room
	err := chessDB.QueryRow(`SELECT * FROM rooms WHERE playersCount = ? LIMIT 1`, 1).Scan(&room.RoomId, &room.PlayersCount)
	db.DBError(err)
	if room.RoomId != "" {
		roomId = room.RoomId
		playersCount = room.PlayersCount
		pos = "2"
	}
	if len(roomId) == 0 {
		roomId = uuid.NewString()
	}
	_, err = chessDB.Exec(`INSERT INTO rooms (roomId, playersCount) VALUES(?,?) ON DUPLICATE KEY UPDATE playersCount=playersCount+1`, roomId, playersCount)
	_, err = chessDB.Exec("INSERT INTO players (id, name, position, roomId) VALUES(?,?,?,?)", id, name, pos, roomId)
	db.DBError(err)
	defer chessDB.Close()
	return pos, nil
}
