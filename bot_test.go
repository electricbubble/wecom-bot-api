package wecom_bot_api

import (
	"github.com/electricbubble/wecom-bot-api/md"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var botKey = ""
var phoneNumber = ""
var userid = ""
var bot WeComBot

func setup(t *testing.T) {
	botKey = os.Getenv("WeCom_Bot_Key")

	phoneNumber = os.Getenv("Phone_Number")
	userid = os.Getenv("Userid")

	var err error
	bot, err = NewWeComBot(botKey)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_PushTextMessage(t *testing.T) {
	setup(t)

	// err := bot.PushTextMessage("广州今日天气：29度，大部分多云，降雨概率：60%",
	// 	MentionByUserid("wangqing"), MentionAllByUserid(),
	// 	MentionByMobile("13800001111"), MentionAllByMobile(),
	// )
	err := bot.PushTextMessage("hi again",
		MentionByMobile(phoneNumber),
		MentionByUserid(userid),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_PushMarkdownMessage(t *testing.T) {
	setup(t)
	SetDebug(true)

	// err := bot.PushMarkdownMessage(md.Heading(1, "H1") + "实时新增用户反馈" + md.WarningText("132例") + "，请相关同事注意。\n" +
	// 	md.QuoteText("类型:"+md.CommentText("用户反馈")) +
	// 	md.QuoteText("普通用户反馈:"+md.CommentText("117例")) +
	// 	md.QuoteText("VIP用户反馈:"+md.CommentText("15例")),
	// )
	err := bot.PushMarkdownMessage(
		md.Heading(1, "H1") + "实时新增用户反馈" + md.WarningText("132例") + "，请相关同事注意。\n" +
			md.QuoteText("类型:"+md.CommentText("用户反馈")) +
			md.QuoteText("普通用户反馈:"+md.CommentText("117例")) +
			md.QuoteText("VIP用户反馈:"+md.CommentText("15例")) +
			md.MentionByUserid(userid),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_PushImageMessage(t *testing.T) {
	setup(t)
	// SetDebug(true)

	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")

	readFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	err = bot.PushImageMessage(readFile)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_PushNewsMessage(t *testing.T) {
	setup(t)
	// SetDebug(true)

	article := NewArticle("中秋节礼品领取", "www.qq.com",
		ArticleDescription("今年中秋节公司有豪礼相送"),
		ArticlePicUrl("http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"),
	)
	article2 := NewArticle("图文标题2", "www.qq.com",
		ArticleDescription("图文描述2"),
	)
	article3 := NewArticle("图文标题3", "www.qq.com",
		ArticleDescription("图文描述3"),
	)
	_, _ = article2, article3
	// err := bot.PushNewsMessage(article)
	err := bot.PushNewsMessage(article, article2, article3)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_PushFileMessage(t *testing.T) {
	setup(t)
	SetDebug(true)

	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")

	media, err := bot.UploadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = bot.PushFileMessage(media)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_UploadFile(t *testing.T) {
	setup(t)
	SetDebug(true)

	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")

	media, err := bot.UploadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(media)
}
