package serverchan

import (
	"testing"
)

func TestPush(t *testing.T) {
	serverChan := NewServerChanAo(ServerChanConfig{
		ScKey: "SCU54877T63422d44343becf98692a1039342e6b75d21f6fbb62e2",
	})

	serverChan.Push("我是标题", "# 我是正文标题\n\n ![图片](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1562517932916&di=db8a301d57a14204fe09a0b54c283303&imgtype=0&src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201701%2F18%2F20170118234754_hxsFf.jpeg)")
}
