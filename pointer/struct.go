package main

import "fmt"


//構造体について


//変数定義する時は大文字で定義してあげる
type  Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v :=Vertex{X:1,Y:2}
	fmt.Println(v)

	fmt.Println(v.X,v.Y)
	//書き換え
	v.X=100

	v2 :=Vertex{X:1}
	fmt.Println(v2)

	v3 :=Vertex{1,2,"test"}
	fmt.Println(v3)

	v4 :=Vertex{}
	fmt.Println(v4)

	//varで変数宣言返してもok
	var v5 Vertex
	fmt.Printf("%T %v",v5,v5)


	//帰り値がポインタで返ってくるパターンは
	//以下二つで実装可能

	v6 :=new(Vertex)
	fmt.Printf("%T %v" ,v6,v6)

	//アドレスがついているとわかりやすい
	v7 :=&Vertex{}
	fmt.Printf("%T  %v",v7,v7)
}
