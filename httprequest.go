package httprequest

import (
	"os"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"mime/multipart"
	"bytes"
	"strings"
	"net/url"
	"github.com/mreiferson/go-httpclient"
	"crypto/tls"
)

type HttpRequest struct {
	Transport *httpclient.Transport
	Client *http.Client
	ProxyFlag bool
	BasicAuthFlag bool
	Username string
	Password string
	ProxyUrl string
}

func NewWithDefaults() (*HttpRequest) {
	timeout := 10000 * time.Millisecond
	return New(timeout, timeout, timeout, true, true, false)
}

func New(connection_timeout time.Duration, response_header_timeout time.Duration, request_timeout time.Duration, keep_alive_flag bool, compression_flag bool, skip_tls_verify bool) (*HttpRequest) {
	transport := &httpclient.Transport {
		DisableKeepAlives: !keep_alive_flag,
		DisableCompression: !compression_flag,
		ConnectTimeout: connection_timeout,
		ResponseHeaderTimeout: response_header_timeout,
		RequestTimeout: request_timeout,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skip_tls_verify},
	}

	
	defer transport.Close()
	client := &http.Client{ Transport: transport }
	r := HttpRequest{ Transport: transport, Client: client, ProxyFlag: false, BasicAuthFlag: false, Username: "", Password: "", ProxyUrl: "" }
	return &r
}

func (r *HttpRequest) Dump() {
	fmt.Println("HttpRequest")
	fmt.Println("-----------")
	fmt.Println("\tProxyFlag:      ", r.ProxyFlag)
	fmt.Println("\tBasicAuthFlag:  ", r.BasicAuthFlag)
	fmt.Println("\tUsername:       ", r.Username)
	fmt.Println("\tPassword:       ", r.Password)
	fmt.Println("\tProxyUrl:       ", r.ProxyUrl)
}

func (r *HttpRequest) UseBasicAuth(flag bool) {
	r.BasicAuthFlag = flag
}

func (r *HttpRequest) UseProxy(flag bool) {
	r.ProxyFlag = flag
}

func (r *HttpRequest) Get(url string, headers map[string]string) ([]byte, int, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, -1, err
	}
	if r.ProxyFlag {
		os.Setenv("HTTP_PROXY", r.ProxyUrl)
	}
	if r.BasicAuthFlag {
		request.SetBasicAuth(r.Username, r.Password)
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, val)
		}
	}
	response, err := r.Client.Do(request)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()
	status_code := response.StatusCode
	contents, err := ioutil.ReadAll(response.Body)
	neterr, ok := err.(net.Error)
	if ok && neterr.Timeout() {
		err = nil
	}
	if err != nil {
		return nil, status_code, err
	}
	return contents, status_code, nil
}

func (r *HttpRequest) PostUrlEncoded(url string, headers map[string]string, data url.Values) ([]byte, int, error) {
	body := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, -1, err
	}
	if r.ProxyFlag {
		os.Setenv("HTTP_PROXY", r.ProxyUrl)
	}
	if r.BasicAuthFlag {
		request.SetBasicAuth(r.Username, r.Password)
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, val)
		}
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := r.Client.Do(request)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()
	status_code := response.StatusCode
	contents, err := ioutil.ReadAll(response.Body)
	neterr, ok := err.(net.Error)
	if ok && neterr.Timeout() {
		err = nil
	}
	if err != nil {
		return nil, status_code, err
	}
	return contents, status_code, nil
}

func (r *HttpRequest) PostMultipart(url string, headers map[string]string, multipart_params map[string]string) ([]byte, int, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range multipart_params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, -1, err
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, -1, err
	}
	if r.ProxyFlag {
		os.Setenv("HTTP_PROXY", r.ProxyUrl)
	}
	if r.BasicAuthFlag {
		request.SetBasicAuth(r.Username, r.Password)
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, val)
		}
	}
	response, err := r.Client.Do(request)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()
	status_code := response.StatusCode
	contents, err := ioutil.ReadAll(response.Body)
	neterr, ok := err.(net.Error)
	if ok && neterr.Timeout() {
		err = nil
	}
	if err != nil {
		return nil, status_code, err
	}
	return contents, status_code, nil
}

func (r *HttpRequest) Put(url string, headers map[string]string, data string) ([]byte, int, error) {
	body := &bytes.Buffer{}
	_, err := body.WriteString(data)
	if err != nil {
		return nil, -1, err
	}
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, -1, err
	}
	if r.ProxyFlag {
		os.Setenv("HTTP_PROXY", r.ProxyUrl)
	}
	if r.BasicAuthFlag {
		request.SetBasicAuth(r.Username, r.Password)
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Set(key, val)
		}
	}
	response, err := r.Client.Do(request)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()
	status_code := response.StatusCode
	contents, err := ioutil.ReadAll(response.Body)
	neterr, ok := err.(net.Error)
	if ok && neterr.Timeout() {
		err = nil
	}
	if err != nil {
		return nil, status_code, err
	}
	return contents, status_code, nil
}
