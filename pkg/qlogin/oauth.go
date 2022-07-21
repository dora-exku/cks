package qlogin

import (
	"fmt"
	"net/http"
	"source_mate/pkg/utils"
	"strings"
	"time"
)

func Authorize(cks []*http.Cookie, clientId, redirectUri string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url := "https://graph.qq.com/oauth2.0/authorize"

	var params string
	params += "response_type=code"
	params += "&client_id=" + clientId
	params += "&redirect_uri=" + redirectUri
	params += "&scope=get_user_info%2Cadd_share%2Cadd_pic_t"
	params += "&state="
	params += "&switch="
	params += "&from_ptlogin=1"
	params += "&src=1"
	params += "&update_auth=1"
	params += "&openapi=80901010"
	params += fmt.Sprintf("&g_tk=%d", Gtk(utils.GetCookie(cks, "p_skey")))
	params += "&auth_time=1658378499340"
	params += fmt.Sprintf("&auth_time=%d", time.Now().UnixMilli())
	params += "&ui=7EF99152-7F24-45F0-816D-F9B17088F57E"

	req, _ := http.NewRequest("POST", url, strings.NewReader(params))

	for _, ck := range cks {
		req.AddCookie(ck)
	}

	req.Header.Set("Host", "graph.qq.com")
	req.Header.Set("Origin", "https://graph.qq.com")
	req.Header.Set("Referer", "https://graph.qq.com/oauth2.0/show?which=Login&display=pc&client_id="+clientId+"&redirect_uri="+redirectUri+"&response_type=code&scope=get_user_info%2Cadd_share%2Cadd_pic_t")

	resp, _ := client.Do(req)

	return resp.Header.Get("Location")
}
