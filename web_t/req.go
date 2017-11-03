package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
	//"strings"
	"io"
	"strings"
)

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

// 构建请求
func req(method string, url string, data io.Reader) (rsp string, err error) {
	// 创建Client
	client := &http.Client{}
	// 构建请求 method(""==GER), url, io.Reader
	req, err := http.NewRequest(method, url, data)
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

func pe(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// 调用拼接
	//data := url.Values{}
	//data.Add("phone", "18080482962")
	//data.Add("password", "123456")
	//fmt.Println(data.Encode())  // password=123456&phone=18080482962
	//req_data := strings.NewReader(data.Encode())


	// 自己写
	//req_data := strings.NewReader("password=123456&phone=18080482962")

	// 自己拼接
	data := make(map[string]string)
	data["phone"] = "18080482962"
	data["password"] = "123456"
	//var info string
	info := ""
	for k, v := range data{
		info = info + fmt.Sprintf("%s=%s&", k,v)
	}
	req_data := strings.NewReader(info[:len(info)-1])
	//rsp, err := req("", "http://120.77.237.231:8001/consumer/index/", req_data)
	rsp, err := req("POST", "http://120.77.237.231:8001/seller/log_in/", req_data)
	pe(err)
	fmt.Println(rsp)

}
