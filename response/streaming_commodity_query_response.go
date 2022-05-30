/**
    @author: dongjs
    @date: 2022/5/27
    @description:查询直播分享套餐列表
**/

package response

type StreamingCommodityQueryResponse struct {
	Id           string                      `json:"id,omitempty"`
	Code         int                         `json:"code,omitempty"`
	Message      string                      `json:"message,omitempty"`
	LocalizedMsg string                      `json:"localizedMsg,omitempty"`
	Data         streamingCommodityQueryData `json:"data,omitempty"`
}

type streamingCommodityQueryData struct {
	CommodityName string `json:"commodityName,omitempty"`
	CommodityCode string `json:"commodityCode,omitempty"`
	Price         string `json:"price,omitempty"`
	Quota         int    `json:"quota,omitempty"`
	Months        int    `json:"months,omitempty"`
}
