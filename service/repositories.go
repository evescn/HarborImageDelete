package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/dao"
	"harbor-image-delete/model"
	"strings"
)

func Repositories(params *model.Projects) (r *[]model.Repositories, err error) {
	// 发起 HTTP 请求
	repositoriesUrl := fmt.Sprintf("/api/v2.0/projects/%s/repositories?page_size=100", params.Name)
	data, err := dao.Get(repositoriesUrl)

	if err != nil {
		return nil, err
	}
	//logger.Info(string(data))

	// 解析 JSON 数据
	repositories := new([]model.Repositories)
	if err := json.Unmarshal(data, &repositories); err != nil {
		logger.Error("Repositories JSON 数据解析报错:", err.Error())
		return nil, errors.New(fmt.Sprintf("Repositories JSON 数据解析报错:", err.Error()))
	}

	// 创建 repositoriesData 切片并复制数据
	repositoriesData := make([]model.Repositories, len(*repositories))

	// 去掉 "NameSpace/" 前缀
	for i, repositorie := range *repositories {
		//logger.Info(tag.Tags[0].Name)
		repositoriesData[i] = model.Repositories{
			Name:          strings.Split(repositorie.Name, "/")[1],
			ArtifactCount: repositorie.ArtifactCount,
			ProjectId:     repositorie.ProjectId,
		}
	}

	return &repositoriesData, nil
}
