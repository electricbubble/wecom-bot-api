package wecom_bot_api

type templateCardMsg struct {
	// 消息类型, 固定为 `template_card`
	MsgType      string       `json:"msgtype"`
	TemplateCard templateCard `json:"template_card"`
}

func newTemplateCardMsg(templateCard *templateCard) *templateCardMsg {
	return &templateCardMsg{
		MsgType:      "template_card",
		TemplateCard: *templateCard,
	}
}

type templateCard struct {
	// 模版卡片类型
	//  文本通知模版卡片类型: `text_notice`
	//  图文展示模版卡片类型: `news_notice`
	CardType        string                       `json:"card_type"`
	Source          *templateCardSource          `json:"source,omitempty"`
	MainTitle       templateCardMainTitle        `json:"main_title"`
	EmphasisContent *templateCardEmphasisContent `json:"emphasis_content,omitempty"`

	CardImage           *templateCardImage            `json:"card_image,omitempty"`
	VerticalContentList []templateCardVerticalContent `json:"vertical_content_list,omitempty"`

	// 二级普通文本
	//  建议不超过112个字
	SubTitleText string `json:"sub_title_text,omitempty"`

	// 二级标题+文本列表，该字段可为空数组
	//
	// 但有数据的话需确认对应字段是否必填
	//
	// 列表长度不超过6
	HorizontalContentList []templateCardHorizontalContent `json:"horizontal_content_list"`

	// 跳转指引样式的列表
	//
	// 该字段可为空数组
	//
	// 但有数据的话需确认对应字段是否必填
	//
	// 列表长度不超过3
	JumpList []templateCardJump `json:"jump_list"`

	// 整体卡片的点击跳转事件
	//
	// text_notice 模版卡片中该字段为必填项
	CardAction templateCardAction `json:"card_action"`
}

func newTemplateCardText(mainTitle TemplateCardMainTitleOption, cardAction TemplateCardAction, opts ...TemplateCardOption) *templateCard {
	var _mainTitle templateCardMainTitle
	mainTitle(&_mainTitle)
	var _cardAction templateCardAction
	cardAction(&_cardAction)
	tplCard := templateCard{
		CardType:              "text_notice",
		Source:                nil,
		MainTitle:             _mainTitle,
		EmphasisContent:       nil,
		VerticalContentList:   make([]templateCardVerticalContent, 0, 4),
		SubTitleText:          "",
		HorizontalContentList: make([]templateCardHorizontalContent, 0, 6),
		JumpList:              make([]templateCardJump, 0, 3),
		CardAction:            _cardAction,
	}
	for _, opt := range opts {
		opt(&tplCard)
	}
	return &tplCard
}

func newTemplateCardNews(mainTitle TemplateCardMainTitleOption, cardImage TemplateCardImageOption, cardAction TemplateCardAction, opts ...TemplateCardOption) *templateCard {
	var (
		_mainTitle  templateCardMainTitle
		_cardAction templateCardAction
		_cardImage  templateCardImage
	)
	mainTitle(&_mainTitle)
	cardAction(&_cardAction)
	cardImage(&_cardImage)
	tplCard := templateCard{
		CardType:              "news_notice",
		Source:                nil,
		MainTitle:             _mainTitle,
		EmphasisContent:       nil,
		CardImage:             &_cardImage,
		VerticalContentList:   make([]templateCardVerticalContent, 0, 4),
		SubTitleText:          "",
		HorizontalContentList: make([]templateCardHorizontalContent, 0, 6),
		JumpList:              make([]templateCardJump, 0, 3),
		CardAction:            _cardAction,
	}
	for _, opt := range opts {
		opt(&tplCard)
	}
	return &tplCard
}

type TemplateCardMainTitleOption func(mt *templateCardMainTitle)

func TemplateCardMainTitle(title string, desc string) TemplateCardMainTitleOption {
	return func(mt *templateCardMainTitle) {
		mt.Title = title
		mt.Desc = desc
	}
}

type TemplateCardOption func(tc *templateCard)

// func TemplateCardSourceIconUrl(iconUrl string) TemplateCardOption {
// 	return func(tc *templateCard) {
// 		tc.Source.IconUrl = iconUrl
// 	}
// }
//
// func TemplateCardSourceDesc(desc string) TemplateCardOption {
// 	return func(tc *templateCard) {
// 		tc.Source.Desc = desc
// 	}
// }

func TemplateCardSource(iconUrl, desc string) TemplateCardOption {
	return func(tc *templateCard) {
		if tc.Source == nil {
			tc.Source = new(templateCardSource)
		}
		tc.Source.IconUrl = iconUrl
		tc.Source.Desc = desc
	}
}

// func TemplateCardEmphasisContentTitle(title string) TemplateCardOption {
// 	return func(tc *templateCard) {
// 		tc.EmphasisContent.Title = title
// 	}
// }
//
// func TemplateCardEmphasisContentDesc(desc string) TemplateCardOption {
// 	return func(tc *templateCard) {
// 		tc.EmphasisContent.Desc = desc
// 	}
// }

func TemplateCardEmphasisContent(title, desc string) TemplateCardOption {
	return func(tc *templateCard) {
		if tc.EmphasisContent == nil {
			tc.EmphasisContent = new(templateCardEmphasisContent)
		}
		tc.EmphasisContent.Title = title
		tc.EmphasisContent.Desc = desc
	}
}

func TemplateCardSubTitleText(subTitle string) TemplateCardOption {
	return func(tc *templateCard) {
		tc.SubTitleText = subTitle
	}
}

// TemplateCardHorizontalContent 二级标题+文本列表
//  列表长度不超过6
func TemplateCardHorizontalContent(keyName string, opt TemplateCardHorizontalContentOption) TemplateCardOption {
	return func(tc *templateCard) {
		horizontalContent := templateCardHorizontalContent{KeyName: keyName}
		if opt != nil {
			opt(&horizontalContent)
		}
		tc.HorizontalContentList = append(tc.HorizontalContentList, horizontalContent)
	}
}

type TemplateCardHorizontalContentOption func(hc *templateCardHorizontalContent)

func TemplateCardHorizontalContentText(text string) TemplateCardHorizontalContentOption {
	return func(hc *templateCardHorizontalContent) {
		hc.Type = 0
		hc.Value = text
	}
}

func TemplateCardHorizontalContentUrl(rawUrl string, text string) TemplateCardHorizontalContentOption {
	return func(hc *templateCardHorizontalContent) {
		hc.Type = 1
		hc.Url = rawUrl
		hc.Value = text
	}
}

func TemplateCardHorizontalContentMedia(filename string, media Media) TemplateCardHorizontalContentOption {
	return func(hc *templateCardHorizontalContent) {
		hc.Type = 2
		hc.Value = filename
		hc.MediaId = media.ID
	}
}

// TemplateCardJump 跳转指引样式的列表
//  列表长度不超过3
func TemplateCardJump(title string, opt TemplateCardJumpOption) TemplateCardOption {
	return func(tc *templateCard) {
		jump := templateCardJump{
			Type:  0,
			Title: title,
		}
		if opt != nil {
			opt(&jump)
		}
		tc.JumpList = append(tc.JumpList, jump)
	}
}

type TemplateCardJumpOption func(j *templateCardJump)

func TemplateCardJumpUrl(rawUrl string) TemplateCardJumpOption {
	return func(j *templateCardJump) {
		j.Type = 1
		j.Url = rawUrl
	}
}

func TemplateCardJumpApp(appID string, pagePath string) TemplateCardJumpOption {
	return func(j *templateCardJump) {
		j.Type = 2
		j.AppID = appID
		j.PagePath = pagePath
	}
}

type TemplateCardAction func(act *templateCardAction)

func TemplateCardActionUrl(rawUrl string) TemplateCardAction {
	return func(act *templateCardAction) {
		act.Type = 1
		act.Url = rawUrl
	}
}

func TemplateCardActionApp(appID string, pagePath string) TemplateCardAction {
	return func(act *templateCardAction) {
		act.Type = 2
		act.Appid = appID
		act.PagePath = pagePath
	}
}

type TemplateCardImageOption func(img *templateCardImage)

// TemplateCardImage
//  aspectRatio 图片的宽高比
//   宽高比要小于2.25
//   大于1.3
//   不填该参数默认1.3
func TemplateCardImage(rawUrl string, aspectRatio ...float64) TemplateCardImageOption {
	return func(img *templateCardImage) {
		img.Url = rawUrl
		if len(aspectRatio) != 0 {
			img.AspectRatio = aspectRatio[0]
		}
	}
}

// TemplateCardVerticalContent 卡片二级垂直内容
//  列表长度不超过4
func TemplateCardVerticalContent(title, desc string) TemplateCardOption {
	return func(tc *templateCard) {
		tc.VerticalContentList = append(tc.VerticalContentList, templateCardVerticalContent{
			Title: title,
			Desc:  desc,
		})
	}
}

type (
	// templateCardEmphasisContent 关键数据样式
	templateCardEmphasisContent struct {
		Title string `json:"title,omitempty"` // 关键数据样式的数据内容，建议不超过10个字
		Desc  string `json:"desc,omitempty"`  // 关键数据样式的数据描述内容，建议不超过15个字
	}

	// templateCardImage 图片样式
	templateCardImage struct {
		Url         string  `json:"url"`                    // 图片的 url
		AspectRatio float64 `json:"aspect_ratio,omitempty"` // 图片的宽高比，宽高比要小于2.25，大于1.3，不填该参数默认1.3
	}

	// templateCardVerticalContent 卡片二级垂直内容，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过4
	templateCardVerticalContent struct {
		Title string `json:"title"`          // 卡片二级标题，建议不超过26个字
		Desc  string `json:"desc,omitempty"` // 二级普通文本，建议不超过112个字
	}
)

type (
	// templateCardSource 卡片来源样式信息，不需要来源样式可不填写
	templateCardSource struct {
		IconUrl string `json:"icon_url,omitempty"` // 来源图片的url
		Desc    string `json:"desc,omitempty"`     // 来源图片的描述，建议不超过13个字
	}

	// templateCardMainTitle 模版卡片的主要内容，包括一级标题和标题辅助信息
	templateCardMainTitle struct {
		Title string `json:"title"`          // 一级标题，建议不超过26个字
		Desc  string `json:"desc,omitempty"` // 标题辅助信息，建议不超过30个字
	}

	// templateCardHorizontalContent 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	templateCardHorizontalContent struct {
		Type    int    `json:"type,omitempty"`     // 链接类型，0或不填代表是普通文本，1 代表跳转url，2 代表下载附件
		KeyName string `json:"keyname"`            // 二级标题，建议不超过5个字
		Value   string `json:"value,omitempty"`    // 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过26个字
		Url     string `json:"url,omitempty"`      // 链接跳转的url，horizontal_content_list.type是1时必填
		MediaId string `json:"media_id,omitempty"` // 附件的media_id，horizontal_content_list.type是2时必填
	}

	// templateCardJump 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
	templateCardJump struct {
		Type     int    `json:"type,omitempty"`     // 跳转链接类型，0 或不填代表不是链接，1 代表跳转 url，2 代表跳转小程序
		Title    string `json:"title"`              // 跳转链接样式的文案内容，建议不超过 13个字
		Url      string `json:"url,omitempty"`      // 跳转链接的 url，jump_list.type 是 1 时必填
		AppID    string `json:"appid,omitempty"`    // 跳转链接的小程序的 appid，jump_list.type 是 2 时必填
		PagePath string `json:"pagepath,omitempty"` // 跳转链接的小程序的 pagepath，jump_list.type 是 2 时选填
	}

	// templateCardAction 整体卡片的点击跳转事件，text_notice 模版卡片中该字段为必填项
	templateCardAction struct {
		Type     int    `json:"type"`               // 卡片跳转类型，0 或不填代表不是链接，1 代表跳转url，2 代表打开小程序。text_notice模版卡片中该字段取值范围为[1,2]
		Url      string `json:"url,omitempty"`      // 跳转事件的 url，card_action.type 是 1 时必填
		Appid    string `json:"appid,omitempty"`    // 跳转事件的小程序的 appid，card_action.type 是 2 时必填
		PagePath string `json:"pagepath,omitempty"` // 跳转事件的小程序的 pagepath，card_action.type 是 2 时选填
	}
)
