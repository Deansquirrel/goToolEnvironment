package goToolEnvironment

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type internetIPInfo struct {
	Ip      string `json:"ip"`
	Country string `json:"country"`
	City    string `json:"city"`
}

//获取公网IP
func GetInternetAddr() (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://ipconfig.io/json")
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result internetIPInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.Ip, nil
}

//获取局域网IP
func GetIntranetAddr() (string, error) {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	addr := ""
	for _, ip := range ips {
		ipNet, ok := ip.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				if addr != "" {
					addr = addr + "|"
				}
				addr = addr + fmt.Sprintf("%s", ipNet.IP.To4())
			}
		}
	}
	return addr, nil
}
