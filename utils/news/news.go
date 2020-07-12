package news


type NewsStruct interface {
	Push()
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