/**
    @author: dongjs
    @date: 2022/5/26
    @description:批量把设备迁移到当前region 返回值
**/

package response

type DeviceRegionDistributeResponse struct {
	Id           string                       `json:"id,omitempty"`
	Code         int                          `json:"code,omitempty"`
	Message      string                       `json:"message,omitempty"`
	LocalizedMsg string                       `json:"localizedMsg,omitempty"`
	Data         []deviceRegionDistributeData `json:"data,omitempty"`
}

type deviceRegionDistributeData struct {
	ProductKey string `json:"productKey"`
	DeviceName string `json:"deviceName"`
	Succeed    bool   `json:"succeed"`
	ResultCode string `json:"resultCode"`
}
