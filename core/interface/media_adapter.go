package _interface

import (
	"ad-exchange-server/core/model"
	"net/http"
)

// MediaAdapter 媒体适配器统一接口
type MediaAdapter interface {

	// UnmarshalRequest 从HTTP请求中解析并转换为内部统一请求模型
	UnmarshalRequest(r *http.Request) *model.AdInternalRequest

	// MarshalResponse 将内部统一响应模型转换为媒体对应的响应格式（返回字节流）
	MarshalResponse(internalResp *model.AdInternalResponse) ([]byte, error)

	// GetMediaName 获取适配器对应的媒体名称
	GetMediaName() string
}
