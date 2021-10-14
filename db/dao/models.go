package db

import "time"

type KeyValue struct {
	Key      string    `json:"key"`
	Value    string    `json:"value"`
	UpdateAt time.Time `json:"updated_at"`
}
