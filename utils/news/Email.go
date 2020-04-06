package news

import "fmt"

type Email struct {
	// 这里个人认为不是继承而是组合，在重构关键技法中这也是一种提倡做法，继承——>组合
	//email NewsStructPush
}

/**
 邮箱发送
 */
func (e *Email) Push() {
	fmt.Print("email Push")
}