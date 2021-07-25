package main

import (
	"flag"

	"github.com/hb-go/echo-web/module/log"
	"github.com/hb-go/echo-web/router"
)

const (
	DefaultConfFilePath = "conf/conf.toml"
)

var (
	confFilePath string
	cmdHelp      bool
)

func init() {
	flag.StringVar(&confFilePath, "c", DefaultConfFilePath, "配置文件路径")
	flag.BoolVar(&cmdHelp, "h", false, "帮助")
	flag.Parse()

}

func main() {
	if cmdHelp {
		flag.PrintDefaults()
		return
	}
	log.Debugf("run with conf:%s", confFilePath)

	// 子域名部署
	router.RunSubdomains(confFilePath)
}
