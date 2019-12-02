package main

import (
	"encoding/json"
)

// Message struct
type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Date   string `json:"date"`
	Room   string `json:"room"`
}

// MarshalBinary -
func (m *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary -
func (m *Message) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	return nil
}
