package wxpush

type WxPushConfig struct {
	CorpId  string `config:"corpid"`
	AgentId int    `config:"agentid"`
	Secert  string `config:"secert"`
}
