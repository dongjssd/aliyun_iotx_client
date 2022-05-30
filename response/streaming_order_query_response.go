/**
    @author: dongjs
    @date: 2022/5/27
    @description: 查询直播分享套餐订单列表 返回值
**/

package response

type StreamingOrderQueryResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         []struct {
		PageNo    int                     `json:"pageNo,omitempty"`
		PageSize  int                     `json:"pageSize,omitempty"`
		PageCount int                     `json:"pageCount,omitempty"`
		Total     int                     `json:"total,omitempty"`
		OrderList []cloudStorageOrderData `json:"orderList,omitempty"`
	} `json:"data,omitempty"`
}
