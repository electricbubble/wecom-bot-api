package wecom_bot_api

type newsMsg struct {
	// 消息类型, 固定为 `news`
	MsgType string `json:"msgtype"`

	// 图文消息, 一个图文消息支持1到8条图文
	News newsData `json:"news"`
}

type newsData struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	// **必填** 标题, 不超过 128 个字节, 超过会自动截断
	Title string `json:"title"`

	// **选填** 描述, 不超过 512 个字节, 超过会自动截断
	Description string `json:"description,omitempty"`

	// **必填** 点击后跳转的链接
	Url string `json:"url"`

	// **选填** 图文消息的图片链接, 支持 `JPG`、`PNG` 格式, 较好的效果为大图 `1068*455`, 小图 `150*150`
	PicUrl string `json:"picurl,omitempty"`
}

func newNewsMsg(art Article, articles ...Article) newsMsg {
	articles = append([]Article{art}, articles...)
	msg := newsMsg{
		MsgType: "news",
		News:    newsData{Articles: articles},
	}

	return msg
}

func NewArticle(title, Url string, opts ...ArticleOption) Article {
	article := Article{
		Title: title,
		Url:   Url,
	}

	for _, opt := range opts {
		opt(&article)
	}

	return article
}

type ArticleOption func(d *Article)

func ArticleDescription(desc string) ArticleOption {
	return func(d *Article) {
		d.Description = desc
	}
}
func ArticlePicUrl(picUrl string) ArticleOption {
	return func(d *Article) {
		d.PicUrl = picUrl
	}
}
