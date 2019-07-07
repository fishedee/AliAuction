package serverchan

import (
	. "github.com/fishedee/app/ioc"
	. "github.com/fishedee/language"
	. "github.com/fishedee/util"
)

type ServerChanAo struct {
	config ServerChanConfig
}

func NewServerChanAo(config ServerChanConfig) IServerChanAo {
	return &ServerChanAo{
		config: config,
	}
}

func (this *ServerChanAo) Push(title string, body string) {
	err := DefaultAjaxPool.Post(&Ajax{
		Url: "https://sc.ftqq.com/" + this.config.ScKey + ".send",
		Data: map[string]interface{}{
			"text": title,
			"desp": body,
		},
	})
	if err != nil {
		Throw(1, err.Error())
	}
}

func init() {
	MustRegisterIoc(NewServerChanAo)
}
