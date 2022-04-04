package ytdl

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadBin(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	sp := strings.Split(url, "/")
	file_name := sp[len(sp)-1]
	file, err := os.Create(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatalf("Failed to download %s", err.Error())
	}
	log.Println("Downloaded...")
	if file_name == FILE_NAME_LINUX {
		err := os.Chmod(FILE_NAME_LINUX, FILE_MODE)
		if err != nil {
			log.Fatal(err)
		}
	}
}
