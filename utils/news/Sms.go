package news

import "fmt"

type Sms struct {
	// 这里个人认为不是继承而是组合，在重构关键技法中这也是一种提倡做法，继承——>组合
	//sms NewsStructPush
}

func (s *Sms) Push() {
	fmt.Print("Sms Push ")
}