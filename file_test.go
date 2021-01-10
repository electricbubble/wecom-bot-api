package wecom_bot_api

import (
	"encoding/json"
	"testing"
)

func Test_newFileMsg(t *testing.T) {
	msg := newFileMsg("3a8asd892asd8asd")

	marshal, err := json.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}
