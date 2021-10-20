package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fghwett/toolsign/config"
	"github.com/fghwett/toolsign/notify"
	"github.com/fghwett/toolsign/task"
	"github.com/fghwett/toolsign/util"
)

var path = flag.String("path", "./config.yml", "配置文件地址")

func main() {
	flag.Parse()

	conf, err := config.Init(*path)
	if err != nil {
		fmt.Printf("读取配置文件失败 err: %s", err)
		os.Exit(-1)
	}

	util.BigSleep(5, 20)

	t := task.New(conf.Config)
	t.Do()

	if err := notify.Send(conf.ServerChan.SecretKey, "Tool.lu签到任务", t.GetResult()); err != nil {
		log.Printf("通知发送失败 %s\n", err)
	}
}
