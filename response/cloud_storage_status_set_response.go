/**
    @author: dongjs
    @date: 2022/5/27
    @description: 设置用户的云存储套餐状态 返回值
**/

package response

type CloudStorageStatusSetResponse struct {
	Id           string                `json:"id,omitempty"`
	Code         int                   `json:"code,omitempty"`
	Message      string                `json:"message,omitempty"`
	LocalizedMsg string                `json:"localizedMsg,omitempty"`
	Data         cloudStorageOrderData `json:"data,omitempty"`
}
