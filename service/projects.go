package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/dao"
	"harbor-image-delete/model"
)

func Projects(projectsUrl string) (p *[]model.Projects, err error) {
	// 发起 HTTP 请求
	data, err := dao.Get(projectsUrl)

	if err != nil {
		return nil, err
	}
	//logger.Info(string(data))

	// 解析 JSON 数据
	projects := new([]model.Projects)
	if err := json.Unmarshal(data, &projects); err != nil {
		logger.Error("Projects JSON 数据解析报错:", err.Error())
		return nil, errors.New(fmt.Sprintf("Projects JSON 数据解析报错:", err.Error()))
	}
	return projects, nil
}
