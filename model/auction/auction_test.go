package auction

import (
	. "github.com/fishedee/assert"
	"testing"
)

func TestFetchItemList(t *testing.T) {
	auctionAo := NewAuctionAo().(*AuctionAo)
	data := auctionAo.fetchItemList(ItemListWhere{
		Category: 50025969,
		City:     "佛山",
		StartSeg: -1,
	})
	t.Log(string(data))
	AssertEqual(t, len(data) != 0, true)
}
