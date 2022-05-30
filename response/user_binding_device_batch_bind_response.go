/**
    @author: dongjs
    @date: 2022/5/26
    @description: 批量绑定设备 返回参数
**/

package response

type UserBindingDeviceBatchBindResponse struct {
	Id           string                           `json:"id,omitempty"`
	Code         int                              `json:"code,omitempty"`
	Message      string                           `json:"message,omitempty"`
	LocalizedMsg string                           `json:"localizedMsg,omitempty"`
	Data         []userBindingDeviceBatchBindData `json:"data,omitempty"`
}

type userBindingDeviceBatchBindData struct {
	ProductKey string `json:"productKey,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	IotId      string `json:"iotId,omitempty"`
	BindCode   string `json:"bindCode,omitempty"`
}
