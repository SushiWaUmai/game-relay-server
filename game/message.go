package game

type Message struct {
	MsgType string `json:"type"`
	Author  uint   `json:"author"`
	Data    any    `json:"data"`
	Targets []uint `json:"targets"`
}
