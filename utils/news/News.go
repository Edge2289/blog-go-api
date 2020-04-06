package news

import (
	"errors"
	"fmt"
)

type StructPush interface {
	Push()
}

type Ambient struct {
	Istruct string
	//AmbientPush StructPush
}

func (n *Ambient) Travel() (string, error) {

	fmt.Print(n.Istruct)
	//switch n.Istruct {
	//	case "sms":
	//		Sms{}.Push()
	//		return n.Istruct+" success", nil
	//		break
	//	case "email":
	//		Email{}.Push()
	//		return n.Istruct+" success", nil
	//		break
	//	default:
	//		return "error", errors.New("not News struct ")
	//		break
	//}
	return "error", errors.New("not News struct ")
}
/**
获取需要使用的对象
 */
//func GetStruct(_type string) (Type, string) {
//
//	switch _type {
//		case "email":
//			break
//		default:
//			return (nil, "not news struct ")
//			break
//
//	}
//}