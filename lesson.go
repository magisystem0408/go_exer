//なんらかのpackageに属している必要がある
package main

import (
	"fmt"
	"os/user"
	"reflect"
	"time"
)


//大文字でかける
const pi =3.14
const(
	Username ="test_user"
	Password ="test_pass"
)


func init()  {
	fmt.Println("コンストラクタだよ")
}

//関数宣言
func bazz()  {
	fmt.Println("timi")
}

//これを必ず書く
func main()  {

	//変数の宣言方法
	//var num int
	//num =1
	num :=1
	fmt.Println(num)

	fmt.Println("hello mamushi","テストマムシ",time.Now())
	fmt.Println(user.Current())

	//データ型の確認
	fmt.Println(reflect.TypeOf(num))

	//関数の実行
	bazz()
}