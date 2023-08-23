package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/dao"
	"harbor-image-delete/model"
)

func Artifacts(params *model.ArtifactsUrl) (a *[]model.Artifacts, err error) {
	// 发起 HTTP 请求
	artifactsUrl := fmt.Sprintf("/api/v2.0/projects/%s/repositories/%s/artifacts?page_size=100", params.ProjectName, params.RepositoriesName)
	data, err := dao.Get(artifactsUrl)

	if err != nil {
		return nil, err
	}
	//logger.Info(string(data))

	// 解析 JSON 数据
	artifacts := new([]model.ArtifactsTmp)
	if err = json.Unmarshal(data, &artifacts); err != nil {
		logger.Error("Artifacts JSON 数据解析报错:", err.Error())
		return nil, errors.New(fmt.Sprintf("Artifacts JSON 数据解析报错:", err.Error()))
	}

	// 创建 artifactsData 切片并复制数据
	artifactsData := make([]model.Artifacts, len(*artifacts))

	// 转换数据格式为 artifactsData 数据格式，并返回 controller 层
	for i, tag := range *artifacts {
		//logger.Info(tag.Tags[0].Name)
		artifactsData[i] = model.Artifacts{Name: tag.Tags[0].Name}
	}
	return &artifactsData, nil
}
