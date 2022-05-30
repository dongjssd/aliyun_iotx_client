/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

type DeviceApplyQueryResponse struct {
	Id      string               `json:"id,omitempty"`
	Code    int                  `json:"code,omitempty"`
	Message string               `json:"message,omitempty"`
	Data    deviceApplyQueryData `json:"data,omitempty"`
}

type deviceApplyQueryData struct {
	TotalNum int64                  `json:"totalNum,omitempty"`
	PageNo   int                    `json:"pageNo,omitempty"`
	PageSize int                    `json:"pageSize,omitempty"`
	Items    []deviceApplyQueryItem `json:"items,omitempty"`
}

type deviceApplyQueryItem struct {
	IotId        string `json:"iotId,omitempty"`
	ProductKey   string `json:"productKey,omitempty"`
	DeviceName   string `json:"deviceName,omitempty"`
	DeviceSecret string `json:"deviceSecret,omitempty"`
	ActiveTime   string `json:"activeTime,omitempty"`
	Status       int    `json:"status,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
}
