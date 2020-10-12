/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ActionChannelType : 渠道信息，标识该条行为发生在何渠道上。
type ActionChannelType string

// List of ActionChannelType
const (
	ActionChannelType_NATURAL   ActionChannelType = "NATURAL"
	ActionChannelType_TENCENT   ActionChannelType = "TENCENT"
	ActionChannelType_BYTEDANCE ActionChannelType = "BYTEDANCE"
	ActionChannelType_KUAISHOU  ActionChannelType = "KUAISHOU"
	ActionChannelType_ALIBABA   ActionChannelType = "ALIBABA"
	ActionChannelType_BAIDU     ActionChannelType = "BAIDU"
	ActionChannelType_OTHERS    ActionChannelType = "OTHERS"
	ActionChannelType_UNKNOWN   ActionChannelType = "UNKNOWN"
)
