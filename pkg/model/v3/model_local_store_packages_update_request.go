/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LocalStorePackagesUpdateRequest struct {
	AccountId             *int64                  `json:"account_id,omitempty"`
	LocalStorePackageId   *int64                  `json:"local_store_package_id,omitempty"`
	LocalStorePackageName *string                 `json:"local_store_package_name,omitempty"`
	LocalStoreList        *[]LocalStoreListStruct `json:"local_store_list,omitempty"`
}