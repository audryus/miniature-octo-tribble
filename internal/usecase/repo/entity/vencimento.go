package entity

import "time"

type Vencimento struct {
	Vencimento time.Time `json:"vencimento" bson:"_id,omitempty"`
}
