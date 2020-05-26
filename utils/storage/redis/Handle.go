package redis

import (
	//"blog-go-api/utils/json"
	//"fmt"
	//"github.com/garyburd/redigo/redis"
	"fmt"
	"sync"
)

//创建一个管道
var RedisChannel = make(map[string]interface{})

var personChan = make(chan Data, 1)

var wg sync.WaitGroup

type Data struct {
	_type    string
	name 	string
	data    string
	ex 	int
}

/**
 * 获取缓存
 */
func Get(name string)  {

	fmt.Print(name)
	//go Handle()
	accessLogMap := make(map[string]interface{})
	accessLogMap["type"] = "set"
	accessLogMap["name"] = name
	fmt.Print(accessLogMap)

	//accessLogJson, _ := json.Encode(accessLogMap)
	//RedisChannel <- accessLogMap
}

/**
 * 设置缓存  没有时间限制
 */
func Set(name string, data string) int {

	wg.Add(1)

	fmt.Print("123123")
	go _Handle()
	//return func() {
		person := Data{"set", name, data, 0}
	fmt.Print("1xxx")
		personChan <- person
	fmt.Print("3xxx")
	wg.Wait()
	fmt.Print("44fffxx")
	return 123213
	//}
}

/**
 * 设置缓存  有时间限制
 */
func SetEx(name string, data string, ex int)  {
	//go Handle()

	//accessLogMap := make(map[string]interface{})
	//accessLogMap["type"] = "setex"
	//accessLogMap["name"] = name
	//accessLogMap["data"] = data
	//accessLogMap["ex"] = ex

	//accessLogMap := 123123
	//accessLogJson, _ := json.Encode(accessLogMap)
	//RedisChannel <- string(accessLogMap)
}

/**
 * 全局处理redis 的函数
 */
func _Handle() int {

	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	return err
	//}
	//defer c.Close()
	//data := <- personChan
	//fmt.Print(data)
	//fmt.Print()
	newPerson := <- personChan
	fmt.Printf("new person %s", newPerson.data)
	defer wg.Add(-1)
	return 123
	//fmt.Print()

	//switch data {
	//case "set":
	//	_, err = c.Do("SET", data.name, data.data)
	//	if err != nil {
	//		return err
	//	}
	//	break
	//case "setex":
	//	_, err = c.Do("SET", data.name, data.data, "EX", data.ex)
	//	if err != nil {
	//		return err
	//	}
	//	break
	//case "get":
	//	data, err := redis.String(c.Do("GET", data.name))
	//	if err != nil {
	//		return err
	//	}
	//	return data
	//	break
	//}
	//return nil
}