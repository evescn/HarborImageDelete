package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"harbor-image-delete/model"
)

func Artifacts(repositoriesUrl string) (r *[][]model.Artifacts, err error) {
	// 发起 HTTP 请求
	data, err := Get(repositoriesUrl)

	if err != nil {
		return nil, err
	}
	//fmt.Println(string(data))

	// 解析 JSON 数据
	artifacts := new([]model.ArtifactsTmp)
	if err := json.Unmarshal(data, &artifacts); err != nil {
		logger.Error("Artifacts JSON 数据解析报错:", err.Error())
		return nil, errors.New(fmt.Sprintf("Artifacts JSON 数据解析报错:", err.Error()))
	}

	artifactsData := new([]model.Artifacts)

	artifactsData.Name
	for _, tag := range artifacts.Tags {
		artifactsData = append(artifactsData, Tag{Name: tag.Name})
	}

	artifactsData = artifacts
	fmt.Println(artifacts)
	return nil, nil
}
