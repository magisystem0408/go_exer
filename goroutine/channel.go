//channelは並列化したときにデータのやりとりを行う
package main

import "fmt"

func goroutine(s []int,c chan int) {
	sum :=0
	for _,v:=range s{
		sum +=v
	}
	c <- sum
}

func main() {
	//スライスとチャネルを作る
	s :=[]int{1,2,3,4,5}
	c :=make(chan int)

	go goroutine(s,c)

	//チャネルから値を受け取る
	x := <-c
	fmt.Println(x)
}