package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

)

type Response struct {
	Status   int    `json:"status"`
	CityName string `json:"city"`
	Data     Data   `json:"data"`
	Date     string `json:"date"`
	Message  string `json:"message"`
	Count    int    `json:"count"`
}

// 响应数据
type Data struct {
	ShiDu     string `json:"shidu"`
	Quality   string `json:"quality"`
	Ganmao    string `json:"ganmao"`
	Yesterday Day    `json:"yesterday"`
	Forecast  []Day  `json:"forecast"`
}

// 某一天的数据
type Day struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float32 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}


func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}


func Print(day string, r Response) {
	fmt.Println("城市:", r.CityName)

	if day == "今天" {
		fmt.Println("湿度:", r.Data.ShiDu)
		fmt.Println("空气质量:", r.Data.Quality)
		fmt.Println("温馨提示:", r.Data.Ganmao)
	} else if day == "昨天" {
		fmt.Println("日期:", r.Data.Yesterday.Date)
		fmt.Println("温度:", r.Data.Yesterday.Low, r.Data.Yesterday.High)
		fmt.Println("风量:", r.Data.Yesterday.Fx, r.Data.Yesterday.Fl)
		fmt.Println("天气:", r.Data.Yesterday.Type)
		fmt.Println("温馨提示:", r.Data.Yesterday.Notice)
	} else if day == "预测" {
		fmt.Println("====================================")
		for _, item := range r.Data.Forecast {
			fmt.Println("日期:", item.Date)
			fmt.Println("温度:", item.Low, item.High)
			fmt.Println("风量:", item.Fx, item.Fl)
			fmt.Println("天气:", item.Type)
			fmt.Println("温馨提示:", item.Notice)
			fmt.Println("====================================")
		}
	} else {
		fmt.Println("大熊你是想刁难我胖虎 ?_?")
	}
}

func main() {
	var city string
	var day string

	flag.StringVar(&city, "c", "成都", "城市中文名")
	flag.StringVar(&day, "d", "今天", "可选: 今天, 昨天, 预测")
	flag.Parse()

	var r Response
	body, err := Request(fmt.Sprintf("http://www.sojson.com/open/api/weather/json.shtml?city=%s",city))
	fmt.Println(body)
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("\nError message: %v", err)
	}
	if r.Status != 200 {
		fmt.Printf("获取天气API出现错误, %s", r.Message)
		return
	}
	Print(day, r)
}
