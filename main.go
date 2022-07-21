package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"source_mate/pkg/qlogin"
	"time"
)

func main() {

	// 千库网 588ku.com
	clientId := "101252414"
	redirectUrl := "https%3A%2F%2F588ku.com%2Fdlogin%2Fcallback%2Fqq"

	//qlogin.OauthCookie()
	//return

	cks := qlogin.BaseCookies(clientId)

	qr, sig, _ := qlogin.GetQr(clientId)
	//fmt.Println(sig, err)
	file, e := os.OpenFile("./qr.png", os.O_CREATE|os.O_WRONLY, 0666)
	if e != nil {
		fmt.Println("文件打开失败")
		return
	}

	file.Write(qr)
	file.Close()

	var lcks []*http.Cookie
	var url string
	for true {
		time.Sleep(time.Duration(1) * time.Second)
		loginCks, loginRes := qlogin.PtQrLogin(cks, sig, clientId)
		if loginRes.Code == 0 {
			url = loginRes.Url
			lcks = loginCks
			break
		}
		fmt.Println(loginRes.Msg)
	}

	ocks := qlogin.CheckSig(lcks, url)

	// 获取三方授权重定向链接
	redrUrl := qlogin.Authorize(ocks, clientId, redirectUrl)
	fmt.Println(redrUrl)

	// 以下为第三方处理逻辑
	client := &http.Client{}
	req, _ := http.NewRequest("GET", redrUrl, nil)
	//req.Header.Set("Host", "588ku.com")
	req.Header.Set("Referer", "https://graph.qq.com/")
	//req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Cookies())
	fmt.Println(string(body))
}
