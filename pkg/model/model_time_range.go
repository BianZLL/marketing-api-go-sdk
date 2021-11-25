/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 时间范围，最长跨度1年。<br>time_type=TIME_TYPE_ACTION_TIME时填写线索提交时间，time_type=TIME_TYPE_CREATED_TIME时填写线索入库时间
type TimeRange struct {
	StartTime *int64 `json:"start_time,omitempty"`
	EndTime   *int64 `json:"end_time,omitempty"`
}
