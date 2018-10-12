package models

type Request_Register struct {
	Id          string `json:"i"`
	Version     int    `json:"v"`
	Workstation string `json:"w"`
	Arch        string `json:"a"`
	Cpu         string `json:"c"`
	Cores       int    `json:"t"`
	Av          string `json:"p"`
	Utility     int    `json:"u"`
	Os          string `json:"o"`
}

type Request_Heartbeat struct {
	Utilization int    `json:"u"`
	Working     bool   `json:"s"`
	Id          string `json:"i"`
}

type Request_Task struct {
	Id string
}
