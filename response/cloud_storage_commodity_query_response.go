/**
    @author: dongjs
    @date: 2022/5/26
    @description:查询云存储套餐列表 返回值
**/

package response

type CloudStorageCommodityQueryResponse struct {
	Id           string                         `json:"id,omitempty"`
	Code         int                            `json:"code,omitempty"`
	Message      string                         `json:"message,omitempty"`
	LocalizedMsg string                         `json:"localizedMsg,omitempty"`
	Data         cloudStorageCommodityQueryData `json:"data,omitempty"`
}

type cloudStorageCommodityQueryData struct {
	CommodityName string `json:"commodityName,omitempty"`
	CommodityCode string `json:"commodityCode,omitempty"`
	Specification string `json:"specification,omitempty"`
	RecordType    int    `json:"recordType,omitempty"`
	Price         string `json:"price,omitempty"`
	Lifecycle     int    `json:"lifecycle,omitempty"`
	Months        int    `json:"months,omitempty"`
}
