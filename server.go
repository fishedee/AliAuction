package main

import (
	. "aliauction/model/auction"
	. "aliauction/model/cache"
	. "aliauction/model/wxpush"
	_ "aliauction/util"
	"fmt"
	. "github.com/fishedee/app/ioc"
	. "github.com/fishedee/app/log"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	"strconv"
	"time"
)

func crawlData(task chan Item, auctionAo IAuctionAo, cacheAo ICacheAo, log Log) {
	defer CatchCrash(func(e Exception) {
		log.Critical("Crash CrawlData![Code:%v][Msg:%v][Stack:%v]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
	})
	defer Catch(func(e Exception) {
		log.Error("Error CrawlData![Code:%v][Msg:%v][Stack:%v]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
	})
	log.Debug("try to crawl data")
	itemListWhere := []ItemListWhere{
		//佛山工业厂房
		ItemListWhere{
			Category: 200788003,
			City:     "%B7%F0%C9%BD",
		},
		//佛山商业用房
		ItemListWhere{
			Category: 200782003,
			City:     "%B7%F0%C9%BD",
		},
		//龙江住宅用房
		ItemListWhere{
			Category:     50025969,
			LocationCode: 440606,
		},
		//禅城住宅用房
		ItemListWhere{
			Category:     50025969,
			LocationCode: 440604,
		},
	}
	for _, where := range itemListWhere {
		for i := 0; i != 10; i++ {
			newWhere := where
			if i != 0 {
				newWhere.Page = i + 1
			}
			data := auctionAo.GetItemList(newWhere)
			for _, single := range data {
				key := strconv.Itoa(single.Id)
				hasData := cacheAo.Has(key)
				if hasData == false {
					jsonData, err := EncodeJson(single)
					if err != nil {
						panic(err)
					}
					cacheAo.Set(key, string(jsonData))
					task <- single
				}
			}
		}
	}
	close(task)
}

func formatPrice(a Decimal) string {
	if a.Cmp("100000000") >= 0 {
		a = a.Div("100000000")
		return a.String() + "亿元"
	} else if a.Cmp("10000") >= 0 {
		a = a.Div("10000")
		return a.String() + "万元"
	} else {
		return a.String() + "元"
	}
}

func pushData(task chan Item, wxPushAo IWxPushAo, log Log) {
	defer CatchCrash(func(e Exception) {
		log.Critical("Crash PushData![Code:%v][Msg:%v][Stack:%v]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
	})
	defer Catch(func(e Exception) {
		log.Error("Error PushData![Code:%v][Msg:%v][Stack:%v]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
	})
	for {
		item, isOk := <-task
		if isOk == false {
			break
		}
		start := time.Unix(int64(item.Start)/1000, int64(item.Start)%1000).Format("2006-01-02 15:04:05")
		end := time.Unix(int64(item.End)/1000, int64(item.End)%1000).Format("2006-01-02 15:04:05")
		title := fmt.Sprintf("当前价:%v", formatPrice(item.CurrentPrice))
		description := fmt.Sprintf("[状态:%v][时间:%v~%v]，%v", item.Status, start, end, item.Title)
		sourceUrl := fmt.Sprintf("https://h5.m.taobao.com/paimai/detail/detailV2.html?type=2&itemId=%v", item.Id)
		picUrl := "https:" + item.PicUrl
		wxPushAo.Push(title, description, picUrl, sourceUrl)
	}
}

//go:generate mock ^./model/.*/.*(ao|db)\.go$ ^.*(Ao|Db)$
func main() {
	MustInvokeIoc(func(auctionAo IAuctionAo, cacheAo ICacheAo, wxPushAo IWxPushAo, log Log) {
		//执行逻辑
		itemChan := make(chan Item, 1024)
		go crawlData(itemChan, auctionAo, cacheAo, log)
		pushData(itemChan, wxPushAo, log)

		//等待
		time.Sleep(time.Minute * 30)
	})
}
