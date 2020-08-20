package qrcode

import "github.com/skip2/go-qrcode"

func Create() {
	qrcode.WriteFile("http://c.biancheng.net/", qrcode.Medium, 256, "./golang_qrcode.png")
}
