/**
    @author: dongjs
    @date: 2022/5/27
    @description: 创建直播分享套餐订单
**/

package response

type StreamingOrderCreateResponse struct {
	Id           string                `json:"id,omitempty"`
	Code         int                   `json:"code,omitempty"`
	Message      string                `json:"message,omitempty"`
	LocalizedMsg string                `json:"localizedMsg,omitempty"`
	Data         cloudStorageOrderData `json:"data,omitempty"`
}
