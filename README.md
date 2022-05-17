# GoDownloader

一個簡單的 Downloader，可以產生各種 Platform 的 binary。

## Setup

1. 修改 main.go 的下載 URL
2. 修改 build.sh 底下的所需的 GOARCH & GOOS

## Usage

```
$ ./build.sh
```

執行後會建立 bin 資料夾，各種 CPU 架構的執行檔會在底下。

```
$ ls bin
go-downloader-386   go-downloader-arm64
go-downloader-ppc64 go-downloader-amd64
go-downloader-mips
```