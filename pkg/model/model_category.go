/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 类目信息
type Category struct {
	Level OcpaCategoryLevel `json:"level,omitempty"`
	Id    int64             `json:"id,omitempty"`
	Name  string            `json:"name,omitempty"`
}
