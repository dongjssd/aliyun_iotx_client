/**
    @author: dongjs
    @date: 2022/5/27
    @description: 设备激活数据详细信息
**/

package response

type DeviceActivateDetailGetResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         struct {
		Data []deviceActiveData `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
