package utils

import "net/http"

func GetCookie(cks []*http.Cookie, name string) string {
	for _, item := range cks {
		if item.Name == name {
			return item.Value
		}
	}
	return ""
}
