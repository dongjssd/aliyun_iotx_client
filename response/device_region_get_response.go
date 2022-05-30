/**
    @author: dongjs
    @date: 2022/5/26
    @description: 获取设备当前所在的站点 返回值
**/

package response

type DeviceRegionGetResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         string `json:"data,omitempty"`
}
