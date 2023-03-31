package game

type Message struct {
	MsgType string `json:"type"`
	Data    any    `json:"data"`
	Targets []int  `json:"targets"`
}
