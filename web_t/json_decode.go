package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/bitly/go-simplejson"
)

type Info struct {
	NewsKindsInfo []NewsKind `json:"news_kinds_info"`
}

type NewsKind struct {
	NewsKind int    `json:"news_kind"`
	Name     string `json:"name"`
}

// 请求
func get() (rsp string, err error) {
	response, err := http.Get("http://120.77.237.231:8001/consumer/index/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func req() (rsp string, err error) {
	// 创建Client
	client := &http.Client{}
	// 构建请求参数
	data := url.Values{}
	data.Add("phone", "18080482962")
	data.Add("password", "123456")
	// io.Reader
	req_data := strings.NewReader(data.Encode())
	// 构建请求 method, url, io.Reader
	req, err := http.NewRequest("POST", "http://120.77.237.231:8001/seller/log_in/", req_data)
	// 构建请求头
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 发起请求
	response, err := client.Do(req)
	// 解析
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	// 返回
	return string(body), nil
}

func main() {
	body, _ := get()
	var info Info
	rsp := simplejson.NewJson([]byte(body))
	// byte  &interface
	json.Unmarshal([]byte(body), &info)
	fmt.Println(info.NewsKindsInfo[0].Name)
}
