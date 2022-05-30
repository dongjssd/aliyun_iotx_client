/**
    @author: dongjs
    @date: 2022/5/26
    @description: 查询网关的子设备列表 返回值
**/

package response

type DeviceSubdeviceQueryResponse struct {
	Id           string                   `json:"id,omitempty"`
	Code         int                      `json:"code,omitempty"`
	Message      string                   `json:"message,omitempty"`
	LocalizedMsg string                   `json:"localizedMsg,omitempty"`
	Data         deviceSubdeviceQueryData `json:"data,omitempty"`
}

type deviceSubdeviceQueryData struct {
	Total    int64                          `json:"total,omitempty"`
	PageNo   int                            `json:"pageNo,omitempty"`
	PageSize int                            `json:"pageSize,omitempty"`
	Items    []deviceSubdeviceQueryDataItem `json:"items,omitempty"`
}

type deviceSubdeviceQueryDataItem struct {
	IotId           string `json:"IotId,omitempty"`
	DeviceName      string `json:"deviceName,omitempty"`
	ProductKey      string `json:"productKey,omitempty"`
	ThingType       string `json:"thingType,omitempty"`
	FirmwareVersion string `json:"firmwareVersion,omitempty"`
	Status          int    `json:"status,omitempty"`
	NodeType        string `json:"nodeType,omitempty"`
	Region          string `json:"region,omitempty"`
	RbacTenantId    string `json:"rbacTenantId,omitempty"`
	GmtCreate       int64  `json:"gmtCreate,omitempty"`
	GmtModified     int64  `json:"gmtModified,omitempty"`
	SdkVersion      string `json:"sdkVersion,omitempty"`
}
