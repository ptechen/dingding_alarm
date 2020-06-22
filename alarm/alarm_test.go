package alarm

import "testing"

func TestAlarm_SendAlarm(t *testing.T) {
	p := &Alarm{
		Url: "https://oapi.dingtalk.com/robot/send?access_token=d7cae6bf6a69bc327b8c69696cb0c6ea08d6cef645891e41aee7715651961f8e",
		Message: &Message{
			Msgtype: "text",
			Text:    &Text{
				Content: "etl:test",
			},
		},
	}
	err := p.SendAlarm()
	if err != nil {
		t.Error(err)
	}
}