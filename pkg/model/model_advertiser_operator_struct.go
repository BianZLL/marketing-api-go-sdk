/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 运营人员结构
type AdvertiserOperatorStruct struct {
	OperatorId      int64  `json:"operator_id,omitempty"`
	OperatorName    string `json:"operator_name,omitempty"`
	Qq              int64  `json:"qq,omitempty"`
	WechatAccountId string `json:"wechat_account_id,omitempty"`
	IsMaster        bool   `json:"is_master,omitempty"`
}
