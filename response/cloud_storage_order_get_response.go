/**
    @author: dongjs
    @date: 2022/5/26
    @description: 获取云存储套餐的订单详情 返回值
**/

package response

type CloudStorageOrderGetResponse struct {
	Id           string                `json:"id,omitempty"`
	Code         int                   `json:"code,omitempty"`
	Message      string                `json:"message,omitempty"`
	LocalizedMsg string                `json:"localizedMsg,omitempty"`
	Data         cloudStorageOrderData `json:"data,omitempty"`
}

type cloudStorageOrderData struct {
	IotId         string `json:"iotId,omitempty"`
	IdentityId    string `json:"identityId,omitempty"`
	CommodityCode string `json:"commodityCode,omitempty"`
	Specification string `json:"specification,omitempty"`
	Copies        int    `json:"copies,omitempty"`
	Price         string `json:"price,omitempty"`
	UserName      string `json:"userName,omitempty"`
	UserId        string `json:"userId,omitempty"`
	OrderId       string `json:"orderId,omitempty'"`
	OutOrderNo    string `json:"outOrderNo,omitempty"`
	Status        int    `json:"status,omitempty"`
	PaymentStatus int    `json:"paymentStatus,omitempty"`
	RecordType    int    `json:"recordType,omitempty"`
	OrderType     int    `json:"orderType,omitempty"`
	StartTime     string `json:"startTime,omitempty"`
	EndTime       string `json:"endTime,omitempty"`
	StartTimeUTC  string `json:"startTimeUTC,omitempty"`
	EndTimeUTC    string `json:"endTimeUTC,omitempty"`
}
