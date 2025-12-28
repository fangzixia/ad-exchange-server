package model

// AdInternalResponse 内部统一广告响应模型（所有媒体协议最终转换为此格式）
type AdInternalResponse struct {
	AdID        string
	AdTitle     string
	AdContent   string
	RedirectURL string
	Price       float64
	IsSuccess   bool
}
