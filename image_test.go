package wecom_bot_api

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func Test_newImageMsg(t *testing.T) {
	userHomeDir, _ := os.UserHomeDir()
	filename := path.Join(userHomeDir, "Pictures", "IMG_5246.jpg")

	readFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	msg := newImageMsg(readFile)

	marshal, err := json.Marshal(&msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
}
