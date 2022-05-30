/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

type ProductNvrDeviceCertQueryResponse struct {
	Id      string                        `json:"id,omitempty"`
	Code    int                           `json:"code,omitempty"`
	Message string                        `json:"message,omitempty"`
	Data    productNvrDeviceCertQueryData `json:"data,omitempty"`
}

type productNvrDeviceCertQueryData struct {
	Total             int64               `json:"total,omitempty"`
	PageNo            int                 `json:"pageNo,omitempty"`
	PageSize          int                 `json:"pageSize,omitempty"`
	NvrDeviceCertList []nvrDeviceCertList `json:"NvrDeviceCertList,omitempty"`
}

type nvrDeviceCertList struct {
	IpcNum            int                     `json:"ipcNum,omitempty"`
	NvrDeviceCertInfo []nvrDeviceCertInfoList `json:"nvrDeviceCertInfo"`
	SubDeviceCertList []nvrDeviceCertInfoList `json:"subDeviceCertList"`
}

type nvrDeviceCertInfoList struct {
	IotId        string `json:"iotId"`        //设备ID，生活物联网平台为设备颁发的ID，设备的唯一标识符
	ProductKey   string `json:"productKey"`   //产品的Key，设备证书信息之一。创建产品时，生活物联网平台为该产品颁发的全局唯一标识
	DeviceName   string `json:"deviceName"`   //设备的名称，设备证书信息之一。在注册设备时，自定义的或系统生成的设备名称，具备产品维度内的唯一性
	DeviceSecret string `json:"deviceSecret"` //设备的密钥，设备证书信息之一。在注册设备时生活物联网平台为设备颁发的设备密钥，用于认证加密。需与DeviceName成对使用
}
