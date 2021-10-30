package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

var (
	file_url string
	filename string
	output   bool
)

func checkArgs() {
	if len(os.Args) <= 1 {
		usage()
	}
	for number, value := range os.Args {
		if validateUrl(os.Args[number]) {
			file_url = value
		}
		if os.Args[number] == "-o" && !validateUrl(os.Args[number+1]) && len(os.Args) >= number+1 {
			output = true
			filename = os.Args[number+1]
		}
	}
	if !output && file_url != "" {
		Odownload()
		os.Exit(0)
	}
	if output == true {
		output = false
		NDownload()
		os.Exit(0)
	}
	usage()
}

func usage() {
	fmt.Println("Get App v1.0")
	fmt.Printf("Usage: %s [-o] [filename] [url]\n", os.Args[0])
	fmt.Printf("Usage: %s [url]\n", os.Args[0])
	os.Exit(0)
}

func NDownload() {
	req, _ := http.NewRequest("GET", file_url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	current_path, _ := os.Getwd()
	filename = fmt.Sprintf("%s\\%s", current_path, filename)
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Println(filename)
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)

}

func Odownload() {
	if file_url == "" {
		fmt.Println("Url is not valid!")
		usage()
	}

	fileURL, err := url.Parse(file_url)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	filename = segments[len(segments)-1]

	req, _ := http.NewRequest("GET", file_url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	current_path, _ := os.Getwd()
	filename = fmt.Sprintf("%s\\%s", current_path, filename)
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Println(filename)
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
}

func validateUrl(_url string) bool {
	_, err := url.ParseRequestURI(_url)
	if err != nil {
		return false
	}

	return true
}

func main() {
	output = false
	checkArgs()
}
