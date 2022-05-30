/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

type NameDeviceGenerateResponse struct {
	Id      string `json:"id,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		BatchId string `json:"batchId,omitempty"`
	} `json:"data,omitempty"`
}
