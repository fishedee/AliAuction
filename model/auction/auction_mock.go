package auction

import ()

type IAuctionAo interface {
	GetItemList(where ItemListWhere) []Item
}

type AuctionAoMock struct {
	GetItemListHandler func(where ItemListWhere) []Item
}

func (this *AuctionAoMock) GetItemList(where ItemListWhere) []Item {
	return this.GetItemListHandler(where)
}
