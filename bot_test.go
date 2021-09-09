package wecom_bot_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/electricbubble/wecom-bot-api/md"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"
)

var botKey = ""
var phoneNumber = ""
var userid = ""
var bot WeComBot

func setup() {
	botKey = os.Getenv("WeCom_Bot_Key")

	phoneNumber = os.Getenv("Phone_Number")
	userid = os.Getenv("Userid")

	bot = NewWeComBot(botKey)
}

func Test_weComBot_PushTextMessage(t *testing.T) {
	setup()

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
	setup()
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
	setup()
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
	setup()
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
	setup()
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
	setup()
	SetDebug(true)

	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")

	media, err := bot.UploadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(media)
}

func TestTemplateCard(t *testing.T) {
	setup()

	var (
		req *http.Request
		err error
	)

	msg := bytes.NewBufferString(`{
    "msgtype":"template_card",
    "template_card":{
        "card_type":"text_notice",
        "source":{
            "icon_url":"https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
            "desc":"企业微信"
        },
        "main_title":{
            "title":"欢迎使用企业微信",
            "desc":"您的好友正在邀请您加入企业微信"
        },
        "emphasis_content":{
            "title":"100",
            "desc":"数据含义"
        },
        "sub_title_text":"下载企业微信还能抢红包！",
        "horizontal_content_list":[
            {
                "keyname":"邀请人",
                "value":"张三"
            },
            {
                "keyname":"企微官网",
                "value":"点击访问",
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi"
            }
        ],
        "jump_list":[
            {
                "type":1,
                "url":"https://work.weixin.qq.com/?from=openApi",
                "title":"企业微信官网"
            }
        ],
        "card_action":{
            "type":1,
            "url":"https://work.weixin.qq.com/?from=openApi",
            "appid":"APPID",
            "pagepath":"PAGEPATH"
        }
    }
}
`,
	).Bytes()

	rawUrl := "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87"

	tplCardText := newTemplateCardText(
		TemplateCardMainTitle("一级标题", "标题辅助信息"), TemplateCardActionUrl(rawUrl),
		// TemplateCardSource("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0", "企业微信"),
		TemplateCardEmphasisContent("关键数据标题", "关键数据描述"),
		TemplateCardSubTitleText("二级普通文本"),
		TemplateCardHorizontalContent("二级标题(text)", TemplateCardHorizontalContentText("二级文本")),
		TemplateCardHorizontalContent("二级标题(url)", TemplateCardHorizontalContentUrl(rawUrl, "api地址")),
		TemplateCardJump("跳转指引", TemplateCardJumpUrl(rawUrl)),
		TemplateCardJump("企业微信官网", TemplateCardJumpUrl("https://work.weixin.qq.com")),
	)
	tplCardMsg := newTemplateCardMsg(tplCardText)

	rawUrl = "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E5%9B%BE%E6%96%87%E5%B1%95%E7%A4%BA%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87"
	imgUrl := "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0"
	tplCardNews := newTemplateCardNews(
		TemplateCardMainTitle("一级标题", "标题辅助信息"), TemplateCardImage(imgUrl), TemplateCardActionUrl(rawUrl),
		TemplateCardSource("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0", "企业微信"),
		TemplateCardVerticalContent("卡片二级标题", "二级普通文本"),
		TemplateCardHorizontalContent("二级标题(text)", TemplateCardHorizontalContentText("二级文本")),
		TemplateCardHorizontalContent("二级标题(url)", TemplateCardHorizontalContentUrl(rawUrl, "api地址")),
		TemplateCardJump("跳转指引", TemplateCardJumpUrl(rawUrl)),
		TemplateCardJump("企业微信官网", TemplateCardJumpUrl("https://work.weixin.qq.com")),
	)
	tplCardMsg = newTemplateCardMsg(tplCardNews)

	bsData, err := json.MarshalIndent(tplCardMsg, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	msg = bsData

	fmt.Println(string(msg))

	// return

	if req, err = newRequest(http.MethodPost,
		fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", botKey),
		msg,
	); err != nil {
		t.Fatal(err)
	}
	rawResp, err := executeHTTP(req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(rawResp))
}

func Test_weComBot_PushTemplateCardTextNotice(t *testing.T) {
	setup()
	SetDebug(true)

	// userHomeDir, _ := os.UserHomeDir()
	// filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")
	//
	// media, err := bot.UploadFile(filename)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// "media_id":"38BHOWH1SHSCZImMcuPmG2TuJSpYikh0AxznKJYSUJAJaFJvDeRu60NTAuj_IKLoR"
	media := Media{ID: "38BHOWH1SHSCZImMcuPmG2TuJSpYikh0AxznKJYSUJAJaFJvDeRu60NTAuj_IKLoR"}

	rawUrl := "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E9%80%9A%E7%9F%A5%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87"

	err := bot.PushTemplateCardTextNotice(
		TemplateCardMainTitle("一级标题", "标题辅助信息"), TemplateCardActionUrl(rawUrl),
		TemplateCardSource("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0", "企业微信"),
		TemplateCardEmphasisContent("关键数据标题", "关键数据描述"),
		TemplateCardSubTitleText("二级普通文本"),
		TemplateCardHorizontalContent("二级标题(text)", TemplateCardHorizontalContentText("二级文本")),
		TemplateCardHorizontalContent("二级标题(url)", TemplateCardHorizontalContentUrl(rawUrl, "api地址")),
		TemplateCardHorizontalContent("二级标题(media)", TemplateCardHorizontalContentMedia("IMG_5246.jpg", media)),
		TemplateCardJump("跳转指引", TemplateCardJumpUrl(rawUrl)),
		TemplateCardJump("企业微信官网", TemplateCardJumpUrl("https://work.weixin.qq.com")),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_weComBot_PushTemplateCardNewsNotice(t *testing.T) {
	setup()
	SetDebug(true)

	// "media_id":"38BHOWH1SHSCZImMcuPmG2TuJSpYikh0AxznKJYSUJAJaFJvDeRu60NTAuj_IKLoR"
	media := Media{ID: "38BHOWH1SHSCZImMcuPmG2TuJSpYikh0AxznKJYSUJAJaFJvDeRu60NTAuj_IKLoR"}

	rawUrl := "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E5%9B%BE%E6%96%87%E5%B1%95%E7%A4%BA%E6%A8%A1%E7%89%88%E5%8D%A1%E7%89%87"
	imgUrl := "https://wework.qpic.cn/wwpic/354393_4zpkKXd7SrGMvfg_1629280616/0"

	err := bot.PushTemplateCardNewsNotice(
		TemplateCardMainTitle("一级标题", "标题辅助信息"), TemplateCardImage(imgUrl), TemplateCardActionUrl(rawUrl),
		TemplateCardSource("https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0", "企业微信"),
		TemplateCardVerticalContent("卡片二级标题", "二级普通文本"),
		TemplateCardHorizontalContent("二级标题(text)", TemplateCardHorizontalContentText("二级文本")),
		TemplateCardHorizontalContent("二级标题(url)", TemplateCardHorizontalContentUrl(rawUrl, "api地址")),
		TemplateCardHorizontalContent("二级标题(media)", TemplateCardHorizontalContentMedia("IMG_5246.jpg", media)),
		TemplateCardJump("跳转指引", TemplateCardJumpUrl(rawUrl)),
		TemplateCardJump("企业微信官网", TemplateCardJumpUrl("https://work.weixin.qq.com")),
	)
	if err != nil {
		t.Fatal(err)
	}
}
