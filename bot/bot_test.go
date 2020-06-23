package bot

import "testing"

func TestAlarm_SendAlarm(t *testing.T) {
	p := &Alarm{
		Url: "https://oapi.dingtalk.com/robot/send?access_token=***",
		Message: &Message{
			Msgtype: "text",
			Text:    &Text{
				Content: "test",
			},
		},
	}
	err := p.SendAlarm()
	if err != nil {
		t.Error(err)
	}
}