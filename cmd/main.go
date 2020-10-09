package main

import (
	"fmt"
	"gatewayH/internal/gateway/server"
	"github.com/alecthomas/kingpin"

	"os"
)

var _version_ = ""
var _branch_ = ""

func main() {

	// 1.从输入参数读取文件名，比如 ./agent --filename=../../configs/goods.yaml
	filename := kingpin.Flag("filename", "Config file name.").String()
	version := kingpin.Flag("version", "Display Version.").Bool()
	kingpin.Parse()
	fmt.Printf("Version: %s\nBranch: %s\n", _version_, _branch_)
	if *version {
		os.Exit(-1)
	}
	server.Run(*filename)
}
