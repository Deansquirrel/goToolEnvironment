package goToolEnvironment

import (
	"errors"
	"github.com/Deansquirrel/goToolCommon"
	"os/exec"
	"runtime"
	"strings"
)

func GetClientId(clientType string) (string, error) {
	var biosFlag, diskFlag, cpuFlag bool
	var err error
	biosSn, err := GetBIOSSerialNumber()
	if err != nil {
		biosFlag = true
	}
	diskSn, err := GetDiskDriverSerialNumber()
	if err != nil {
		diskFlag = true
	}
	cpuId, err := GetCPUPorcessorID()
	if err != nil {
		cpuFlag = true
	}
	currPath, err := goToolCommon.GetCurrPath()
	if err != nil {
		return "", err
	}
	if biosFlag && diskFlag && cpuFlag {
		return "", errors.New("key factor err")
	}
	strKey := ""
	strKey = strKey + clientType
	if !biosFlag {
		strKey = strKey + biosSn
	}
	if !diskFlag {
		strKey = strKey + diskSn
	}
	if !cpuFlag {
		strKey = strKey + cpuId
	}
	strKey = strKey + currPath
	return strings.ToUpper(goToolCommon.Md5([]byte(strKey))), nil
}

func GetPhysicalId() (string, error) {
	biosSn, err := GetBIOSSerialNumber()
	if err != nil {
		return "", err
	}
	diskSn, err := GetDiskDriverSerialNumber()
	if err != nil {
		return "", err
	}
	cpuId, err := GetCPUPorcessorID()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(goToolCommon.Md5([]byte(biosSn + diskSn + cpuId))), nil
}

//获取硬盘SerialNumber
func GetDiskDriverSerialNumber() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return diskDriverSerialNumberOnWindows()
	case "linux":
		return "", errors.New("unsupported os")
	case "darwin":
		return "", errors.New("unsupported os")
	default:
		return "", errors.New("unknown os")
	}
}

func diskDriverSerialNumberOnWindows() (string, error) {
	cmd := exec.Command("CMD", "/C", "WMIC DISKDRIVE GET SERIALNUMBER")
	serialNo, err := cmd.Output()
	if err != nil {
		return "", err
	}
	l := strings.Split(string(serialNo), "\n")
	if len(l) >= 2 {
		return l[1], nil
	} else {
		return "", errors.New("return split length less 2")
	}
}

//获取硬盘SerialNumber
func GetBIOSSerialNumber() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return biosSerialNumberOnWindows()
	case "linux":
		return "", errors.New("unsupported os")
	case "darwin":
		return "", errors.New("unsupported os")
	default:
		return "", errors.New("unknown os")
	}
}

func biosSerialNumberOnWindows() (string, error) {
	cmd := exec.Command("CMD", "/C", "WMIC BIOS GET SERIALNUMBER")
	serialNo, err := cmd.Output()
	if err != nil {
		return "", err
	}
	l := strings.Split(string(serialNo), "\n")
	if len(l) >= 2 {
		return l[1], nil
	} else {
		return "", errors.New("return split length less 2")
	}
}

//获取CPU PorcessorID
func GetCPUPorcessorID() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return cpuPorcessorIDOnWindows()
	case "linux":
		return "", errors.New("unsupported os")
	case "darwin":
		return "", errors.New("unsupported os")
	default:
		return "", errors.New("unknown os")
	}
}

func cpuPorcessorIDOnWindows() (string, error) {
	cmd := exec.Command("CMD", "/C", "WMIC CPU GET ProcessorID")
	serialNo, err := cmd.Output()
	if err != nil {
		return "", err
	}
	l := strings.Split(string(serialNo), "\n")
	if len(l) >= 2 {
		return l[1], nil
	} else {
		return "", errors.New("return split length less 2")
	}
}
