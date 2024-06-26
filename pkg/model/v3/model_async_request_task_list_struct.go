/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type AsyncRequestTaskListStruct struct {
	TaskId               *int64           `json:"task_id,omitempty"`
	TaskName             *string          `json:"task_name,omitempty"`
	TaskType             *string          `json:"task_type,omitempty"`
	TaskScope            TaskScope        `json:"task_scope,omitempty"`
	Status               TaskStatus       `json:"status,omitempty"`
	ResultStatus         TaskResultStatus `json:"result_status,omitempty"`
	CreatedTime          *int64           `json:"created_time,omitempty"`
	EndTime              *int64           `json:"end_time,omitempty"`
	ScopeObjectIdList    *[]int64         `json:"scope_object_id_list,omitempty"`
	ScopeObjectIdStrList *[]string        `json:"scope_object_id_str_list,omitempty"`
}
