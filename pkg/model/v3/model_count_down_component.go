/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 倒计时组件
type CountDownComponent struct {
	ComponentId *int64           `json:"component_id,omitempty"`
	Value       *CountDownStruct `json:"value,omitempty"`
}