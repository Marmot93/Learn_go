package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Info struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
	Data map[string]string `json:"data"`
}

func query(phone string, realname string, cardno string) (info string) {
	uri := "http://v.apix.cn/apixcredit/idcheck/mobile/"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		fmt.Println("构建请求错误", err)
	}
	params := req.URL.Query()
	params.Add("phone", phone)
	params.Add("name", realname)
	params.Add("cardno", cardno)
	params.Add("type", "mobile")
	req.URL.RawQuery = params.Encode()
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("apix-key", "****************************")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("请求错误", err)
	}
	defer rsp.Body.Close()
	body, _ := ioutil.ReadAll(rsp.Body)
	return string(body)
}

func main() {
	body := query("18080482962", "牟文峰", "511123199306190****")
	// {"msg": "成功：验证信息一致", "code": 0,
	// "data": {"moible_prov": "四川", "sex": "M", "birthday": "1993-06-19", "address": "四川省乐山市犍为县",
	// "mobile_operator": "四川电信CDMA卡", "mobile_city": "成都"}
	// }
	fmt.Println(body)
	var info Info
	json.Unmarshal([]byte(body), &info)
	fmt.Println(info.Msg)
	fmt.Println(info.Data["address"])
}
