# WeCom-Bot-API

企业微信-群机器人-API

## Installation

```shell
go get github.com/electricbubble/wecom-bot-api
```

## Usage

#### 纯文本消息

```go
package main

import (
	botApi "github.com/electricbubble/wecom-bot-api"
	"os"
)

func main() {
	botKey := os.Getenv("WeCom_Bot_Key")
	phoneNumber := os.Getenv("Phone_Number")
	userid := os.Getenv("Userid")

	bot := botApi.NewWeComBot(botKey)

	// 仅发送文本内容
	_ = bot.PushTextMessage("hi")

	// 通过群成员 `手机号码` 进行 `@` 提醒
	_ = bot.PushTextMessage("hi again", botApi.MentionByMobile(phoneNumber))

	// 通过群成员 `userid` 进行 `@` 提醒
	_ = bot.PushTextMessage("hi again", botApi.MentionByUserid(userid))

	// @全部成员
	_ = bot.PushTextMessage("hi again",
		botApi.MentionAllByMobile(),
		// botApi.MentionAllByUserid(),
	)
}

```

#### Markdown 消息

```go
package main

import (
	"bytes"
	botApi "github.com/electricbubble/wecom-bot-api"
	"github.com/electricbubble/wecom-bot-api/md"
	"os"
)

func main() {
	botKey := os.Getenv("WeCom_Bot_Key")
	userid := os.Getenv("Userid")

	bot := botApi.NewWeComBot(botKey)

	content := bytes.NewBufferString(md.Heading(1, "H1"))
	content.WriteString("实时新增用户反馈" + md.WarningText("132例") + "，请相关同事注意。\n")
	content.WriteString(md.QuoteText("类型:" + md.CommentText("用户反馈")))
	content.WriteString(md.QuoteText("普通用户反馈:" + md.CommentText("117例")))
	content.WriteString(md.QuoteText("VIP用户反馈:" + md.CommentText("15例")))
	// 👆效果等同于👇
	/*
		# H1
		实时新增用户反馈 <font color="warning">132例</font>，请相关同事注意。\n
		> 类型:<font color="comment">用户反馈</font>
		> 普通用户反馈:<font color="comment">117例</font>
		> VIP用户反馈:<font color="comment">15例</font>
	*/

	// 仅发送 `markdown` 格式的文本
	_ = bot.PushMarkdownMessage(content.String())

	// 通过群成员 `userid` 进行 `@` 提醒
	_ = bot.PushMarkdownMessage(
		md.Heading(2, "H2") + md.Bold("hi") + "\n" + "> again\n" +
			md.MentionByUserid(userid),
	)
}

```

#### 图片消息

```go
package main

import (
	botApi "github.com/electricbubble/wecom-bot-api"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	botKey := os.Getenv("WeCom_Bot_Key")
	bot := botApi.NewWeComBot(botKey)

	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")
	readFile, _ := ioutil.ReadFile(filename)

	// 发送 图片消息
	_ = bot.PushImageMessage(readFile)
}

```

#### 图文消息

```go
package main

import (
	botApi "github.com/electricbubble/wecom-bot-api"
	"os"
)

func main() {
	botKey := os.Getenv("WeCom_Bot_Key")
	bot := botApi.NewWeComBot(botKey)

	article := botApi.NewArticle("中秋节礼品领取", "www.qq.com",
		botApi.ArticleDescription("今年中秋节公司有豪礼相送"),
		botApi.ArticlePicUrl("http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png"),
	)
	article2 := botApi.NewArticle("图文标题2", "www.qq.com",
		botApi.ArticleDescription("图文描述2"),
	)
	article3 := botApi.NewArticle("图文标题3", "www.qq.com",
		botApi.ArticleDescription("图文描述3"),
	)
	_, _ = article2, article3

	// 发送 `1条图文` 消息
	_ = bot.PushNewsMessage(article)

	// 发送 `多条图文` 消息 (一个图文消息支持 `1~8条` 图文)
	_ = bot.PushNewsMessage(article, article2, article3)
}

```

#### 文件消息

```go
package main

import (
	botApi "github.com/electricbubble/wecom-bot-api"
	"os"
	"path"
)

func main() {
	botKey := os.Getenv("WeCom_Bot_Key")
	bot := botApi.NewWeComBot(botKey)

	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")

	// 必须先通过企业微信上传文件接口, 获取 `media_id` (仅三天内有效)
	// https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0%E6%8E%A5%E5%8F%A3
	media, _ := bot.UploadFile(filename)
	// 发送 文件消息
	_ = bot.PushFileMessage(media)
}

```