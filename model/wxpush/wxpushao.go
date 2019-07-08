package wxpush

import (
	. "aliauction/model/cache"
	. "github.com/fishedee/app/ioc"
	. "github.com/fishedee/app/timer"
	. "github.com/fishedee/sdk"
	//. "github.com/fishedee/util"
	"time"
)

type WxPushAo struct {
	cacheAo ICacheAo
	config  WxPushConfig
}

func NewWxPushAo(cacheAo ICacheAo, config WxPushConfig, timer Timer) IWxPushAo {
	wxPushAo := &WxPushAo{
		cacheAo: cacheAo,
		config:  config,
	}
	wxPushAo.refreshAccessToken()
	timer.MustCron("0 */10 * * * *", wxPushAo.refreshAccessToken)
	return wxPushAo
}

func (this *WxPushAo) refreshAccessToken() {
	now := time.Now().Add(time.Minute * 30)
	wxAccessTokenExpire := time.Time{}
	wxAccessTokenExpireStr := this.cacheAo.Get("accessTokenExpire")
	if wxAccessTokenExpireStr != "" {
		var err error
		wxAccessTokenExpire, err = time.ParseInLocation("2006-01-02T15:04:05", wxAccessTokenExpireStr, time.Local)
		if err != nil {
			wxAccessTokenExpire = time.Time{}
		}
	}
	if wxAccessTokenExpire.IsZero() == false &&
		wxAccessTokenExpire.After(now) {
		//尚未超时
		return
	}

	wxQySdk := &WxQySdk{
		CorpId:  this.config.CorpId,
		AgentId: this.config.AgentId,
		Secert:  this.config.Secert,
	}
	token, err := wxQySdk.GetAccessToken()
	if err != nil {
		panic(err)
	}
	accessToken := token.AccessToken
	expireTime := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	this.cacheAo.Set("accessTokenExpire", expireTime.Format("2006-01-02T15:04:05"))
	this.cacheAo.Set("accessToken", accessToken)
}

func (this *WxPushAo) Push(title string, description string, imageUrl string, sourceUrl string) {
	wxQySdk := &WxQySdk{
		CorpId:      this.config.CorpId,
		AgentId:     this.config.AgentId,
		Secert:      this.config.Secert,
		AccessToken: this.cacheAo.Get("accessToken"),
	}
	//群发图文消息
	_, err := wxQySdk.SendMessage(WxQySdkSendMessage{
		ToUser:  "@all",
		MsgType: "news",
		News: WxQySdkSendNewsMessage{
			Articles: []WxQySdkSendNewsArticleMessage{
				WxQySdkSendNewsArticleMessage{
					Title:       title,
					Description: description,
					Url:         sourceUrl,
					PicUrl:      imageUrl,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func init() {
	MustRegisterIoc(NewWxPushAo)
}
