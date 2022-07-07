package lark // import "github.com/tttlkkkl/lark"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

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

// GetUser 获取单个用户信息
func (l *Lark) GetUser(req UserRequest) (rep UserResponse) {
	if req.UserID == "" {
		rep.Code = -1
		rep.Message = "UserID 是必须的"
		return
	}
	q := make(url.Values)
	url := fmt.Sprintf(getUserURL, req.UserID)
	if req.UserIDType != "" {
		q.Add("user_id_type", string(req.UserIDType))
	}
	if req.DepartmentIDType != "" {
		q.Add("department_id_type", string(req.DepartmentIDType))
	}
	rs, err := l.httpGet(url, q)
	if err != nil {
		rep.Code = -1
		rep.Message = err.Error()
		return
	}
	err = json.Unmarshal(rs, &rep)
	if err != nil {
		rep.Code = -1
		rep.Message = err.Error()
		return
	}
	return
}

// copyFileURL          = "https://open.larksuite.com/drive/explorer/v2/file/copy/files/%s"
// updatePermissionsURL = "https://open.larksuite.com/drive/v1/permissions/%s/public?type=%s"
func (l *Lark) CopyDocument(token string, req CopyRequest) CopyResponse {
	url := fmt.Sprintf(copyFileURL, token)
	rep := CopyResponse{}
	s, err := json.Marshal(req)
	if err != nil {
		Log.Error("消息序列化失败", err)
		rep.Code = -1
		rep.Message = err.Error()
		return rep
	}
	reqBody := bytes.NewBuffer(s)
	body, err := l.httpPost(url, reqBody)
	err = json.Unmarshal(body, &rep)
	if err != nil {
		rep.Code = -1
		rep.Message = err.Error()
		return rep
	}
	return rep
}

func (l *Lark) UpdatePermissionPublic(token string, fileType DocType, req UpdatePermissionRequest) UpdatePermissionResponse {
	url := fmt.Sprintf(updatePermissionsURL, token, fileType)
	rep := UpdatePermissionResponse{}
	s, err := json.Marshal(req)
	if err != nil {
		Log.Error("消息序列化失败", err)
		rep.Code = -1
		rep.Message = err.Error()
		return rep
	}
	reqBody := bytes.NewBuffer(s)
	body, err := l.httpPatch(url, reqBody)
	err = json.Unmarshal(body, &rep)
	if err != nil {
		rep.Code = -1
		rep.Message = err.Error()
		return rep
	}
	return rep
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

// TenantAccessToken 默认自动使用 TenantAccessToken 访问
func (l *Lark) httpGet(url string, q url.Values, h ...CustomHeader) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = q.Encode()
	l.setHeader(req, h...)
	rep, err := l.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()
	return ioutil.ReadAll(rep.Body)
}

// CustomHeader 自定义 http header
type CustomHeader struct {
	Key string
	Val string
}

// TenantAccessToken 默认自动使用 TenantAccessToken 访问
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

// TenantAccessToken 默认自动使用 TenantAccessToken 访问
func (l *Lark) httpPatch(url string, body io.Reader, h ...CustomHeader) ([]byte, error) {
	req, err := http.NewRequest("PATCH", url, body)
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

// Post 发送 post 请求，自动附加 TenantAccessToken
// 但不会覆盖自定义 header 中的 Authorization 字段
func (l *Lark) Post(url string, body io.Reader, h ...CustomHeader) ([]byte, error) {
	return l.httpPost(url, body, h...)
}

// Get 发送 get 请求，自动附加 TenantAccessToken
// 但不会覆盖自定义 header 中的 Authorization 字段
func (l *Lark) Get(url string, q url.Values, h ...CustomHeader) ([]byte, error) {
	return l.httpGet(url, q, h...)
}

func (l *Lark) setHeader(req *http.Request, h ...CustomHeader) {
	for _, v := range h {
		if v.Key != "" {
			req.Header.Set(v.Key, v.Val)
		}
	}
	// 尝试自动加注 tenant access token
	if req.Header.Get("Authorization") == "" {
		l.setTenantAccessToken(req)
	}
}

func (l *Lark) setAppAccessToken(req *http.Request) {
	tk, err := l.GetAppAccessToken()
	if err != nil {
		Log.Error("无法获取正确的 app access token", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+tk)
}

func (l *Lark) setTenantAccessToken(req *http.Request) {
	tk, err := l.GetTenantAccessToken()
	if err != nil {
		Log.Error("无法获取正确的 tenant access token", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+tk)
}
