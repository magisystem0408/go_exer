package main

import (
	"fmt"
	"regexp"
)

func main() {

	//正規表現を一回のみ実行するパターン
	match, _ := regexp.MatchString("a([a-z]+)e","apple")
	fmt.Println(match)

	//複数回正規表現を実行したいパターン
	r := regexp.MustCompile("(a[a-z]+)e")
	ms :=r.MatchString("apple")
	fmt.Println(ms)

}