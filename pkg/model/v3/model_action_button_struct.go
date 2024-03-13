/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 行动按钮结构
type ActionButtonStruct struct {
	ButtonText *string         `json:"button_text,omitempty"`
	JumpInfo   *JumpinfoStruct `json:"jump_info,omitempty"`
}