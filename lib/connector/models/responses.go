package models

type Response_Heartbeat struct {
	NeedToWork bool
	Tasks      []string
}

type Response_Task struct {
	Action int
	Args   []string
}
