/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductItemsDetailGetResponseData struct {
	FeedId          *int64           `json:"feed_id,omitempty"`
	SystemStatus    AuditStatus      `json:"system_status,omitempty"`
	RejectMessage   *string          `json:"reject_message,omitempty"`
	ProductItemSpec *ProductInfoSpec `json:"product_item_spec,omitempty"`
}