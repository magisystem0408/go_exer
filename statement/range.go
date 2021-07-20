package main

import "fmt"

func main() {
	l := []string{"python","mamushi","java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i,l[i])
	}

	//以下のようにも記述することができる
	for i,v :=range l{
		fmt.Println(i,v)
	}

	//iのindex番号を使いたくない
	for _,v :=range l{
		fmt.Println(v)
	}

	m:= map[string]int{"apple":100,"banana":200}
	for k,v :=range m{
		fmt.Println(k,v)
	}

	//valueはいらない時は記述しなくて良い

	for k :=range m{
		fmt.Println(k)
	}

	//keyが欲しい時はアンダースコアをつける

	for _,k :=range m{
		fmt.Println(k)
	}

}