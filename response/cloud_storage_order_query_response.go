/**
    @author: dongjs
    @date: 2022/5/27
    @description: 查询云存储套餐的订单列表 返回值
**/

package response

type CloudStorageOrderQueryResponse struct {
	Id           string                     `json:"id,omitempty"`
	Code         int                        `json:"code,omitempty"`
	Message      string                     `json:"message,omitempty"`
	LocalizedMsg string                     `json:"localizedMsg,omitempty"`
	Data         cloudStorageOrderQueryData `json:"data,omitempty"`
}

type cloudStorageOrderQueryData struct {
	PageNo    int                     `json:"pageNo,omitempty"`
	PageSize  int                     `json:"pageSize,omitempty"`
	PageCount int                     `json:"pageCount,omitempty"`
	Total     int                     `json:"total,omitempty"`
	OrderList []cloudStorageOrderData `json:"order_list,omitempty"`
}
