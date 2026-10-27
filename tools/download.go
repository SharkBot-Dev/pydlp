package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v93/github"
)

func DownloadFile(url, dest string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("RequestError (%s): %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("FileCreateError (%s): %v\n", dest, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("SaveError (%s): %v\n", dest, err)
		return
	}

	fmt.Println("DownloadOk.")
}

func UpdateYtdlp() bool {
	folder := "yt-dlps"
	device, suffixFile := GetDeviceInfo()

	if device == "null" {
		fmt.Printf("This model is not supported.")
		return false
	}

	client, err := github.NewClient()
	if err != nil {
		panic(err)
	}

	release, _, repErr := client.Repositories.GetLatestRelease(context.Background(), "yt-dlp", "yt-dlp")
	if repErr != nil {
		panic(repErr)
	}

	assets := release.GetAssets()
	for _, asset := range assets {
		if strings.HasSuffix(*asset.BrowserDownloadURL, device) {
			fmt.Println("Downloading..")
			fmt.Println("Url: " + *asset.BrowserDownloadURL)

			downloadFolder := folder + "/" + release.TagName + "/" + "yt-dlp" + suffixFile

			err := os.MkdirAll(folder+"/"+release.TagName, 0755)
			if err != nil {
				fmt.Printf("SaveFolderCreateError: %v\n", err)
				return false
			}

			fmt.Println("SavedFolder: " + downloadFolder)

			DownloadFile(*asset.BrowserDownloadURL, downloadFolder)

			writeErr := os.WriteFile("version.txt", []byte(release.TagName), 0644)
			if writeErr != nil {
				panic(writeErr)
			}
			fmt.Println("UpdatedVersion: " + release.TagName)

			return true
		}
	}

	return false
}

func DownloadVersionYtdlp(version string) bool {
	folder := "yt-dlps"
	device, suffixFile := GetDeviceInfo()

	if device == "null" {
		fmt.Printf("This model is not supported.")
		return false
	}

	client, err := github.NewClient()
	if err != nil {
		panic(err)
	}

	release, _, repErr := client.Repositories.GetReleaseByTag(context.Background(), "yt-dlp", "yt-dlp", version)
	if repErr != nil {
		panic(repErr)
	}

	assets := release.GetAssets()
	for _, asset := range assets {
		if strings.HasSuffix(*asset.BrowserDownloadURL, device) {
			fmt.Println("Downloading..")
			fmt.Println("Url: " + *asset.BrowserDownloadURL)

			downloadFolder := folder + "/" + release.TagName + "/" + "yt-dlp" + suffixFile

			err := os.MkdirAll(folder+"/"+release.TagName, 0755)
			if err != nil {
				fmt.Printf("SaveFolderCreateError: %v\n", err)
				return false
			}

			fmt.Println("SavedFolder: " + downloadFolder)

			DownloadFile(*asset.BrowserDownloadURL, downloadFolder)

			writeErr := os.WriteFile("version.txt", []byte(release.TagName), 0644)
			if writeErr != nil {
				panic(writeErr)
			}
			fmt.Println("UpdatedVersion: " + release.TagName)

			return true
		}
	}

	return false
}
