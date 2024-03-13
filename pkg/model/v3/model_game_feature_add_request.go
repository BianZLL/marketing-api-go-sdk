/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type GameFeatureAddRequest struct {
	AccountId                  *int64                          `json:"account_id,omitempty"`
	MarketingTargetType        MarketingTargetType             `json:"marketing_target_type,omitempty"`
	MarketingTargetDetailId    *string                         `json:"marketing_target_detail_id,omitempty"`
	ProfitModeId               *int64                          `json:"profit_mode_id,omitempty"`
	GameTypeId                 *int64                          `json:"game_type_id,omitempty"`
	GameplayIdList             *[]int64                        `json:"gameplay_id_list,omitempty"`
	GameThemeId                *int64                          `json:"game_theme_id,omitempty"`
	GameSubThemeId             *int64                          `json:"game_sub_theme_id,omitempty"`
	GameMarketingLifecycleList *[]GameMarketingLifecycleStruct `json:"game_marketing_lifecycle_list,omitempty"`
}