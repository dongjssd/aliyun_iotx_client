/**
    @author: dongjs
    @date: 2022/5/27
    @description:
**/

package response

type StreamingOrderGetResponse struct {
	Id           string                `json:"id,omitempty"`
	Code         int                   `json:"code,omitempty"`
	Message      string                `json:"message,omitempty"`
	LocalizedMsg string                `json:"localizedMsg,omitempty"`
	Data         cloudStorageOrderData `json:"data,omitempty"`
}
