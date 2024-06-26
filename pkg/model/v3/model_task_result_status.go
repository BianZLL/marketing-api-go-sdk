/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// TaskResultStatus : 任务结果状态
type TaskResultStatus string

// List of TaskResultStatus
const (
	TaskResultStatus_UNKNOWN      TaskResultStatus = "TASK_RESULT_STATUS_UNKNOWN"
	TaskResultStatus_PENDING      TaskResultStatus = "TASK_RESULT_STATUS_PENDING"
	TaskResultStatus_PROCESSING   TaskResultStatus = "TASK_RESULT_STATUS_PROCESSING"
	TaskResultStatus_SUCCESS      TaskResultStatus = "TASK_RESULT_STATUS_SUCCESS"
	TaskResultStatus_FAIL         TaskResultStatus = "TASK_RESULT_STATUS_FAIL"
	TaskResultStatus_PARTIAL_FAIL TaskResultStatus = "TASK_RESULT_STATUS_PARTIAL_FAIL"
	TaskResultStatus_SYSTEM_ERROR TaskResultStatus = "TASK_RESULT_STATUS_SYSTEM_ERROR"
	TaskResultStatus_DELETED      TaskResultStatus = "TASK_RESULT_STATUS_DELETED"
)
