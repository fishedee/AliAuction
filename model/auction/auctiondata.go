package auction

import (
	. "github.com/fishedee/language"
	"time"
)

type ItemListWhere struct {
	Category int
	City     string
	StartSeg int
}

type Item struct {
	Id           string
	Status       string
	ItemUrl      string
	Title        string
	PicUrl       string
	InitialPrice Decimal
	CurrentPrice Decimal
	ConsultPrice Decimal
	TimeToStart  time.Time
	TimeToEnd    time.Time
}
