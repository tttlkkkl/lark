package lark

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReceiveMessage 事件与消息接收
type ReceiveMessage struct {
	Token      string
	EncryptKey string
}

// Encrypt 加密串解析
type Encrypt struct {
	Encrypt string `json:"encrypt"`
}

// EventMessage 事件响应公共字段
type EventMessage struct {
	Token       string `json:"token"`
	Challenge   string `json:"challenge"`
	MessageType MT     `json:"type"`
	UUID        string `json:"uuid"`
	Ts          string `json:"ts"`
	Event       Values `json:"event"`
	Schema      string `json:"schema"`
	Header      Header `json:"header"`
}

// CardMessage 卡片消息
type CardMessage struct {
	OpenID        string `json:"open_id"`
	UserID        string `json:"user_id"`
	OpenMessageID string `json:"open_message_id"`
	TenantKey     string `json:"tenant_key"`
	Token         string `json:"token"`
	Action        Action `json:"action"`
	Challenge     string `json:"challenge"`
}

// Action 操作类型
type Action struct {
	Val Values `json:"value"`
	Tag string `json:"tag"`
}

// Header 2.0 事件 header
type Header struct {
	EventID    string        `json:"event_id"`
	Token      string        `json:"token"`
	CreateTime string        `json:"create_time"`
	EventType  MT            `json:"event_type"`
	TenantKey  string        `json:"tenant_key"`
	AppID      string        `json:"app_id"`
	ResourceID string        `json:"resource_id"`
	UserList   []interface{} `json:"user_list"`
}

// NewReceiveMessage 消息事件处理
func newReceiveMessage(token, encryptKey string) (*ReceiveMessage, error) {
	e := &ReceiveMessage{
		Token:      token,
		EncryptKey: encryptKey,
	}
	return e, nil
}

// Handle 消息接收处理
func (e *ReceiveMessage) Handle(req *http.Request, rep http.ResponseWriter) (*EventMessage, error) {
	postData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Log.Error("获取 http body 失败:", err)
		return nil, err
	}
	Log.Debug("收到原始数据:", string(postData))
	data := e.getEncryptString(postData)
	Log.Debug("展开后的数据:", string(data))
	var msg EventMessage
	err = json.Unmarshal(data, &msg)
	if err != nil {
		Log.Error("json 解包失败:", err)
		return nil, err
	}
	// Verification Token 比对
	if msg.GetToken() != e.Token {
		return nil, errors.New("Verification Token 验证失败，请确认配置")
	}
	// url 检验自动返回
	if msg.Challenge != "" {
		rep.Header().Set("Content-Type", "application/json")
		rep.Write([]byte(fmt.Sprintf(`{"challenge":"%s"}`, msg.Challenge)))
	}
	return &msg, nil
}

// HandleCard 消息接收处理
func (e *ReceiveMessage) HandleCard(req *http.Request, rep http.ResponseWriter) (*CardMessage, error) {
	postData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Log.Error("获取 http body 失败:", err)
		return nil, err
	}
	Log.Debug("收到原始数据:", string(postData))
	data := e.getEncryptString(postData)
	Log.Debug("展开后的数据:", string(data))
	var msg CardMessage
	err = json.Unmarshal(data, &msg)
	if err != nil {
		Log.Error("json 解包失败:", err)
		return nil, err
	}
	// url 检验自动返回
	if msg.Challenge != "" {
		rep.Header().Set("Content-Type", "application/json")
		rep.Write([]byte(fmt.Sprintf(`{"challenge":"%s"}`, msg.Challenge)))
	}
	return &msg, nil
}

func (e *ReceiveMessage) getEncryptString(data []byte) []byte {
	var en Encrypt
	err := json.Unmarshal(data, &en)
	if err == nil && en.Encrypt != "" {
		Log.Debug("启用解密，解密前字符串:", en.Encrypt)
		aesRs, err := AESDecrypt(en.Encrypt, []byte(e.EncryptKey))
		if err != nil {
			Log.Error("解密失败:", err)
			return data
		}
		return aesRs
	}
	return data
}

// GetToken 获取 Verification Token
func (e *EventMessage) GetToken() string {
	if e.Schema == "2.0" {
		return e.Header.Token
	}
	return e.Token
}

// GetToken 获取 Verification Token
func (c *CardMessage) GetToken() string {
	return c.Token
}

// GetMessageType 获取事件消息类型
func (e *EventMessage) GetMessageType() MT {
	if e.Schema == "2.0" {
		return e.Header.EventType
	}
	return e.MessageType
}
