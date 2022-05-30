/**
    @author: dongjs
    @date: 2022/5/26
    @description:购买云存储套餐 返回值
**/

package response

type CloudStorageCommodityBuyResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         struct {
		OrderId bool `json:"orderId,omitempty"`
	} `json:"data,omitempty"`
}
