package main

import (
	"fmt"
	"sort"
)

func main() {
	i :=[]int{5,3,2,8,7}
	s :=[]string{"d","a","f"}
	p :=[]struct{
		Name string
		Age int
	}{
		{"mamushi",20},
		{"ねこ",5},
		{"ニャンコ",10},
		{"timitimiジャンクそん",20},


	}
	fmt.Println(i,s,p)

	//ここからソートする
	sort.Ints(i)
	sort.Strings(s)



	fmt.Println(i)

}