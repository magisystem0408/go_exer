package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//0666は読み書き
	loggfile,_ :=os.OpenFile(logFile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)

	multiLogFile :=io.MultiWriter(os.Stderr,loggfile)
	//フラグのセット
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)

}

func main() {

	LoggingSettings("test.log")

	log.Println("logging!")
	log.Printf("%T %v","test","test")

	//これやるとプログラムが終了する
	log.Fatalln("error!!!")


	//応用例-ファイルを開くとき
	_,err :=os.Open("aaaaaaaa")
	if err !=nil{
		log.Fatalln("Exit")
	}

}