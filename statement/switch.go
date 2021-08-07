package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("mamushi")
	case "windows":
		fmt.Println("nekomamushi")
	default:
		fmt.Println("デフォルトで設定されています")
	}

	//省略記法
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mamushi")
	case "windows":
		fmt.Println("ねこ")
	default:
		fmt.Println("Default!")
	}

	t:=time.Now()
	fmt.Println()
	switch {
	case t.Hour() <12:
		fmt.Println("mornig")
	case t.Hour() <17:
		fmt.Println("afternoon")
	}
}