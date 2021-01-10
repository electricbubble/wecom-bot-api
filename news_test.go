package wecom_bot_api

import (
	"encoding/json"
	"testing"
)

func Test_newNewsMessage(t *testing.T) {
	article := NewArticle("中秋节礼品领取", "www.qq.com",
		ArticleDescription("今年中秋节公司有豪礼相送"),
		ArticlePicUrl("http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"),
	)
	msg := newNewsMsg(article)
	// msg := newNewsMsg(article, article)

	marshal, err := json.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}
