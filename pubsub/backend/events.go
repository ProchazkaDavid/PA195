package main

import "encoding/json"

// Message struct - Message should be returned for the `send_msg` event
type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Date   string `json:"date"`
	Room   string `json:"room"`
}

// Room struct - Room should be returned for the `create_room` event
type Room struct {
	Room string `json:"room"`
}

// FetchRoom struct - []FetchRoom should be returned for the `fetch_all` event
type FetchRoom struct {
	Room string `json:"room"`
	Msgs []Msg  `json:"msgs"`
}

// Msg is modified Message - no Room field :(
type Msg struct {
	Text   string `json:"text"`
	Sender string `json:"sender"`
	Date   string `json:"date"`
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

// MarshalBinary -
func (r *Room) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

// UnmarshalBinary -
func (r *Room) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	return nil
}
