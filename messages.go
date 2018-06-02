package vk

import (
	"encoding/json"
	"github.com/dasrmipt/go-vk-api/obj"
)

// Messages https://new.vk.com/dev/messages
type Messages struct {
	client *VK
}

// Send https://new.vk.com/dev/messages.send
func (messages *Messages) Send(params RequestParams) (int64, error) {
	resp, err := messages.client.CallMethod("messages.send", params)

	if err != nil {
		return 0, err
	}

	type JSONBody struct {
		MessageID int64 `json:"response"`
	}

	var body JSONBody

	if err := json.Unmarshal(resp, &body); err != nil {
		return 0, err
	}

	return body.MessageID, nil
}

func (messages *Messages) GetChat(params RequestParams) (*obj.Chat, error) {
	resp, err := messages.client.CallMethod("messages.getChat", params)

	if err != nil {
		return nil, err
	}

	type JSONBody struct {
		Chat obj.Chat `json:"response"`
	}

	var body JSONBody

	if err := json.Unmarshal(resp, &body); err != nil {
		return nil, err
	}

	return &body.Chat, nil
}