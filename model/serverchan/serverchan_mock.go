package serverchan

import ()

type IServerChanAo interface {
	Push(title string, body string)
}

type ServerChanAoMock struct {
	PushHandler func(title string, body string)
}

func (this *ServerChanAoMock) Push(title string, body string) {
	this.PushHandler(title, body)
}
