package lark

import (
	"encoding/json"
	"strings"
)

// Response 公共返回结果
type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

// StringArray 数组字符串
type StringArray []string

// MarshalJSON 自定义数组字符串json编码
func (s *StringArray) MarshalJSON() ([]byte, error) {
	str := strings.Join(*s, "|")
	return json.Marshal(str)
}

// UnmarshalJSON 自定义数组字符串json解码
func (s *StringArray) UnmarshalJSON(b []byte) error {
	strArr := strings.Split(string(b), "|")
	*s = strArr
	return nil
}

// IsSuccess 是否成功
func (r *Response) IsSuccess() bool {
	return r.Code == 0 && (strings.ToLower(r.Message) == "ok" || strings.ToLower(r.Message) == "success")
}
