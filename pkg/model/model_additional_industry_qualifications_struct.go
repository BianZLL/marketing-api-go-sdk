/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 附加行业资质信息
type AdditionalIndustryQualificationsStruct struct {
	QualificationId     int64               `json:"qualification_id,omitempty"`
	SystemIndustryId    int64               `json:"system_industry_id,omitempty"`
	BusinessScopeId     int64               `json:"business_scope_id,omitempty"`
	QualificationCode   string              `json:"qualification_code,omitempty"`
	ImageIdList         *[]string           `json:"image_id_list,omitempty"`
	QualificationStatus QualificationStatus `json:"qualification_status,omitempty"`
	ExpiredDate         string              `json:"expired_date,omitempty"`
	RejectMessage       string              `json:"reject_message,omitempty"`
	CreatedTime         int64               `json:"created_time,omitempty"`
	LastModifiedTime    int64               `json:"last_modified_time,omitempty"`
}
