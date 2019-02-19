package goToolEnvironment

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

//获取操作系统类型
func GetOsType() string {
	return runtime.GOOS
}

//获取当前操作系统名称
func GetOsName() string {
	return runtime.GOOS
}

//获取当前操作系统版本号（Win10版本号获取错误）
func GetOsVer() (string, error) {
	version, err := syscall.GetVersion()
	if err != nil {
		return "", err
	} else {
		ver := fmt.Sprintf("%d.%d.%d", byte(version), uint8(version>>8), version>>16)
		return ver, nil
	}
}

//获取主机名称
func GetHostName() (string, error) {
	return os.Hostname()
}
