package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"harbor-image-delete/dao"
	"harbor-image-delete/model"
	"strings"
)

// DeleteFromProjectsAndRepositories 基于 Projects 和 Repositories 删除对应 Repositories 下多余镜像，默认保留最后 20 次
func DeleteFromProjectsAndRepositories(params *model.ArtifactsUrl, artifactsData *[]model.Artifacts) (data []model.Artifacts, err error) {
	// 获取需要删除 Image 数据，循环并发删除
	deleteData := (*artifactsData)[params.Total:]

	// 循环高并发删除数据
	errorsCh := make(chan error)
	for _, imageTag := range deleteData {
		// logger.Info(imageTag.Name)
		model.Wg.Add(1)
		// 拼接 URL 地址 /api/v2.0/projects/t1/repositories/ltask/artifacts/t1-20230601-1440
		artifactsImagesUrl := fmt.Sprintf("/api/v2.0/projects/%s/repositories/%s/artifacts/%s", params.ProjectName, params.RepositoriesName, imageTag.Name)
		// 发起 HTTP 请求
		go func() {
			_, err = dao.Delete(artifactsImagesUrl)
			if err != nil {
				errorsCh <- err
			}
		}()
	}
	model.Wg.Wait()
	close(errorsCh)

	var allErrors []string
	for err = range errorsCh {
		allErrors = append(allErrors, err.Error())
	}

	if len(allErrors) > 0 {
		combinedErrors := strings.Join(allErrors, "; ")
		logger.Error("All Errors: %s\n", combinedErrors)
		return nil, errors.New(fmt.Sprintf("All Errors: %s\n", combinedErrors))
	}

	//logger.Info(string(data))
	return deleteData, nil
}

// DeleteFromProjects 基于 Projects 删除对应 Projects 所有 Repositories 多余镜像，默认保留最后 20 次
func DeleteFromProjects(params *model.ProjectsUrl, repositoriesData *[]model.Repositories) (data []model.ArtifactsUrl, err error) {
	// 处理数据，获取 Image 数多余 params.Total 的 Repositories
	tmpData := make([]model.Repositories, 0)
	for _, artifactCount := range *repositoriesData {
		//logger.Info(artifactCount.ArtifactCount)
		if artifactCount.ArtifactCount > params.Total {
			tmpData = append(tmpData, artifactCount)
		}
	}

	logger.Info(tmpData)

	// 循环 查询 Repositories 下的 Artifacts，并使用高并发删除数据
	errorsCh := make(chan error)
	deleteDataAll := make([]model.ArtifactsUrl, len(tmpData))
	for i, repositoriesName := range tmpData {
		model.WgProjects.Add(1)
		go func() {
			new_params := &model.ArtifactsUrl{
				ProjectName:      params.ProjectName,
				RepositoriesName: repositoriesName.Name,
				Total:            params.Total,
			}

			// 调用 service 方法，获取 Artifacts 列表
			artifactsData, err := Artifacts(new_params)
			if err != nil {
				errorsCh <- err
			}
			logger.Info(artifactsData)

			// 调用 service 方法，删除 Image 数据
			deleteData, err := DeleteFromProjectsAndRepositories(new_params, artifactsData)
			if err != nil {
				errorsCh <- err
			}

			// 统计需要删除 Image 总数
			total := len(deleteData)
			deleteDataAll[i] = model.ArtifactsUrl{
				ProjectName:      params.ProjectName,
				RepositoriesName: repositoriesName.Name,
				Total:            total,
			}
			model.WgProjects.Done()
		}()
	}
	model.WgProjects.Wait()
	close(errorsCh)

	var allErrors []string
	for err = range errorsCh {
		allErrors = append(allErrors, err.Error())
	}

	if len(allErrors) > 0 {
		combinedErrors := strings.Join(allErrors, "; ")
		logger.Error("All Errors: %s\n", combinedErrors)
		return nil, errors.New(fmt.Sprintf("All Errors: %s\n", combinedErrors))
	}

	return deleteDataAll, nil
}

// DeleteALL 删除所有 Projects 下 Repositories 多余镜像，默认保留最后 20 次
func DeleteALL(c *gin.Context) {
	return
}
