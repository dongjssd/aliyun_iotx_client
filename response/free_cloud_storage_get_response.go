/**
    @author: dongjs
    @date: 2022/5/26
    @description: 查询免费的云存储套餐详情 返回值
**/

package response

type FreeCloudStorageGetResponse struct {
	Id           string               `json:"id,omitempty"`
	Code         int                  `json:"code,omitempty"`
	Message      string               `json:"message,omitempty"`
	LocalizedMsg string               `json:"localizedMsg,omitempty"`
	Data         freeCloudStorageData `json:"data,omitempty"`
}
