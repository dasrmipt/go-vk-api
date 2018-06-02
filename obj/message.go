package obj

import (
	"strconv"
	"strings"
	"fmt"
)

type Message struct {
	ID          int64             `json:"id"`
	UserId      int64             `json:"user_id"`
	FromId      int64             `json:"from_id"`
	Date        int64             `json:"date"`
	ReadState   byte              `json:"read_state"`
	Out         byte              `json:"out"`
	Title       string            `json:"title"`
	Body        string            `json:"body"`
	Attachments map[string]string `json:"attachments"`
	FwdMessages []Message         `json:"fwd_messages"`
	Emoji       byte              `json:"emoji"`
	Important   byte              `json:"important"`
	Deleted     byte              `json:"deleted"`
	Random_id   int64             `json:"random_id"`
}

type MessageToSend struct {
	UserId      int64             `json:"user_id"`
	Random_id   int64             `json:"random_id"`
	PeerId      int64             `json:"peer_id"`
	Domain      string            `json:"domain"`
	ChatId      int64             `json:"chat_id"`
	UserIds     []int64           `json:"user_ids"`
	Message     string            `json:"message"`
	Attachments map[string]string `json:"attachments"`
	FwdMessages []int64           `json:"forward_messages"`
	StickerId   int64             `json:"sticker_id"`
}

func intToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func arrayIntToStr(num []int64) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(num)), ","), "[]")
}

func (msg *MessageToSend) ToRequestParams() (params map[string]string) {
	params = map[string]string{
		"message": msg.Message,
	}
	if msg.PeerId != 0 {
		params["peer_id"] = intToStr(msg.PeerId)
	} else if msg.UserId != 0 {
		params["user_id"] = intToStr(msg.UserId)
	} else if msg.Domain != "" {
		params["domain"] = msg.Domain
	} else if msg.ChatId != 0 {
		params["chat_id"] = intToStr(msg.ChatId)
	} else if msg.UserIds != nil {
		params["user_ids"] = arrayIntToStr(msg.UserIds)
	}

	if msg.FwdMessages != nil {
		params["forward_messages"] = arrayIntToStr(msg.FwdMessages)
	}
	return
}
