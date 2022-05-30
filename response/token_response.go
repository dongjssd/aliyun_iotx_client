/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

// TokenResponse 返回data
type TokenResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         token  `json:"data,omitempty"`
}

type token struct {
	CloudToken string `json:"cloudToken,omitempty"` //云端Token
	ExpireIn   int64  `json:"expireIn,omitempty"`   //云端Token的有效期，单位：ms
}
