package wecom_bot_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type weComBot struct {
	webhook string
	key     string
}

func NewWeComBot(key string) WeComBot {
	bot := new(weComBot)
	bot.webhook = fmt.Sprintf(BotSendUrl, key)

	bot.key = key
	return bot
}

func (b *weComBot) PushTextMessage(content string, opts ...TextMsgOption) (err error) {
	msg := newTextMsg(content, opts...)
	return b.pushMsg(msg)
}

func (b *weComBot) PushMarkdownMessage(content string) (err error) {
	msg := newMarkdownMsg(content)
	return b.pushMsg(msg)
}

func (b *weComBot) PushImageMessage(img []byte) (err error) {
	msg := newImageMsg(img)
	return b.pushMsg(msg)
}

func (b *weComBot) PushNewsMessage(art Article, articles ...Article) (err error) {
	msg := newNewsMsg(art, articles...)
	return b.pushMsg(msg)
}

func (b *weComBot) PushFileMessage(media Media) error {
	msg := newFileMsg(media.ID)
	return b.pushMsg(msg)
}

func (b *weComBot) PushTemplateCardTextNotice(mainTitle TemplateCardMainTitleOption, cardAction TemplateCardAction, opts ...TemplateCardOption) error {
	tplCardMsg := newTemplateCardMsg(newTemplateCardText(mainTitle, cardAction, opts...))
	return b.pushMsg(tplCardMsg)
}

func (b *weComBot) PushTemplateCardNewsNotice(mainTitle TemplateCardMainTitleOption, cardImage TemplateCardImageOption, cardAction TemplateCardAction, opts ...TemplateCardOption) error {
	tplCardMsg := newTemplateCardMsg(newTemplateCardNews(mainTitle, cardImage, cardAction, opts...))
	return b.pushMsg(tplCardMsg)
}

func (b *weComBot) pushMsg(msg interface{}) (err error) {
	var bsJSON []byte
	if bsJSON, err = json.Marshal(msg); err != nil {
		return err
	}
	var req *http.Request
	if req, err = newRequest(http.MethodPost, b.webhook, bsJSON); err != nil {
		return err
	}
	_, err = executeHTTP(req)
	return
}

func (b *weComBot) UploadFile(filename string) (media Media, err error) {
	var req *http.Request
	if req, err = newUploadRequest(http.MethodPost, fmt.Sprintf(UploadMediaUrl, b.key), filename); err != nil {
		return Media{}, err
	}
	var rawResp []byte = nil
	if rawResp, err = executeHTTP(req); err != nil {
		return Media{}, err
	}

	var reply = new(struct {
		ErrCode   int    `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		Type      string `json:"type"`
		MediaId   string `json:"media_id"`
		CreatedAt string `json:"created_at"`
	})
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return Media{}, fmt.Errorf("unknown response: %w\nraw response: %s", err, rawResp)
	}
	media = Media{ID: reply.MediaId}
	return
}
