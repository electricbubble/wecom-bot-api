package wecom_bot_api

type markdownMsg struct {
	// 消息类型, 固定为 `markdown`
	MsgType  string       `json:"msgtype"`
	Markdown markdownData `json:"markdown"`
}

type markdownData struct {
	// `markdown` 内容, 最长不超过 4096 个字节, 必须是 `utf8` 编码
	Content string `json:"content"`
}

func newMarkdownMsg(content string) markdownMsg {
	msg := markdownMsg{
		MsgType:  "markdown",
		Markdown: markdownData{Content: content},
	}
	return msg
}
