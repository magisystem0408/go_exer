package _struct

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string{
	//中身を書き換えたい時は↑の引数に、アスタリスクをつける
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human)  {
	if human.Say() =="Mr.Mike"{
		fmt.Println("Run")
	}else {
		fmt.Println("Get out")
	}
}

func main() {
	//中身を書き換えるときはポインタレシーバーの&をつけないといけない。
	var mike Human = &Person{"Mike"}

	//humanの中にsayというメソッドが入っていないとだめ
	mike.Say()


	DriveCar(mike)
}
