package alarm

import (
	"encoding/json"
	"github.com/kirinlabs/HttpRequest"
)

type Alarm struct {
	Url     string
	Secret  string
	Message *Message
}

type Message struct {
	Msgtype  string    `json:"msgtype"`
	Text     *Text     `json:"text"`
	Link     *Link     `json:"link"`
	Markdown *Markdown `json:"markdown"`
}

type Text struct {
	Content string `json:"content"`
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"pic_url"`
	MessageUrl string `json:"messageUrl"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (p *Alarm) SendAlarm() (err error) {
	data := make(map[string]interface{})
	dataBytes, _ := json.Marshal(p.Message)
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return err
	}
	req := HttpRequest.NewRequest()
	req = req.SetHeaders(map[string]string{"Content-Type": "application/json"})
	_, err = req.Post(p.Url, data)
	return err
}
