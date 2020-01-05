package models

type Process struct {
	TotalValue   float64
	Items        []Item
	Installments int64
	Seller       Seller
}

type Item struct {
	Name  string
	Value float64
}

type Seller struct {
	Name    string
	CNPJ    string
	Address SellerAddress
}

type SellerAddress struct {
	Street  string
	Number  int64
	ZipCode string
}
