package wxpush

import ()

type IWxPushAo interface {
	Push(title string, description string, imageUrl string, sourceUrl string)
}

type WxPushAoMock struct {
	PushHandler func(title string, description string, imageUrl string, sourceUrl string)
}

func (this *WxPushAoMock) Push(title string, description string, imageUrl string, sourceUrl string) {
	this.PushHandler(title, description, imageUrl, sourceUrl)
}
