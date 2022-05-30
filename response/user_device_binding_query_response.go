/**
    @author: dongjs
    @date: 2022/5/26
    @description: 根据设备ID查找绑定的用户 返回参数
**/

package response

type UserDeviceBindingQueryResponse struct {
	Id           string                     `json:"id,omitempty"`
	Code         int                        `json:"code,omitempty"`
	Message      string                     `json:"message,omitempty"`
	LocalizedMsg string                     `json:"localizedMsg,omitempty"`
	Data         userDeviceBindingQueryData `json:"data,omitempty"`
}

type userDeviceBindingQueryData struct {
	Total          int                                    `json:"total"`
	AccountDevList []userDeviceBindingQueryDataAccountDev `json:"accountDevList,omitempty"`
}

type userDeviceBindingQueryDataAccountDev struct {
	IotId       string `json:"iotId,omitempty"`
	ProductKey  string `json:"productKey,omitempty"`
	DeviceName  string `json:"deviceName,omitempty"`
	ProductName string `json:"productName,omitempty"`
	NickName    string `json:"nickName,omitempty"`
	Owned       int    `json:"owned,omitempty"`
	IdentityId  string `json:"identityId,omitempty"`
}
