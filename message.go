package lark

import "encoding/json"

// MessageHeader 消息公共头部
type MessageHeader struct {
	OpenID      string  `json:"open_id"`
	UserID      string  `json:"user_id"`
	Email       string  `json:"email"`
	ChatID      string  `json:"chat_id"`
	RootID      string  `json:"root_id"`
	MessageType MsgType `json:"msg_type"`
}

// MessageBatchHeader 消息公共头部
type MessageBatchHeader struct {
	DepartmentIDs StringArray `json:"department_ids"`
	OpenIDs       StringArray `json:"open_ids"`
	UserIDs       StringArray `json:"user_ids"`
	MessageType   MsgType     `json:"msg_type"`
}

// MessageResponse 消息发送返回
type MessageResponse struct {
	Response
	Data struct {
		MessageID            string      `json:"data"`
		InvalidDepartmentIDs StringArray `json:"invalid_department_ids"`
		InvalidOpenIDs       StringArray `json:"invalid_open_ids"`
		InvalidUserIDs       StringArray `json:"invalid_user_ids"`
	} `json:"data"`
}

// MessageCard 卡片消息
type MessageCard struct {
	MessageHeader
	Card        interface{} `json:"card"`
	UpdateMulti bool        `json:"update_multi"`
}

// NewMessageCard 实例化一个卡片消息
// card 卡片结构 json 字符串
// isShare 是否共享卡片即 update_multi
func NewMessageCard(card string, isShare bool) (MessageCard, error) {
	m := MessageCard{
		MessageHeader: MessageHeader{MessageType: MsgTypeCard},
		UpdateMulti:   isShare,
	}
	var cd Values
	err := json.Unmarshal([]byte(card), &cd)
	if err != nil {
		return m, err
	}
	m.Card = cd
	return m, nil
}
