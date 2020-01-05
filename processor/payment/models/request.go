package models

type Request struct {
	Token      string   `json:"token"`
	Card       CardData `json:"card"`
	Process    Process  `json:"process"`
	AcquirerID int64    `json:"acquirer_id"`
}
