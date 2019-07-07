package main

import (
	. "aliauction/model/auction"
	. "aliauction/model/cache"
	. "aliauction/model/serverchan"
	_ "aliauction/util"
	"fmt"
	. "github.com/fishedee/app/ioc"
	. "github.com/fishedee/app/log"
	. "github.com/fishedee/encoding"
	. "github.com/fishedee/language"
	"strconv"
	"time"
)

//go:generate mock ^./model/.*/.*(ao|db)\.go$ ^.*(Ao|Db)$
func main() {
	MustInvokeIoc(func(auctionAo IAuctionAo, cacheAo ICacheAo, serverChanAo IServerChanAo, log Log) {
		func() {
			defer CatchCrash(func(e Exception) {
				log.Critical("Crash![Code:%v][Msg:%v][Stack:%v]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
			})
			defer Catch(func(e Exception) {
				log.Error("Crash![Code:%v][Msg:%v][Stack:%v]", e.GetCode(), e.GetMessage(), e.GetStackTrace())
			})
			log.Debug("try to crawl data")
			data := auctionAo.GetItemList(ItemListWhere{
				Category: 50025969,
				City:     "%B7%F0%C9%BD",
				StartSeg: -1,
			})
			for _, single := range data {
				key := strconv.Itoa(single.Id)
				hasData := cacheAo.Has(key)
				if hasData == false {
					log.Debug("%v,hasData!")
					jsonData, err := EncodeJson(single)
					if err != nil {
						panic(err)
					}
					cacheAo.Set(key, string(jsonData))
					title := fmt.Sprintf("【当前价:%v】【状态:%v】%v", single.CurrentPrice, single.Status, single.Title)
					serverChanAo.Push(title, "")
				}
			}
		}()
		time.Sleep(time.Minute * 30)
		fmt.Println("Hello World")
	})
}
