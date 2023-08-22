package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/model"
)

func Repositories(repositoriesUrl string) (r *[]model.Repositories, err error) {
	// 发起 HTTP 请求
	data, err := Get(repositoriesUrl)

	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))

	// 解析 JSON 数据
	repositories := new([]model.Repositories)
	if err := json.Unmarshal(data, &repositories); err != nil {
		logger.Error("Repositories JSON 数据解析报错:", err.Error())
		return nil, errors.New(fmt.Sprintf("Repositories JSON 数据解析报错:", err.Error()))
	}
	return repositories, nil
}
