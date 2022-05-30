/**
    @author: dongjs
    @date: 2022/5/25
    @description:获取量产批次ID（NVR产品专用）返回值
**/

package response

type ProductNvrDeviceGenerateResponse struct {
	Id      string `json:"id,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}
