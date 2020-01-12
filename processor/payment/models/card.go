package models

type CardSensitiveData struct {
	Number string
	CVV    string
}

type CardOpenData struct {
	Name string
	Flag string
	Date string
}

type CardData struct {
	Open      CardOpenData
	Sensitive CardSensitiveData
}
