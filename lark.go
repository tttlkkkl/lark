package lark // import "github.com/tttlkkkl/lark"

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/tttlkkkl/lark/cache"
)

// Lark api client
type Lark struct {
	Options
	AppID          string
	AppSecret      string
	ReceiveMessage *ReceiveMessage
	Cache          cache.Cache
	HTTPClient     *http.Client
}

// Options 配置选项
type Options struct {
	// ReceiveMessageAPI 接收消息服务器配置
	ReceiveMessageAPI struct {
		Token      string
		EncryptKey string
	}
}

// Option 设置配置选项
type Option func(*Options)

// SetReceiveMessageAPI 设置消息接收服务器
func SetReceiveMessageAPI(token, encryptKey string) Option {
	return func(o *Options) {
		o.ReceiveMessageAPI.Token = token
		o.ReceiveMessageAPI.EncryptKey = encryptKey
	}
}

// NewLark 客户端初始化
func NewLark(appID, appSecret string, opt ...Option) (*Lark, error) {
	opts := new(Options)
	for _, o := range opt {
		o(opts)
	}
	var err error
	rm := &ReceiveMessage{}
	if opts.ReceiveMessageAPI.Token != "" && opts.ReceiveMessageAPI.EncryptKey != "" {
		rm, err = newReceiveMessage(opts.ReceiveMessageAPI.Token, opts.ReceiveMessageAPI.EncryptKey)
		if err != nil {
			return nil, err
		}
	}
	return &Lark{
		Options:        *opts,
		AppID:          appID,
		AppSecret:      appSecret,
		ReceiveMessage: rm,
		Cache:          cache.NewMemory(),
		HTTPClient:     http.DefaultClient,
	}, nil
}

// SendMessage 发送一条自定义结构的消息
func (l *Lark) SendMessage(message interface{}) MessageResponse {
	return l.send(messageSendURL, message)
}

// SendBatchMessage 批量发送自定义结构的消息
func (l *Lark) SendBatchMessage(message interface{}) MessageResponse {
	return l.send(messageBatchSendURL, message)
}

func (l *Lark) send(url string, message interface{}) MessageResponse {
	repBody := MessageResponse{}
	s, err := json.Marshal(message)
	if err != nil {
		Log.Error("消息序列化失败", err)
		repBody.Code = -1
		repBody.Message = err.Error()
		return repBody
	}
	reqBody := bytes.NewBuffer(s)
	body, err := l.httpPost(url, reqBody)
	err = json.Unmarshal(body, &repBody)
	if err != nil {
		repBody.Code = -1
		repBody.Message = err.Error()
		return repBody
	}
	return repBody
}

func (l *Lark) httpGet(url string) ([]byte, error) {
	rep, err := l.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(rep.Body)
}

// CustomHeader 自定义 http header
type CustomHeader struct {
	Key string
	Val string
}

func (l *Lark) httpPost(url string, body io.Reader, h ...CustomHeader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	l.setHeader(req, h...)
	rep, err := l.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()
	return ioutil.ReadAll(rep.Body)
}

// Post 发送 post 请求，自动附加 token
func (l *Lark) Post(url string, body io.Reader, h ...CustomHeader) ([]byte, error) {
	return l.httpPost(url, body, h...)
}

func (l *Lark) setHeader(req *http.Request, h ...CustomHeader) {
	for _, v := range h {
		if v.Key != "" {
			req.Header.Set(v.Key, v.Val)
		}
	}
	// 尝试自动加注鉴权信息
	if req.Header.Get("Authorization") == "" {
		tk, err := l.GetAccessToken()
		if err != nil {
			Log.Error("无法获取正确的 token", err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+tk)
	}
}
