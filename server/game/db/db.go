package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "chess-admin:root@tcp(mysql_chess)/chess")
	DBError(err)
	return db
}
func SetupDatabase() {
	fmt.Println("Setting up database connection")
	db, err := sql.Open("mysql", "chess-admin:root@tcp(mysql_chess)/chess")
	DBError(err)
	DBError(err)
	_, err = db.Exec(`DROP TABLE IF EXISTS rooms`)
	_, err = db.Exec(`DROP TABLE IF EXISTS players`)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS rooms (
		roomId VARCHAR(255) NOT NULL PRIMARY KEY,
		playersCount INT NOT NULL
	)`)
	DBError(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS players (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		position TEXT NOT NULL,
		roomId VARCHAR(255) NOT NULL  REFERENCES rooms(roomId)
		)`)
	DBError(err)
	defer db.Close()

}

func DBError(err error) {
	if err != nil {
		fmt.Errorf("!!DB Error", err.Error())
	}
}
