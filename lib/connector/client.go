package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/monsun69/idle_service/lib/connector/models"
)

type Instance struct {
	http      http.Client
	baseUrl   string
	UserAgent string
	Auth      http.Cookie
}

func (i *Instance) Init(url string) {
	i.baseUrl = url
	var TLSCc *tls.Config
	if strings.HasPrefix(i.baseUrl, "http://") {
		TLSCc = nil
	} else {
		TLSCc = &tls.Config{InsecureSkipVerify: false}
	}

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   10,
		MaxIdleConns:          10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
		TLSClientConfig:       TLSCc,
	}
	defer tr.CloseIdleConnections()
	i.http = http.Client{
		Transport: tr,
		Timeout:   60 * time.Second,
	}
}

func (i *Instance) Register(data models.Request_Register) (bool, error) {
	// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/join", i.baseUrl), bytes.NewBuffer(jsonValue))
	if err != nil {
		return false, err
	}
	req.Header.Set("User-Agent", i.UserAgent)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := i.http.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != 200 {
		return false, nil
	}
	if len(resp.Cookies()) < 1 {
		return false, nil
	}
	i.Auth = *resp.Cookies()[0]
	defer resp.Body.Close()
	return true, nil
}

func (i *Instance) Heartbeat(data models.Request_Heartbeat) (result models.Response_Heartbeat, err error) {
	// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/heartbeat", i.baseUrl), bytes.NewBuffer(jsonValue))
	if err != nil {
		return result, err
	}
	req.Header.Set("User-Agent", i.UserAgent)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.AddCookie(&i.Auth)
	resp, err := i.http.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, fmt.Errorf("Unauth")
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	return result, nil
}

func (i *Instance) Task(data models.Request_Task) (result models.Response_Task, err error) {
	// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/task", i.baseUrl), bytes.NewBuffer(jsonValue))
	if err != nil {
		return result, err
	}
	req.Header.Set("User-Agent", i.UserAgent)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.AddCookie(&i.Auth)
	resp, err := i.http.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, fmt.Errorf("Unauth")
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	return result, nil
}
