/**
    @author: dongjs
    @date: 2022/5/27
    @description:用户云存套餐转移
**/

package response

type CloudStorageTransferResponse struct {
	Id           string                          `json:"id,omitempty"`
	Code         int                             `json:"code,omitempty"`
	Message      string                          `json:"message,omitempty"`
	LocalizedMsg string                          `json:"localizedMsg,omitempty"`
	Data         cloudStorageCommodityEnableData `json:"data,omitempty"`
}
