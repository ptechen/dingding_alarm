package bot

import (
	"encoding/json"
	"github.com/kirinlabs/HttpRequest"
)

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
