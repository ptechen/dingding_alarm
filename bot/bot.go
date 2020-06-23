package bot

import (
	"encoding/json"
	"github.com/kirinlabs/HttpRequest"
)

//go:generate kratos tool protoc --grpc bot.proto

func (p *Bot) SendAlarm() (err error) {
	if p.IsAtAll && p.Message.At != nil {
		p.Message.At.IsAtAll = true
	}
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

type ParamsBot struct {
	BotName map[string]string `json:"bot_name"`
	GroupAt map[string]*At    `json:"group_at"`
}

func (p *Bot) AlarmConfig(params *ParamsBot) {
	p.Url = params.BotName[p.BotName]
	p.Message.At, _ = params.GroupAt[p.GroupAt]
}

func (p *Bot) Send(params *ParamsBot) (err error) {
	p.AlarmConfig(params)
	err = p.SendAlarm()
	return err
}