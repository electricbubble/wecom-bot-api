package wecom_bot_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strings"
	"time"
)

// BotSendUrl 企业微信群机器人 webhook
var BotSendUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"

// UploadMediaUrl 企业微信上传文件接口 `url`, `type` 固定传 `file`
var UploadMediaUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=%s&type=file"

type WeComBot interface {
	PushTextMessage(content string, opts ...TextMsgOption) error
	PushMarkdownMessage(content string) error
	PushImageMessage(img []byte) error
	PushNewsMessage(art Article, articles ...Article) error
	PushFileMessage(media Media) error
	UploadFile(filename string) (media Media, err error)
}

var HTTPClient = http.DefaultClient

func newRequest(method string, rawUrl string, rawBody []byte) (request *http.Request, err error) {
	debugLog(fmt.Sprintf("--> %s %s\n%s", method, rawUrl, rawBody))

	if request, err = http.NewRequest(method, rawUrl, bytes.NewBuffer(rawBody)); err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	return
}

func newUploadRequest(method string, rawUrl string, filename string) (request *http.Request, err error) {
	bodyBuffer := bytes.NewBufferString("")
	writer := multipart.NewWriter(bodyBuffer)

	mediaFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() { _ = mediaFile.Close() }()

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			"media", escapeQuotes(path.Base(filename))))
	h.Set("Content-Type", "application/octet-stream")
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}
	_, _ = io.Copy(part, mediaFile)

	_ = writer.Close()

	debugLog(fmt.Sprintf("--> %s %s [MEDIA_DATA: Don't display]", method, rawUrl))

	if request, err = http.NewRequest(method, rawUrl, bodyBuffer); err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	return
}

func executeHTTP(req *http.Request) (rawResp []byte, err error) {
	start := time.Now()
	var resp *http.Response
	if resp, err = HTTPClient.Do(req); err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	rawResp, err = ioutil.ReadAll(resp.Body)
	debugLog(fmt.Sprintf("<-- %s %s %d %s\n%s\n", req.Method, req.URL.String(), resp.StatusCode, time.Since(start), rawResp))
	if err != nil {
		return nil, err
	}

	var reply = new(struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	})
	if err = json.Unmarshal(rawResp, reply); err != nil {
		return nil, fmt.Errorf("unknown response: %w\nraw response: %s", err, rawResp)
	}
	if reply.ErrMsg != "ok" {
		return nil, fmt.Errorf("unknown response: %s", rawResp)
	}

	return
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

var debugFlag = false

func SetDebug(debug bool) {
	debugFlag = debug
}

func debugLog(msg string) {
	if !debugFlag {
		return
	}
	log.Println("[DEBUG-WeCom-Bot-API] " + msg)
}
