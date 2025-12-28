package _interface

import "ad-exchange-server/core/model"

// PlatformAdapter 平台方适配器统一接口
type PlatformAdapter interface {

	// MarshalRequest 将内部平台方请求模型转换为平台方对应的协议格式（返回字节流）
	MarshalRequest(internalReq *model.AdInternalRequest) ([]byte, error)

	// UnmarshalResponse 将平台方响应字节流转换为内部统一响应模型
	UnmarshalResponse(respBytes []byte) (*model.AdInternalResponse, error)

	// GetPlatformName 获取适配器对应的平台方类型
	GetPlatformName() string

	// GetPlatformURL 获取平台方接口地址
	GetPlatformURL() string
}
