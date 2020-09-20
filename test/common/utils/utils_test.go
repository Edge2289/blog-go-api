package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestStringArray1(t *testing.T) {
	authString := "1,2,3,4,5,67,8"
	kv := strings.Split(authString, ",")
	fmt.Println(kv)
}

func TestStringArray(t *testing.T)  {
	var src = []int{1,2,3,4}
	var temp = make([]string, len(src))
	for k, v := range src {
		temp[k] = fmt.Sprintf("%d", v)
	}
	var result1 = "[" + strings.Join(temp, ",") + "]"
	fmt.Println(result1)

	var a = []int{1,2,3,4}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	var result = string(b)
	fmt.Println(result)

}
