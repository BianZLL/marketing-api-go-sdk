/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意规格和投放权限数据结构
type AdcreativeTemplateListStructAdpermit struct {
	AdcreativeTemplateId          *int64                `json:"adcreative_template_id,omitempty"`
	AdcreativeTemplateStyle       *string               `json:"adcreative_template_style,omitempty"`
	AdcreativeTemplateAppellation *string               `json:"adcreative_template_appellation,omitempty"`
	AdcreativeSampleImage         *string               `json:"adcreative_sample_image,omitempty"`
	PromotedObjectType            PromotedObjectType    `json:"promoted_object_type,omitempty"`
	SiteSet                       SiteSetDefinition     `json:"site_set,omitempty"`
	SupportBillingSpecList        *[]SupportBillingSpec `json:"support_billing_spec_list,omitempty"`
	SupportBidModeList            *[]string             `json:"support_bid_mode_list,omitempty"`
}