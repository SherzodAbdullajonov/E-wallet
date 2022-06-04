package models

import "time"

type TopUps struct{
	Id int32 `json:"id"`
	Amount float32 `json:"amount"`
	AccountId int32 `json:"account_id"`
	Created_at time.Time `json:"created_at"`
}