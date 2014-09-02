package httprequest

import (
	"testing"
	"fmt"
)

func TestNewWithDefaults(t *testing.T) {
	fmt.Println("httprequest.NewWithDefaults()")
	hr := NewWithDefaults()
	fmt.Println("httprequest.Dump()")
	hr.Dump()
}

func TestHttpProxy(t *testing.T) {
	fmt.Println("httprequest.UseProxy()")
	hr := NewWithDefaults()
	proxy_url := "http://127.0.0.1"
	proxy_flag := true
	hr.ProxyUrl = proxy_url
	if hr.ProxyUrl != proxy_url {
		t.Fail()
	}
	hr.UseProxy(proxy_flag)
	if hr.ProxyFlag != proxy_flag {
		t.Fail()
	}
	hr.Dump()
}

func TestHttpBasicAuth(t *testing.T) {
	fmt.Println("httprequest.UseBasicAuth()")
	hr := NewWithDefaults()
	basic_auth_flag := true
	username := "Hello"
	password := "World"
	hr.UseBasicAuth(basic_auth_flag)
	if hr.BasicAuthFlag != basic_auth_flag {
		t.Fail()
	}
	hr.Username = username
	if hr.Username != username {
		t.Fail()
	}
	hr.Password = password
	if hr.Password != password {
		t.Fail()
	}
	hr.Dump()
}

func TestHttpGet(t *testing.T) {
	fmt.Println("httprequest.Get()")
	hr := NewWithDefaults()
	url := "http://www.google.com"
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	body, status_code, err := hr.Get(url, headers)
	if err == nil {
		fmt.Println("Status Code: ", status_code)
		fmt.Println("Body:        ", string(body))
	} else {
		fmt.Println("Error:       ", err)
	}
}

func TestHttpsGet(t *testing.T) {
	fmt.Println("httprequest.Get()")
	hr := NewWithDefaults()
	url := "https://www.google.com"
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	body, status_code, err := hr.Get(url, headers)
	if err == nil {
		fmt.Println("Status Code: ", status_code)
		fmt.Println("Body:        ", string(body))
	} else {
		fmt.Println("Error:       ", err)
	}
}