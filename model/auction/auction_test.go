package auction

import (
	. "github.com/fishedee/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestFetchItemList(t *testing.T) {
	auctionAo := NewAuctionAo().(*AuctionAo)
	data := auctionAo.fetchItemList(ItemListWhere{
		Category: 50025969,
		City:     "%B7%F0%C9%BD",
		StartSeg: -1,
	})
	err := ioutil.WriteFile("testdata/itemList", data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	AssertEqual(t, len(data) != 0, true)
}

func TestAnalyseItemList(t *testing.T) {
	auctionAo := NewAuctionAo().(*AuctionAo)
	input, err := ioutil.ReadFile("testdata/demo")
	if err != nil {
		panic(err)
	}
	data := auctionAo.analyseItemList(input)
	t.Log(data)
}
