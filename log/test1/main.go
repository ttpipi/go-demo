package main

import (
	"log"
	"os"
)

func init() {
	log.SetPrefix("[test1]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err := os.OpenFile("E:/golang/golang/log/test1/log-file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败: ", err)
	}
	log.SetOutput(logFile)
}

func main() {
	log.Println("log.SetFlags")
}
