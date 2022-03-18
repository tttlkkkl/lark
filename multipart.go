package lark

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

// UploudImaesResponse 图片上传返回
type UploudImaesResponse struct {
	Response
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data"`
}

// UploudImages lark 上传消息图片附件
// filePath 文件路径或者名称，fileBody 为 nil 时尝试从 filePath 中读取文件内容
// fileBody 文件内容，当 fileBody 不为 nil 时直接作为文件内容，此时不再尝试从 filePath 读取文件
func (l *Lark) UploudImages(imageType ImageType, filePath string, fileBody io.Reader) (*UploudImaesResponse, error) {
	var rs = UploudImaesResponse{}
	if fileBody == nil && filePath == "" {
		return &rs, errors.New("请提供文件路径或者文件内容")
	}
	var err error
	if fileBody == nil {
		fp, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer fp.Close()
		fileBody = fp
	}
	var body bytes.Buffer
	wtr := multipart.NewWriter(&body)
	part, err := wtr.CreateFormFile("image", filepath.Base(filePath))
	if err != nil {
		return &rs, err
	}
	if _, err = io.Copy(part, fileBody); err != nil {
		return nil, err
	}
	if err = wtr.WriteField("image_type", string(imageType)); err != nil {
		return nil, err
	}
	h := []CustomHeader{
		{
			Key: "Content-Type",
			Val: wtr.FormDataContentType(),
		},
		{
			Key: "Content-Length",
			Val: strconv.FormatInt(int64(body.Len()), 10),
		},
		{
			Key: "User-Agent",
			Val: "go lark lib",
		},
	}
	// 这里不关闭的话会缺少 boundary 封闭付号，导致对端无法解析
	// defer wtr.close() 也不行
	wtr.Close()
	btRs, err := l.httpPost(uploudImaesURL, &body, h...)
	if err != nil {
		return &rs, err
	}
	if err = json.Unmarshal(btRs, &rs); err != nil {
		return &rs, err
	}
	return &rs, nil
}
