/**
    @author: dongjs
    @date: 2022/5/25
    @description: 创建量产批次接口返回参数
**/

package response

type DeviceNameUploadResponse struct {
	Id      string `json:"id,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}
