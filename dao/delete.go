package dao

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/config"
	"harbor-image-delete/model"
	"io"
	"net/http"
)

func Delete(tmpUrl string) (body []byte, err error) {

	// 定义url路径及参数
	apiUrl := config.HarborURL + tmpUrl
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, apiUrl, nil)

	if err != nil {
		logger.Error("New HTTP 报错: ", err.Error())
		return nil, errors.New(fmt.Sprintf("New HTTP 请求报错: ", err.Error()))
	}

	// 添加 Basic Authentication 头
	auth := config.UserPassword
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)

	// 发起请求
	res, err := client.Do(req)
	if err != nil {
		logger.Error("HTTP 请求报错: ", err.Error())
		return nil, errors.New(fmt.Sprintf("HTTP 请求报错: ", err.Error()))
	}
	defer res.Body.Close()
	defer model.Wg.Done()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		logger.Error("IO 数据解析报错: ", err.Error())
		return nil, errors.New(fmt.Sprintf("IO 数据解析报错: ", err.Error()))
	}

	return body, nil
}
