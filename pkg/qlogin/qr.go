package qlogin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"source_mate/pkg/utils"
	"time"
)

func BaseCookies() []*http.Cookie {
	url := `https://xui.ptlogin2.qq.com/cgi-bin/xlogin?appid=716027609&daid=383&style=33&login_text=%E7%99%BB%E5%BD%95&hide_title_bar=1&hide_border=1&target=self&s_url=https%3A%2F%2Fgraph.qq.com%2Foauth2.0%2Flogin_jump&pt_3rd_aid=101252414&pt_feedback_link=https%3A%2F%2Fsupport.qq.com%2Fproducts%2F77942%3FcustomInfo%3D588ku.com.appid101252414&theme=2&verify_theme=`
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	return resp.Cookies()
}

func GetQr() ([]byte, string, error) {
	url := `https://ssl.ptlogin2.qq.com/ptqrshow?appid=716027609&e=2&l=M&s=3&d=72&v=4&t=0.10998353940684247&daid=383&pt_3rd_aid=101252414`

	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	cks := resp.Cookies()
	var qrsig string
	for _, item := range cks {
		if item.Name == "qrsig" {
			qrsig = item.Value
		}
	}

	return body, qrsig, nil
}

func PtQrLogin(cks []*http.Cookie, sig string) {
	url := `https://ssl.ptlogin2.qq.com/ptqrlogin?`
	url += "u1=https%3A%2F%2Fgraph.qq.com%2Foauth2.0%2Flogin_jump"
	url += fmt.Sprintf("&ptqrtoken=%d", Hash33(sig))
	url += "&ptredirect=0"
	url += "&h=1"
	url += "&t=1"
	url += "&g=1"
	url += "&from_ui=1"
	url += "&ptlang=2052"
	url += fmt.Sprintf("&action=1-1-%d", time.Now().UnixMilli())
	url += "&js_ver=22071217"
	url += "&js_type=1"
	url += "&login_sig=" + utils.GetCookie(cks, "pt_login_sig")
	url += "&pt_uistyle=40"
	url += "&aid=716027609"
	url += "&daid=383"
	url += "&pt_3rd_aid=101252414"
	url += "&o1vId=278c614cd260765a7fb4addeb4b97bf5"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	for _, ck := range cks {
		req.AddCookie(&http.Cookie{
			Name:  ck.Name,
			Value: ck.Value,
		})
	}
	req.AddCookie(&http.Cookie{
		Name:  "qrsig",
		Value: sig,
	})
	req.Header.Set("Host", "ssl.ptlogin2.qq.com")
	req.Header.Set("Referer", "https://xui.ptlogin2.qq.com/")

	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Cookies())
	fmt.Println(string(body))
}
