/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 场景定向
type SceneTargetingForWrite struct {
	MobileUnionIndustry         *[]string    `json:"mobile_union_industry,omitempty"`
	UnionPositionPackage        *[]int64     `json:"union_position_package,omitempty"`
	ExcludeUnionPositionPackage *[]int64     `json:"exclude_union_position_package,omitempty"`
	DisplayScene                *[]string    `json:"display_scene,omitempty"`
	MobileUnionCategory         *[]int64     `json:"mobile_union_category,omitempty"`
	TencentNews                 *[]string    `json:"tencent_news,omitempty"`
	WechatScene                 *WechatScene `json:"wechat_scene,omitempty"`
	WechatPosition              *[]int64     `json:"wechat_position,omitempty"`
}
