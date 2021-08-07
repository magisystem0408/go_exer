package main

import "fmt"

//panic：強制的にエラーを出してくれる
//recover：panicをリカバーする

func thirdPartyConnectDB(){
	//強制的にエラーを出してくれる
	panic("Unable to connect database!")
}

func save() {
	//最初の方に宣言する必要がある。
	defer func() {
		s:=recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("Ok?")
}


