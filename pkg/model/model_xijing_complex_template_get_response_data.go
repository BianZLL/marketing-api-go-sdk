/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type XijingComplexTemplateGetResponseData struct {
	PageTemplateId       *string           `json:"page_template_id,omitempty"`
	PageTemplateName     *string           `json:"page_template_name,omitempty"`
	PageTemplateCoverUrl *string           `json:"page_template_cover_url,omitempty"`
	PageName             *string           `json:"page_name,omitempty"`
	PageTitle            *string           `json:"page_title,omitempty"`
	PageConfig           *[]XjConfigStruct `json:"page_config,omitempty"`
}
