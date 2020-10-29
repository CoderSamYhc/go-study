package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init () {

	// 日志格式化为json
	log.SetFormatter(&log.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)

	// 记录级别
	log.SetLevel(log.InfoLevel)

}

func main () {
	log.WithFields(log.Fields{
		"code" : "values",
	}).Info("Info...")

	log.WithFields(log.Fields{
		"code" : "values",
	}).Warn("Warn...")

	log.WithFields(log.Fields{
		"code" : "values",
	}).Fatal("Fatal...")
}
