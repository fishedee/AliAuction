package auction

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	. "github.com/fishedee/app/ioc"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
	"strings"
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
		"provice":        "",
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
	url := []string{}
	for key, value := range query {
		url = append(url, fmt.Sprintf("%s=%v", key, value))
	}
	var data []byte
	err := DefaultAjaxPool.Get(&Ajax{
		Url: "https://sf.taobao.com/item_list.htm?" + Implode(url, "&"),
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

func (this *AuctionAo) analyseItemList(dataGbk []byte) []Item {
	dec := mahonia.NewDecoder("gbk")
	data := dec.ConvertString(string(dataGbk))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		Throw(1, err.Error())
	}
	script := doc.Find("#sf-item-list-data").Text()
	var result struct {
		Data []Item
	}
	err = DecodeJson([]byte(script), &result)
	if err != nil {
		Throw(1, err.Error())
	}
	return result.Data
}

func NewAuctionAo() IAuctionAo {
	return &AuctionAo{}
}

func init() {
	MustRegisterIoc(NewAuctionAo)
}
