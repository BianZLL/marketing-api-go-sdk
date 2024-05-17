/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 任务结果信息
type MuseAiMaterialSaveResultItem struct {
	Status         MuseAiMaterialPushStatus `json:"status,omitempty"`
	MuseMaterialId *int64                   `json:"muse_material_id,omitempty"`
	MediaId        *string                  `json:"media_id,omitempty"`
	MaterialType   TemplateType             `json:"material_type,omitempty"`
}