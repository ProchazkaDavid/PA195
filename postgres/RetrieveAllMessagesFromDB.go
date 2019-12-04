package postgres

import (
	"log"

	_ "github.com/lib/pq"
)

type ChatMessage struct {
	Id, Sender, Time_sent, Channel, Content string
}

func retrieveAllMessages(limit int) []ChatMessage {
	db := getDBConnection()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	statement, err := db.Prepare("SELECT * FROM messages LIMIT $1")
	if err != nil {
		panic(err)
	}

	var data []ChatMessage

	rows, err := statement.Query(limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, sender, time_sent, channel, content string
		if err := rows.Scan(&id, &sender, &time_sent, &channel, &content); err != nil {
			panic(err)
		}
		row := &ChatMessage{Id: id, Sender: sender, Time_sent: time_sent, Channel: channel, Content: content}
		data = append(data, *row)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v", data)

	return data

}
