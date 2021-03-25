/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AutoAcquisitionStatus : 一键起量状态
type AutoAcquisitionStatus string

// List of AutoAcquisitionStatus
const (
	AutoAcquisitionStatus_ACQUISITION_STATUS_PENDING                 AutoAcquisitionStatus = "AUTO_ACQUISITION_STATUS_PENDING"
	AutoAcquisitionStatus_ACQUISITION_STATUS_COMPLETED               AutoAcquisitionStatus = "AUTO_ACQUISITION_STATUS_COMPLETED"
	AutoAcquisitionStatus_ACQUISITION_STATUS_END                     AutoAcquisitionStatus = "AUTO_ACQUISITION_STATUS_END"
	AutoAcquisitionStatus_ACQUISITION_STATUS_SUSPEND                 AutoAcquisitionStatus = "AUTO_ACQUISITION_STATUS_SUSPEND"
	AutoAcquisitionStatus_ACQUISTION_STATUS_UNKNOW                   AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_UNKNOW"
	AutoAcquisitionStatus_ACQUISTION_STATUS_PENDING                  AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_PENDING"
	AutoAcquisitionStatus_ACQUISTION_STATUS_END_LESS_THAN_24_H       AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_END_LESS_THAN_24H"
	AutoAcquisitionStatus_ACQUISTION_STATUS_END_MORE_THAN_24_H       AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_END_MORE_THAN_24H"
	AutoAcquisitionStatus_ACQUISTION_STATUS_COMPLETED                AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_COMPLETED"
	AutoAcquisitionStatus_ACQUISTION_STATUS_SUSPEND_ON_LEARNING_FAIL AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_SUSPEND_ON_LEARNING_FAIL"
	AutoAcquisitionStatus_ACQUISTION_STATUS_SUSPEND_ON_PLAYING_FAIL  AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_SUSPEND_ON_PLAYING_FAIL"
	AutoAcquisitionStatus_ACQUISTION_STATUS_ADVERTISER_CLOSED        AutoAcquisitionStatus = "AUTO_ACQUISTION_STATUS_ADVERTISER_CLOSED"
)