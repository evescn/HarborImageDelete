package model

import "sync"

var Wg sync.WaitGroup

type Projects struct {
	Name      string `json:"name" form:"name" binding:"required"`
	ProjectId int    `json:"project_id"`
}

type ProjectsUrl struct {
	ProjectName string `json:"project_name" form:"project_name" binding:"required"`
	Total       int    `json:"total" form:"total"`
}

type Repositories struct {
	Name          string `json:"name" form:"name" binding:"required"`
	ArtifactCount int    `json:"artifact_count"`
	ProjectId     int    `json:"project_id"`
}

type ArtifactsUrl struct {
	ProjectName      string `json:"project_name" form:"project_name" binding:"required"`
	RepositoriesName string `json:"repositories_name" form:"repositories_name" binding:"required"`
	Total            int    `json:"total" form:"total"`
}

type ArtifactsTmp struct {
	Tags []Artifacts `json:"tags"`
}

type Artifacts struct {
	Name string `json:"name"`
}

type SystemGcSchedule struct {
	JobStatus string `json:"job_status"`
}
