package ytdl

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

func checkFile(goos string) bool {
	if goos == "linux" {
		_, err := os.ReadFile(FILE_NAME_LINUX)
		if err != nil {
			return false
		}
	}
	if goos == "windows" {
		_, err := os.ReadFile(FILE_NAME_WINDOWS)
		if err != nil {
			return false
		}
	}
	return true
}

//gets the download URL according to os
func getBin(os string) string {
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	jsonParsed, err := gabs.ParseJSON(b)
	if err != nil {
		panic(err)
	}
	for _, child := range jsonParsed.Path("0.assets").Children() {
		data := child.Data().(map[string]interface{})
		downloadURL := data[DOWNLOAD_URL_NAME].(string)
		sp := strings.Split(downloadURL, "/")
		file_name := sp[len(sp)-1]
		if file_name == FILE_NAME_LINUX && os == "linux" {
			return downloadURL
		}
		if file_name == FILE_NAME_WINDOWS && os == "windows" {
			return downloadURL
		}
	}
	return ""
}
