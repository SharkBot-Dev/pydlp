package main

import (
	"fmt"
	"os"
	"os/exec"

	"pydlp/tools"
)

func createHelp() {
	fmt.Println("How to use: pydlp [Commands] [Options]\n\nCommands:\ninstall (update): Installs the specified version\nversion: Check or change the version.\nversions: Displays a list of versions.\nrun: Run using the current version.")
}

func main() {
	folder := "yt-dlps"

	err := os.MkdirAll(folder, 0755)
	if err != nil {
		fmt.Printf("SaveFolderCreateError: %v\n", err)
		return
	}

	if len(os.Args) < 2 {
		createHelp()
		return
	}

	switch os.Args[1] {
	case "run":
		bytesData, err := os.ReadFile("version.txt")
		if err != nil {
			fmt.Printf("No version has been specified yet.")
			return
		}
		version := string(bytesData)

		suff := tools.GetDeviceSuff()
		if suff == "null" {
			fmt.Printf("This model is not supported.")
			return
		}
		targetCmd := "yt-dlps/" + version + "/yt-dlp" + suff

		cmd := exec.Command(targetCmd, os.Args[2:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "update":
		if len(os.Args) == 3 {
			if tools.DownloadVersionYtdlp(os.Args[2]) {
				fmt.Println("UpdateSuccess.")
			} else {
				fmt.Println("UpdateFailure.")
			}
			return
		}

		if tools.UpdateYtdlp() {
			fmt.Println("UpdateSuccess.")
		} else {
			fmt.Println("UpdateFailure.")
		}
	case "install":
		if len(os.Args) == 3 {
			if tools.DownloadVersionYtdlp(os.Args[2]) {
				fmt.Println("InstallSuccess.")
			} else {
				fmt.Println("InstallFailure.")
			}
			return
		}

		if tools.UpdateYtdlp() {
			fmt.Println("InstallSuccess.")
		} else {
			fmt.Println("InstallFailure.")
		}
	case "version":
		if len(os.Args) == 3 {
			writeErr := os.WriteFile("version.txt", []byte(os.Args[2]), 0644)
			if writeErr != nil {
				panic(writeErr)
			}
			fmt.Println("UpdatedVersion: " + os.Args[2])
			return
		}

		bytesData, err := os.ReadFile("version.txt")
		if err != nil {
			fmt.Printf("No version has been specified yet.")
			return
		}
		version := string(bytesData)

		fmt.Println(version)
	case "versions":
		entries, err := os.ReadDir(folder)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Versions:")

		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Println(entry.Name())
			} else {
				fmt.Println(entry.Name())
			}
		}

		bytesData, err := os.ReadFile("version.txt")
		if err != nil {
			return
		}
		version := string(bytesData)

		fmt.Println("\nNowVersion: " + version)
	default:
		createHelp()
		os.Exit(1)
	}
}
