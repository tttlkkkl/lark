package lark

const (
	// accessTokenURL accessToken 接口地址
	accessTokenURL      = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	messageSendURL      = "https://open.feishu.cn/open-apis/message/v4/send/"
	messageBatchSendURL = "https://open.feishu.cn/open-apis/message/v4/batch_send/"
	chatListURL         = "https://open.feishu.cn/open-apis/chat/v4/list"
	cardUpdateURL       = "https://open.feishu.cn/open-apis/interactive/v1/card/update/"
)

// MT 事件消息类型
type MT string

const (
	// MTURLVerification xx
	MTURLVerification MT = "url_verification"
	// MTEventCallback 回调事件
	MTEventCallback MT = "event_callback"
)

// MsgType 消息类型
type MsgType string

const (
	// MsgTypeCard 卡片消息类型
	MsgTypeCard MsgType = "interactive"
	// MsgTypeText 文本消息
	MsgTypeText MsgType = "text"
	// MsgTypeImage 图片消息
	MsgTypeImage MsgType = "image"
	// MsgTypePOST 富文本消息
	MsgTypePOST MsgType = "post"
	// MsgTypeShareChat 群名片消息
	MsgTypeShareChat MsgType = "share_chat"
)
