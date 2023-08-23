# HarborImageDelete

> harbor镜像仓库清理服务

## 服务接口

地址 | 类型 | 接口参数 | 用途 
-- | -- | -- | --
/Projects | GET | | 获取当前 Harbor 中的所有项目
/Repositories | GET | name=${ProjectName} | 获取 Harbor 选中项目的 Repositories 信息
/Artifacts | GET | project_name=${ProjectName}&repositories_name=${RepositoriesName} | 获取 Harbor 选中 Repositories 的 Artifacts 信息
/DeleteFromProjectsAndRepositories | DELETE | { "project_name": "${ProjectName}", "repositories_name": "${RepositoriesName}" } | 基于 Harbor 选中项目和 Repositories 删除多余 Artifacts
/DeleteFromProjects | DELETE | { "project_name": "${ProjectName}" } | 基于 Harbor 选中项目，删除项目内所有 Repositories 的多余 Artifacts
/SystemGcSchedule | POST | 创建清理 Harbor 主机磁盘空间任务，上述删除镜像只是逻辑删除，物理删除需要调用此接口执行磁盘删除

> DeleteFromProjectsAndRepositories 和 DeleteFromProjects 接口最后会自动调用一次 SystemGcSchedule 接口

