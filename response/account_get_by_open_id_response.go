/**
    @author: dongjs
    @date: 2022/5/26
    @description: 通过三方外标查询账号信息 返回参数
**/

package response

type AccountGetByOpenIdResponse struct {
	Id           string      `json:"id,omitempty"`
	Code         int         `json:"code,omitempty"`
	Message      string      `json:"message,omitempty"`
	LocalizedMsg string      `json:"localizedMsg,omitempty"`
	Data         accountData `json:"data,omitempty"`
}
