package wecom_bot_api

import (
	"encoding/json"
	"testing"
)

func TestNewTextMessage(t *testing.T) {
	msg := newTextMsg("广州今日天气：29度，大部分多云，降雨概率：60%",
		MentionByUserid("wangqing"), MentionAllByUserid(),
		MentionByMobile("13800001111"), MentionAllByMobile(),
	)

	marshal, err := json.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}
