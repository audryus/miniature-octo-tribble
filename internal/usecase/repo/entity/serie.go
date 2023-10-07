package entity

import "time"

type Serie struct {
	Serie        string    `json:"ser" bson:"_id,omitempty"`
	Strike       float64   `json:"prEx" bson:"strike,omitempty"`
	RazSoc       string    `json:"razSocEmi" bson:"razao_social,omitempty"`
	Tipo         string    `json:"tipo,omitempty" bson:"tipo,omitempty"`
	Vencimento   string    `json:"dtVen" bson:"vencimento,omitempty"`
	DtVencimento time.Time `bson:"tsVencimento,omitempty"`
	Price        float64   `json:"price" bson:"price,omitempty"`
}
