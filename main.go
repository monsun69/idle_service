package main

import (
	"C"
	"log"
	"os"
	"time"

	"github.com/monsun69/idle_service/winsvc"
)

//export EntryPoint
func EntryPoint() {
	f, err := os.OpenFile("C:\\idle_svc.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Panicf("Error on use log file. CODE: %v\n", err)
	}
	log.SetOutput(f)

	if err := winsvc.OurServiceInit(); err != 0 {
		log.Panicf("Error on service initialization. CODE: %v\n", err)
	}

	for {
		time.Sleep(time.Second * 30)
		log.Printf("Idle service is alive.")
	}
}

func main() {
}
