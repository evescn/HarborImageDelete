package dao

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/config"
	"harbor-image-delete/model"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Put
// 发起一个 Get 请求
func Put(tmpUrl string) (p *[]model.Projects, err error) {
	//定义url路径及参数
	apiUrl := config.HarborURL + tmpUrl

	//设置请求参数
	data := url.Values{}
	data.Set("Content-Type", "application/json")

	// 添加 Basic Authentication 头
	auth := config.UserPassword
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	data.Set("Authorization", authHeader)

	//拼接完整url
	u, _ := url.ParseRequestURI(apiUrl)
	u.RawQuery = data.Encode()
	fmt.Println("请求完整路径为", u.String())

	//发起请求
	resp, _ := http.Get(u.String())
	body, _ := ioutil.ReadAll(resp.Body)

	// 解析 JSON 数据
	projects := new([]model.Projects)
	if err := json.Unmarshal(body, &projects); err != nil {
		logger.Error("Projects JSON 数据解析报错:", err)
		return nil, errors.New(fmt.Sprintf("Projects JSON 数据解析报错:", err))
	}
	return projects, nil
}
