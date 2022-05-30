/**
    @author: dongjs
    @date: 2022/5/26
    @description:
**/

package response

type DeviceCustomizemessageNotifyResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         struct {
		MessageId string      `json:"messageId,omitempty"`
		Data      interface{} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
