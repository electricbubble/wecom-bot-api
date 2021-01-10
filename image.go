package wecom_bot_api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

type imageMsg struct {
	// 消息类型, 固定为 `image`
	MsgType string    `json:"msgtype"`
	Image   imageData `json:"image"`
}

type imageData struct {
	// 图片内容的 `base64` 编码
	Base64 string `json:"base64"`

	// 图片内容（ `base64` 编码前）的 `md5` 值
	Md5 string `json:"md5"`
}

// 注：图片（base64编码前）最大不能超过2M，支持JPG,PNG格式

func newImageMsg(img []byte) imageMsg {
	encodeToString := base64.StdEncoding.EncodeToString(img)
	hash := md5.New()
	hash.Write(img)
	toString := hex.EncodeToString(hash.Sum(nil))

	msg := imageMsg{
		MsgType: "image",
		Image:   imageData{Base64: encodeToString, Md5: toString},
	}
	return msg
}
