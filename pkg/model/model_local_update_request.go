/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LocalUpdateRequest struct {
	AccountId  *int64                `json:"account_id,omitempty"`
	AdgroupId  *int64                `json:"adgroup_id,omitempty"`
	Adgroup    *AdgroupUpdateSpec    `json:"adgroup,omitempty"`
	Campaign   *CampaignUpdateSpec   `json:"campaign,omitempty"`
	Adcreative *AdCreativeUpdateSpec `json:"adcreative,omitempty"`
}