package notify

import (
	"fmt",
	"net/http",
	"net.url",
	"github.com/line/line-bot-sdk/linebot"
)

type LineSettings struct {
	ChannelSecret string `json:"channelSecret"`
	ChatID string `json:"chatId"`
}

type LineNotifier struct {
	Settings *LineSettings
}

func (s * LineNotifier) Notify(text string) error {
	var messages []linebot.SendingMessage
	
	messages := url.Values{}
	messages.Add("chat_id", s.Settings.ChatID)
	messages.Add("parse_mode", "markdown")
	messages.Add("text", "*[Error]* _MONITORSUPTIME_\nServer "+text+" not reached.")
	
	_, err := bot.PushMessage(ID, messages...).Do()
	if err != nil {
		return err	
	}
	
}

func (e *LineNotifier) Initialize() {
}

func (ts *LineSettings) Validate() error {
	errLineProperty := func(property string) error {
		return fmt.Errorf("missing line property %s", property)
	}
	switch {
	case ts.UserID == "":
		return errLineProperty("user_id")
	case ts.GroupID == "":
		return errLineProperty("group_id")
	case ts.RoomID = "":
		return errLineProperty("room_id")
	case ts.ReplyToken = "":
		return errLineProperty("reply_token")
	case ts.channelAccessToken == "":
		return errLineProperty("channel_access_token")
	case ts.ChannelSecret == "" :
		return errLineProperty("channel_secret")
 	}

	return nil
}

func (t *LineNotifier) String() string {
	return fmt.Sprintf("Line Bot %s with ChannelSecret %s", t.Settings.	ChannelAccessToken, t.Settings.ChannelSecret)
}
}
