package wecom_bot_api

type fileMsg struct {
	// 消息类型, 固定为 `file`
	MsgType string `json:"msgtype"`
	File    Media  `json:"file"`
}

type Media struct {
	// 文件id, 通过文件上传接口获取
	ID string `json:"media_id"`
}

func newFileMsg(mediaID string) fileMsg {
	msg := fileMsg{
		MsgType: "file",
		File:    Media{ID: mediaID},
	}
	return msg
}
