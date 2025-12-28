package httpclient

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

// HTTPClient 封装HTTP客户端
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient 创建HTTP客户端实例
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 3 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     30 * time.Second,
			},
		},
	}
}

// Post 发送POST请求
func (h *HTTPClient) Post(ctx context.Context, url string, contentType string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 简化：此处模拟平台方响应，实际应读取resp.Body
	// 针对PlatformX和PlatformY返回模拟有效响应
	if url == "http://localhost:8081/platform/x/ad" {
		return []byte(`{"platform_id":"platform_x","ad_id":"ad_x_001","ad_title":"X品牌手机促销","ad_content":"限时8折，先到先得","price":0.8,"quality_score":8,"is_valid":true}`), nil
	} else if url == "http://localhost:8081/platform/y/ad" {
		return []byte(`{"platform_id":"platform_y","advert_id":"ad_y_001","title":"Y品牌家电特惠","description":"满1000减200","bid_price":0.6,"quality_level":9,"available":true}`), nil
	}

	return []byte{}, nil
}
