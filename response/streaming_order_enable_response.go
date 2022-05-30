/**
    @author: dongjs
    @date: 2022/5/27
    @description:使直播分享套餐订单立即生效 返回值
**/

package response

type StreamingOrderEnable struct {
	Id           string                `json:"id,omitempty"`
	Code         int                   `json:"code,omitempty"`
	Message      string                `json:"message,omitempty"`
	LocalizedMsg string                `json:"localizedMsg,omitempty"`
	Data         cloudStorageOrderData `json:"data,omitempty"`
}
