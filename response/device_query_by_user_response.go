/**
    @author: dongjs
    @date: 2022/5/26
    @description: 获取用户绑定的设备列表详情 返回参数
**/

package response

type DeviceQueryByUserResponse struct {
	Id           string                `json:"id,omitempty"`
	Code         int                   `json:"code,omitempty"`
	Message      string                `json:"message,omitempty"`
	LocalizedMsg string                `json:"localizedMsg,omitempty"`
	Data         deviceQueryByUserData `json:"data,omitempty"`
}

type deviceQueryByUserData struct {
	IdentifyId    string `json:"identifyId,omitempty"`
	IotId         string `json:"iotId,omitempty"`
	ProductKey    string `json:"productKey,omitempty"`
	DeviceName    string `json:"deviceName,omitempty"`
	ProductName   string `json:"productName,omitempty"`
	CategoryImage string `json:"categoryImage,omitempty"`
	ProductModel  string `json:"productModel,omitempty"`
	NickName      string `json:"nickName,omitempty"`
	NetType       string `json:"netType,omitempty"`
	ThingType     string `json:"thingType,omitempty"`
	Owned         int    `json:"owned,omitempty"`
}
