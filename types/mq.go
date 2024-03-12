package types

// MessageDuobaoPayload 消息夺宝负载
type MessageDuobaoPayload struct {
	Type   int `json:"type"`
	UserId int `json:"userId"`
}
