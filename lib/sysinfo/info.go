package sysinfo

import (
	"bytes"
	"fmt"
	"net"

	"github.com/StackExchange/wmi"
)

type Win32_Processor struct {
	Name string
}

type AntiVirusProduct struct {
	Name string
}

type Win32_OperatingSystem struct {
	Caption        string
	Version        string
	OSArchitecture string
}

type Win32_ProcessStartup struct {
	CreateFlags          uint32
	EnvironmentVariables []string
	ErrorMode            uint16
	FillAttribute        uint32
	PriorityClass        uint32
	ShowWindow           uint16
	Title                string
	WinstationDesktop    string
	X                    uint32
	XCountChars          uint32
	XSize                uint32
	Y                    uint32
	YCountChars          uint32
	YSize                uint32
}

func GetMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func GetCpuName() string {
	var dst []Win32_Processor
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return "Err"
	}
	for _, v := range dst {
		return v.Name
	}
	return "NaN"
}

func GetAvName() string {
	var dst []AntiVirusProduct
	q := wmi.CreateQuery(&dst, "")
	err := wmi.QueryNamespace(q, &dst, "root\\SecurityCenter2")
	if err != nil {
		return err.Error()
	}
	for _, v := range dst {
		return v.Name
	}
	return "NaN"
}

func GetOsName() string {
	var dst []Win32_OperatingSystem
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return "Err"
	}
	for _, v := range dst {
		return fmt.Sprintf("%v %v %v", v.Caption, v.Version, v.OSArchitecture)
	}
	return "NaN"
}

func RunWmiProcess(processName string) string {
	var dst []Win32_ProcessStartup
	q := wmi.CreateQuery(&dst, processName)
	err := wmi.QueryNamespace(q, &dst, "root\\cimv2")
	if err != nil {
		return err.Error()
	}
	for _, v := range dst {
		return v.Title
	}
	return "NaN"
}
