package dao

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/config"
	"io"
	"net/http"
	"net/url"
)

func Get(tmpUrl string) (body []byte, err error) {
	//定义url路径及参数
	apiUrl := config.HarborURL + tmpUrl

	//设置请求参数
	data := url.Values{}
	data.Set("Content-Type", "application/json")
	data.Set("page_size", "100")

	// 添加 Basic Authentication 头
	auth := config.UserPassword
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	data.Set("Authorization", authHeader)

	// 拼接完整url
	u, _ := url.ParseRequestURI(apiUrl)
	u.RawQuery = data.Encode()
	//fmt.Println("请求完整路径为", u.String())

	// 发起请求
	resp, err := http.Get(u.String())
	if err != nil {
		logger.Error("HTTP 请求报错: ", err.Error())
		return nil, errors.New(fmt.Sprintf("HTTP 请求报错: ", err.Error()))
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("IO 获取数据报错: ", err.Error())
		return nil, errors.New(fmt.Sprintf("HTTP 获取数据报错: ", err.Error()))
	}

	return body, nil
}
