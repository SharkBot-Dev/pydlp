package tools

import (
	"fmt"
	"runtime"
)

func GetDeviceInfo() (string, string) {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Device: Windows")
		return "_x86.exe", ".exe"
	case "linux":
		fmt.Println("Device: Linux")
		return "_linux", ""
	default:
		return "null", "null"
	}
}

func GetDeviceSuff() string {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Device: Windows")
		return ".exe"
	case "linux":
		fmt.Println("Device: Linux")
		return ""
	default:
		return "null"
	}
}
