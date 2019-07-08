package wxpush

import (
	. "aliauction/model/cache"
	. "github.com/fishedee/app/log"
	. "github.com/fishedee/app/timer"
	"testing"
)

func TestWxPushAo(t *testing.T) {
	log, err := NewLog(LogConfig{})
	if err != nil {
		panic(err)
	}
	timer, err := NewTimer(log)
	if err != nil {
		panic(err)
	}
	cacheData := map[string]string{}
	cacheAo := &CacheAoMock{
		GetHandler: func(name string) string {
			return cacheData[name]
		},
		SetHandler: func(name string, value string) {
			cacheData[name] = value
		},
	}
	wxPushAo := NewWxPushAo(cacheAo, WxPushConfig{
		CorpId:  "wwf63fcb2f3002fe28",
		AgentId: 1000002,
		Secert:  "tziieNCWzo1rHT4LjuTOh0btL29CJ9XgW2wgj-RFzJU",
	}, timer).(*WxPushAo)

	wxPushAo.refreshAccessToken()
	wxPushAo.Push("标题", "描述", "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1562588826387&di=5aa7a5df81aab1d2447797f148d8ee6d&imgtype=0&src=http%3A%2F%2Fs10.sinaimg.cn%2Fmw690%2F005Cwy9hzy7d9Vym6wx09%26690", "http://www.baidu.com")

}
