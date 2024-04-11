/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 推广资产信息
type DetailStruct struct {
	MarketingAssetId    *int64              `json:"marketing_asset_id,omitempty"`
	MarketingAssetName  *string             `json:"marketing_asset_name,omitempty"`
	MarketingAssetType  MarketingAssetType  `json:"marketing_asset_type,omitempty"`
	MarketingTargetType MarketingTargetType `json:"marketing_target_type,omitempty"`
	CreatedTime         *int64              `json:"created_time,omitempty"`
	Properties          *[]PropertyStruct   `json:"properties,omitempty"`
	ExtraProperties     *[]ExtraProperty    `json:"extra_properties,omitempty"`
}