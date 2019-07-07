package auction

import (
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
)

type AuctionAo struct {
}

func (this *AuctionAo) GetItemList(where ItemListWhere) []Item {
	data := this.fetchItemList(where)
	return this.analyseItemList(data)
}

func (this *AuctionAo) fetchItemList(where ItemListWhere) []byte {
	query := map[string]interface{}{
		"auction_source": 0,
		"st_param":       -1,
	}
	if where.Category != 0 {
		query["category"] = where.Category
	}
	if where.City != "" {
		query["city"] = where.City
	}
	if where.StartSeg != 0 {
		query["auction_start_seg"] = where.StartSeg
	}
	queryUrl, err := EncodeUrlQuery(query)
	if err != nil {
		Throw(1, err.Error())
	}
	var data []byte
	err = DefaultAjaxPool.Get(&Ajax{
		Url: "https://sf.taobao.com/item_list.htm?" + string(queryUrl),
		Header: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36",
		},
		ResponseData: &data,
	})
	if err != nil {
		Throw(1, err.Error())
	}
	return data
}

func (this *AuctionAo) analyseItemList(data []byte) []Item {
	return nil
}

func NewAuctionAo() IAuctionAo {
	return &AuctionAo{}
}
