/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 表单组件
type FormComponent struct {
	ComponentId *int64      `json:"component_id,omitempty"`
	Value       *FormStruct `json:"value,omitempty"`
}