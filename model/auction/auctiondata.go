package auction

import (
	. "github.com/fishedee/language"
)

type ItemListWhere struct {
	Category int
	City     string
	StartSeg int
}

type Item struct {
	Id           int
	Status       string
	ItemUrl      string
	Title        string
	PicUrl       string
	InitialPrice Decimal
	CurrentPrice Decimal
	ConsultPrice Decimal
	Start        int
	End          int
	TimeToStart  int
	TimeToEnd    int
	ViewerCount  int
	BidCount     int
}
