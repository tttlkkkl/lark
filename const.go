package lark

const (
	// appAccessTokenURL accessToken 接口地址
	appAccessTokenURL    = "https://open.larksuite.com/open-apis/auth/v3/app_access_token/internal/"
	tenantAccessTokenURL = "https://open.larksuite.com/open-apis/auth/v3/tenant_access_token/internal/"
	messageSendURL       = "https://open.larksuite.com/open-apis/message/v4/send/"
	messageBatchSendURL  = "https://open.larksuite.com/open-apis/message/v4/batch_send/"
	chatListURL          = "https://open.larksuite.com/open-apis/chat/v4/list"
	cardUpdateURL        = "https://open.larksuite.com/open-apis/interactive/v1/card/update/"
	getUserURL           = "https://open.larksuite.com/open-apis/contact/v3/users/%s"
	uploudImaesURL       = "https://open.larksuite.com/open-apis/im/v1/images"
	copyFileURL          = "https://open.larksuite.com/open-apis/drive/explorer/v2/file/copy/files/%s"
	updatePermissionsURL = "https://open.larksuite.com/open-apis/drive/v1/permissions/%s/public?type=%s"
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

// ImageType 图片用途类型
type ImageType string

const (
	// ImageTypeMessage 图片用于发送消息
	ImageTypeMessage = "message"
	// ImageTypeAvatar 图片用作头像
	ImageTypeAvatar = "avatar"
)

// DocType 云文档类型
type DocType string

const (
	DocTypeDoc     DocType = "doc"
	DocTypeSheet   DocType = "sheet"
	DocTypeDocx    DocType = "docx"
	DocTypeFile    DocType = "file"
	DocTypeWiki    DocType = "wiki"
	DocTypeBitable DocType = "bitable"
)
