/**
    @author: dongjs
    @date: 2022/5/26
    @description: 查询云存储套餐是否可以购买 返回值
**/

package response

type CloudStorageCommodityCheckResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         struct {
		Available bool `json:"available,omitempty"`
	} `json:"data,omitempty"`
}
