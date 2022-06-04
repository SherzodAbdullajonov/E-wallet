package models

import "time"

type Users struct {
	Id         int32     `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
