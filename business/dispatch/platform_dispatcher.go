package dispatch

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/core/model"
	"ad-exchange-server/infra/httpclient"
	"context"
	"log"
	"sync"
	"time"
)

// PlatformDispatcher 平台方分发器
type PlatformDispatcher struct {
	httpClient *httpclient.HTTPClient
	timeout    time.Duration
}

// NewPlatformDispatcher 创建平台方分发器实例
func NewPlatformDispatcher() *PlatformDispatcher {
	return &PlatformDispatcher{
		httpClient: httpclient.NewHTTPClient(),
		timeout:    3 * time.Second, // 3秒超时
	}
}

// Dispatch 并发分发请求给所有平台方
func (d *PlatformDispatcher) Dispatch(c *model.AdPlatformContent, adapters []_interface.PlatformAdapter) map[string]*model.AdInternalResponse {
	var (
		wg          sync.WaitGroup
		respChan    = make(chan *model.AdInternalResponse, len(adapters))
		doneChan    = make(chan struct{})
		ctx, cancel = context.WithTimeout(context.Background(), d.timeout)
	)
	defer cancel()
	defer close(respChan)
	log.Println("开始处理转发")
	// 并发请求每个平台方
	for _, adapter := range adapters {
		wg.Add(1)
		go func(adapter _interface.PlatformAdapter) {
			defer func() {
				log.Println("平台 defer")
				defer wg.Done()
			}()
			// 1. 内部请求 -> 平台方协议请求
			reqBytes, err := adapter.MarshalRequest(c)
			if err != nil {
				log.Printf("平台方[%s]请求序列化失败: %v", adapter.GetPlatformName(), err)
				return
			}

			// 2. 发送HTTP请求
			timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelFunc()
			respBytes, err := d.httpClient.Post(timeout, adapter.GetPlatformURL(), "application/json", reqBytes)
			if err != nil {
				log.Printf("平台方[%s]HTTP请求失败: %v", adapter.GetPlatformName(), err)
				return
			}

			// 3. 平台方响应 -> 内部统一响应
			internalResp, err := adapter.UnmarshalResponse(c, respBytes)
			if err != nil {
				log.Printf("平台方[%s]响应反序列化失败: %v", adapter.GetPlatformName(), err)
				return
			}

			if internalResp.AdInfos != nil && len(internalResp.AdInfos) > 0 {
				log.Println("平台序列化完成")
				respChan <- internalResp
			}
		}(adapter)
	}
	// 等待所有goroutine结束
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	// 收集响应
	var platformResponses = make(map[string]*model.AdInternalResponse)
	select {
	case <-ctx.Done():
	case <-doneChan:
	}
	log.Println("通道处理完成")
	// 读取通道中所有响应
	for {
		select {
		case resp := <-respChan:
			log.Println("处理完成")
			platformResponses["titanvol"] = resp
		default:
			return platformResponses
		}
	}

}
