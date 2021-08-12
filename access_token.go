package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// AccessTokenResponse accessToken 请求返回
type AccessTokenResponse struct {
	Response
	AccessToken string `json:"tenant_access_token"`
	ExpiresIn   int64  `json:"expire"`
}

// PostParamsAccessToken access token 请求参数
type PostParamsAccessToken struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// GetAppAccessToken 获取 App Access Token
func (l *Lark) GetAppAccessToken() (string, error) {
	return l.getAccessToken(appAccessTokenURL, "a")
}

// GetTenantAccessToken 获取 tenantAccessToken
func (l *Lark) GetTenantAccessToken() (string, error) {
	return l.getAccessToken(tenantAccessTokenURL, "t")
}

// GetUserAccessToken 获取 tenantAccessToken
func (l *Lark) GetUserAccessToken() (string, error) {
	return l.getAccessToken(tenantAccessTokenURL, "u")
}

// getAccessToken 获取accessToken
func (l *Lark) getAccessToken(url, k string) (string, error) {
	token := l.Cache.Get(l.getAccessTokenCacheKey(k))
	if token != nil {
		accessToken, ok := token.(string)
		if ok {
			return accessToken, nil
		}
	}
	pm := PostParamsAccessToken{
		AppID:     l.AppID,
		AppSecret: l.AppSecret,
	}
	pmBt, _ := json.Marshal(&pm)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pmBt))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	r, err := l.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	var rep AccessTokenResponse
	err = json.Unmarshal(body, &rep)
	if err != nil {
		return "", err
	}
	if rep.Code != 0 {
		return "", err
	}
	rep.ExpiresIn = rep.ExpiresIn - 30
	l.Cache.Set(l.getAccessTokenCacheKey(k), rep.AccessToken, time.Second*time.Duration(rep.ExpiresIn))
	return rep.AccessToken, nil
}

func (l *Lark) getAccessTokenCacheKey(k string) string {
	return fmt.Sprintf("%s-%s", k, l.AppID)
}
