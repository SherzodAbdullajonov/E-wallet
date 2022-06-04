package models

import "time"

type Accounts struct {
	Id         int32     `json:"id"`
	User_id    int32     `json:"user_id"`
	Balance    float32   `json:"balance"`
	Identified bool      `json:"identified"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
