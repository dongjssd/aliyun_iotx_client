/**
    @author: dongjs
    @date: 2022/5/27
    @description: 设置云存储套餐立即生效
**/

package response

type CloudStorageCommodityEnableResponse struct {
	Id           string                          `json:"id,omitempty"`
	Code         int                             `json:"code,omitempty"`
	Message      string                          `json:"message,omitempty"`
	LocalizedMsg string                          `json:"localizedMsg,omitempty"`
	Data         cloudStorageCommodityEnableData `json:"data,omitempty"`
}

type cloudStorageCommodityEnableData struct {
	OrderId       string `json:"orderId,omitempty"`
	CommodityCode string `json:"commodityCode,omitempty"`
	Specification string `json:"specification,omitempty"`
	Copies        int    `json:"copies,omitempty"`
	StartTime     string `json:"startTime,omitempty"`
	EndTime       string `json:"endTime,omitempty"`
	Expired       int    `json:"expired,omitempty"`
	PaymentStatus int    `json:"paymentStatus,omitempty"`
	Price         string `json:"price,omitempty"`
	CommodityType string `json:"commodityType,omitempty"`
	StartTimeUTC  string `json:"startTimeUTC,omitempty"`
	EndTimeUTC    string `json:"endTimeUTC,omitempty"`
}
