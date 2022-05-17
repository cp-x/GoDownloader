package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func downloadFile(url string) {
	filename := "downloaded_file"
	out, _ := os.Create(filename)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
}

func main() {
	os := runtime.GOOS
	arch := runtime.GOARCH
	url := fmt.Sprintf("https://example.com/?os=%s&arch=%s", os, arch)
	fmt.Println(url)
	downloadFile(url)
}
