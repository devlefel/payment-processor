package models

type Request struct {
	Token      string   `json:"token"`
	Card       CardData `json:"card"`
	Process    Process  `json:"process"`
	AcquirerID int64    `json:"acquirer_id"`
}

type Response struct {
	Success bool  `json:"success"`
	Errors  Error `json:"errors"`
}
