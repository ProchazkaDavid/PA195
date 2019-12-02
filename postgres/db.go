package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func getDBConnection() *sql.DB {
	// Closing db instance left for caller. Call `defer db.Close()`
	host := os.Getenv("PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func insertMessage(db *sql.DB, message string, senderId int, channelId int) int {
	insertStatement := `
	INSERT INTO messages (sender, time_sent, channel, content)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(insertStatement, senderId, time.Now().UTC().String(), channelId, message).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
