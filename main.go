package main

import (
	"fmt"
	"os"
	"source_mate/pkg/qlogin"
	"time"
)

func main() {

	cks := qlogin.BaseCookies()

	qr, sig, _ := qlogin.GetQr()
	//fmt.Println(sig, err)
	file, e := os.OpenFile("./qr.png", os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		fmt.Println("文件打开失败")
		return
	}

	file.Write(qr)
	file.Close()

	for true {
		time.Sleep(time.Duration(1) * time.Second)
		qlogin.PtQrLogin(cks, sig)
	}
}
