package main

import "fmt"

func main() {

	//メモリにポインタが入るところを確保するには
	var p *int =new(int)
	//アドレスが帰ってくる
	fmt.Println(p)
	//実体の数字は
	fmt.Println(*p)
	*p++
	fmt.Println(*p)


	//以下で実行するとnewでメモリを確保していないのでnilが帰ってくる
	var p2 *int
	fmt.Println(p2)


	//ポインタが帰ってくる帰り値に関しては、newを使用しそう出ない場合はmakeを使用する

	s :=make([]int,0)
	fmt.Printf("%T\n",s)

	m := make(map[string]int)
	fmt.Printf("%T\n",m)


	var p *int =new(int)
	fmt.Printf("%T\n",p)


}