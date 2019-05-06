package goToolEnvironment

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//======================================================================================================================
//获取公网IP
type internetIPTool struct {
	address string
	data    internetIPInfo
}

type internetIPInfo interface {
	GetIp() string
}

type internetIPInfo1 struct {
	Ip string `json:"ip"`
}

func (ip *internetIPInfo1) GetIp() string {
	return ip.Ip
}

type internetIPInfo2 struct {
	Ip string `json:"ip_addr"`
}

func (ip *internetIPInfo2) GetIp() string {
	return ip.Ip
}

type internetIPInfo3 struct {
	Ip string `json:"IP"`
}

func (ip *internetIPInfo3) GetIp() string {
	return ip.Ip
}

func GetInternetAddr() (ip string, err error) {
	list := getInternetIPTool()
	for _, t := range list {
		ip, err = getInternetAddr(t)
		if err == nil {
			return
		}
	}
	return
}

func getInternetIPTool() []*internetIPTool {
	list := make([]*internetIPTool, 0)
	list = append(list, &internetIPTool{
		address: "https://ipconfig.io/json",
		data:    &internetIPInfo1{},
	})
	list = append(list, &internetIPTool{
		address: "https://ifconfig.me/all.json",
		data:    &internetIPInfo2{},
	})
	list = append(list, &internetIPTool{
		address: "https://ifconfig.minidump.info/all.json",
		data:    &internetIPInfo3{},
	})
	return list
}

func getInternetAddr(tool *internetIPTool) (string, error) {
	//https://www.baidu.com/s?wd=IP
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
	resp, err := client.Get(tool.address)
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
	err = json.Unmarshal(body, tool.data)
	if err != nil {
		return "", err
	}
	return tool.data.GetIp(), nil
}

//======================================================================================================================
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
