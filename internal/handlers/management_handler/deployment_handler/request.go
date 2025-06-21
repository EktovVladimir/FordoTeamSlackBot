package deployment_handler

// CreateDeploymentRequest model
//
//	@Description	Запрос на создание деплоя
type CreateDeploymentRequest struct {
	//	@Description	Идентификатор треда в Slack
	//	@Example:		1234567890.123456
	//	@Required		true
	ThreadTs string `json:"thread_ts" validate:"required"`

	//	@Description	Номер Pull Request
	//	@Example:		42
	//	@Required		true
	PullRequestNumber string `json:"pull_request_number" validate:"required,numeric"`

	//	@Description	ID Workflow Run в GitHub
	//	@Example:		987654321
	WorkflowRunId string `json:"workflow_run_id" validate:"omitempty"`

	//	@Description	Статус деплоя
	//	@Example:		pending
	//	@Enum			pending,success,failed,running
	//	@Required		true
	Status string `json:"status" validate:"required,oneof=pending success failed running"`
}

// UpdateDeploymentRequest model
//
//	@Description	Запрос на обновление деплоя
type UpdateDeploymentRequest struct {
	//	@Description	Идентификатор треда в Slack
	//	@Example:		1234567890.123456
	ThreadTs *string `json:"thread_ts" validate:"omitempty,slack_ts"`

	//	@Description	Номер Pull Request
	//	@Example:		42
	PullRequestNumber *string `json:"pull_request_number" validate:"omitempty,numeric"`

	//	@Description	ID Workflow Run в GitHub
	//	@Example:		987654321
	WorkflowRunId *string `json:"workflow_run_id" validate:"omitempty"`

	//	@Description	Статус деплоя
	//	@Example:		success
	//	@Enum			pending,success,failed,running
	Status *string `json:"status" validate:"omitempty,oneof=pending success failed running"`
}
