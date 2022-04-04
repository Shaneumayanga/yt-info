package ytdl

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const URL = "https://api.github.com/repos/ytdl-org/youtube-dl/releases?per_page=1"
const FILE_NAME_LINUX = "youtube-dl"
const FILE_NAME_WINDOWS = "youtube-dl.exe"
const DOWNLOAD_URL_NAME = "browser_download_url"
const FILE_MODE = 493

func init() {
	goos := runtime.GOOS
	if !checkFile(goos) {
		log.Println("Starting download")
		download_url := getBin(goos)
		downloadBin(download_url)
	}
}

func Info(url string) *Output {
	goos := runtime.GOOS
	data := &Output{
		Formats: make([]Format, 0),
	}
	if goos == "linux" {
		output, err := os.Create("output.json")
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()
		cmd := exec.Command("./"+FILE_NAME_LINUX, url, "--dump-single-json")
		cmd.Stdout = output
		if err := cmd.Run(); err != nil {
			log.Println(err)
		}
		b, err := ioutil.ReadFile("output.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(b, data)
		if err != nil {
			panic(err)
		}
		return data

	} else if goos == "windows" {
		output, err := os.Create("output.json")
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()
		cmd := exec.Command(FILE_NAME_WINDOWS, url, "--dump-single-json")
		cmd.Stdout = output
		if err := cmd.Run(); err != nil {
			log.Println(err)
		}
		b, err := ioutil.ReadFile("output.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(b, data)
		if err != nil {
			panic(err)
		}
		return data
	}
	return nil
}
