/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatAdvertiserLocalBusinessGetResponseData struct {
	HeadImageUrl        string `json:"head_image_url,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	ContactPerson       string `json:"contact_person,omitempty"`
	ContactPersonMobile string `json:"contact_person_mobile,omitempty"`
	ContactPersonCardId string `json:"contact_person_card_id,omitempty"`
	ContactPersonTele   string `json:"contact_person_tele,omitempty"`
	Corporation         string `json:"corporation,omitempty"`
	CorporationLicence  string `json:"corporation_licence,omitempty"`
	BusinessContent     string `json:"business_content,omitempty"`
	IndustryId          int64  `json:"industry_id,omitempty"`
	AccountId           int64  `json:"account_id,omitempty"`
	BusinessId          string `json:"business_id,omitempty"`
}
