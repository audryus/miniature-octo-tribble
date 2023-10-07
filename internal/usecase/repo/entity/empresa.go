package entity

type Empresa struct {
	Symbol string  `json:"id,omitempty" bson:"_id,omitempty"`
	RazSoc string  `json:"razSocEmi" bson:"razao_social,omitempty"`
	Tipo   string  `json:"tipo" bson:"tipo,omitempty"`
	Price  float64 `json:"price" bson:"price,omitempty"`
}
