/**
    @author: dongjs
    @date: 2022/5/26
    @description:获取物的基本信息 返回值
**/

package response

type ThingInfoGetResponse struct {
	Id      string           `json:"id,omitempty"`
	Code    int              `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Data    thingInfoGetData `json:"data,omitempty"`
}

type thingInfoGetData struct {
	GmtModified     int64  `json:"gmtModified,omitempty"`
	ActiveTime      int64  `json:"activeTime,omitempty"`
	GmtCreate       int64  `json:"gmtCreate,omitempty"`
	ProductKey      string `json:"productKey,omitempty"`
	StatusLast      int    `json:"statusLast,omitempty"`
	Mac             string `json:"mac,omitempty"`
	DeviceKey       string `json:"deviceKey,omitempty"`
	IotId           string `json:"iotId,omitempty"`
	Name            string `json:"name,omitempty"`
	Nickname        string `json:"nickname,omitempty"`
	SdkVersion      string `json:"sdkVersion,omitempty"`
	Sn              string `json:"sn,omitempty"`
	ThingType       string `json:"thingType,omitempty"`
	Region          string `json:"region,omitempty"`
	FirmwareVersion string `json:"firmwareVersion,omitempty"`
	RbacTenantId    string `json:"rbacTenantId,omitempty"`
	NetAddress      string `json:"netAddress,omitempty"`
	DeviceSecret    string `json:"deviceSecret,omitempty"`
	TenantId        string `json:"tenantId,omitempty"`
	Status          int    `json:"status,omitempty"`
}
