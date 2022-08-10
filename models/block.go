package models

import "time"

type Block struct {
	Index     int
	Timestamp time.Time
	BPM       int
	Hash      string
	PrevHash  string
}

type Message struct {
	BPM int `json:"bpm"`
}
