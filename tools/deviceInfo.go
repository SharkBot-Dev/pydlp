package tools

import (
	"fmt"
	"runtime"
)

func PrintDeviceString() string {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Device: Windows")
		return "Windows"
	case "linux":
		fmt.Println("Device: Linux")
		return "Linux"
	default:
		return "null"
	}
}

func GetDeviceInfo() (string, string) {
	switch runtime.GOOS {
	case "windows":
		return "_x86.exe", ".exe"
	case "linux":
		return "_linux", ""
	default:
		return "null", "null"
	}
}

func GetDeviceSuff() string {
	switch runtime.GOOS {
	case "windows":
		return ".exe"
	case "linux":
		return ""
	default:
		return "null"
	}
}
