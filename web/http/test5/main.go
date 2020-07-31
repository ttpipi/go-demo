/**
cookie的使用
*/

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

func setCookie(rw http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "go",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "go go",
		HttpOnly: true,
	}

	//use Header
	//rw.Header().Set("Set-Cookie", c1.String())
	//rw.Header().Add("Set-Cookie", c2.String())

	//or http.SetCookie
	http.SetCookie(rw, &c1)
	http.SetCookie(rw, &c2)
}

func getCookie(rw http.ResponseWriter, r *http.Request) {
	//use Header
	//c := r.Header["Cookie"]
	//fmt.Fprintln(rw, c)

	//or r.Cookie, r.Cookies
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(rw, "Cannot get the first cookie")
	}

	cs := r.Cookies()
	fmt.Fprintln(rw, c1)
	fmt.Fprintln(rw, cs)
}

func setMessage(rw http.ResponseWriter, r *http.Request) {
	msg := []byte("hello world")
	c := http.Cookie{
		Name: "flash",
		//满足响应首部对cookie值的URL编码要求, 如果没有空格或百分号等特殊字符, 可以不进行编码
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(rw, &c)
}

func showMessage(rw http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(rw, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(rw, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(rw, string(val))
	}
}

func main() {
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	//实现闪现消息
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/get_message", showMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
