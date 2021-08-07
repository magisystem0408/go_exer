package main

import (
	"fmt"
	"strconv"
)

func main() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 1, f)

	fmt.Println(true && true)
	fmt.Println(false || false)
	fmt.Println(!true)

	//型変換について
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %\n", xx, xx, xx)


	var y float64 =1.2
	yy :=int(y)
	fmt.Printf("%T %v %d\n",yy,yy,yy)


	//文字列から数値変換するとき
	var s string="14"
	i,err :=strconv.Atoi(s)
	if err !=nil{
		fmt.Println("エラー発生した")
	}

	fmt.Printf("%T %v",i,i)


	//文字列を出す時に、バイト変換されるので
	//string型でキャスト変換する必要がある。

	h:="Hello World"
	fmt.Println(string(h[0]))


	//配列の定義
	var a [2]int
	a[0] =100
	a[1] =200

	fmt.Println(a)

	//初期に値を入れたい場合
	var b [2]int =[2]int{100,200}

	//配列はサイズを変更することができない。
	//b =append(b,300)

	fmt.Println(b)

	//こっちであればつけられる
	var z []int =[]int{100,200}
	z=append(z,300)
	fmt.Println(z)

}
