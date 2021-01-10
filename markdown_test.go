package wecom_bot_api

import (
	"encoding/json"
	"github.com/electricbubble/wecom-bot-api/md"
	"testing"
)

func Test_newMarkdownMessage(t *testing.T) {
	msg := newMarkdownMsg("实时新增用户反馈" + md.WarningText("132例") + "，请相关同事注意。\n" +
		md.QuoteText("类型:"+md.CommentText("用户反馈")) +
		md.QuoteText("普通用户反馈:"+md.CommentText("117例")) +
		md.QuoteText("VIP用户反馈:"+md.CommentText("15例")),
	)

	marshal, err := json.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}
