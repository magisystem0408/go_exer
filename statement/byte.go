package main

import "fmt"

func main() {
	b := []byte{72,73}
	fmt.Println(b)

	fmt.Println(string(b))
	//出力結果
	//HI

	c := []byte("HI")
	fmt.Println(c)
}

