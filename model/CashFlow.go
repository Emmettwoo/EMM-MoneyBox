package model

type CashFlow struct {
	Title  string `xml:"name,attr"`
	Amount float32
	Desc   string
}
