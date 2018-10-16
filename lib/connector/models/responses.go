package models

type Response_Heartbeat struct {
	NeedToWork bool     `json:"e"`
	Tasks      []string `json:"t"`
}

type Response_Task struct {
	Action int
	Args   []string
}
