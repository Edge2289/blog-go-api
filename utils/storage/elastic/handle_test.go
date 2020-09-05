package elastic

import "testing"

func TestGet(t *testing.T)  {
	//create()
	gets()
}


func TestInsert(t *testing.T)  {
	//create()
	_,err := Insert(100001, `{"first_name":"xiaoxie","last_name":"D.M 001","age":100,"about":"I ove to go rock climbing","interests":["sports","music"]}`)
	if err != nil {
		panic(err)
	}
}