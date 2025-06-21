package deployment_handler

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

// DeploymentResponse model
//
//	@Description	Информация о процессе деплоя
type DeploymentResponse struct {
	//	@Description	Уникальный ID деплоя
	//	@Example:		1
	Id types.UniqId `json:"id"`

	//	@Description	Идентификатор треда в Slack
	//	@Example:		1234567890.123456
	ThreadTs string `json:"thread_ts"`

	//	@Description	Номер Pull Request
	//	@Example:		42
	PullRequestNumber string `json:"pull_request_number"`

	//	@Description	ID Workflow Run в GitHub
	//	@Example:		987654321
	WorkflowRunId string `json:"workflow_run_id"`

	//	@Description	Статус деплоя
	//	@Example:		success
	//	@Enum			pending,success,failed,running
	Status string `json:"status"`
}
