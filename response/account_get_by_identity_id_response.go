/**
    @author: dongjs
    @date: 2022/5/26
    @description: 通过identityid查询账户的详细信息 返回值
**/

package response

type AccountGetByIdentityIdResponse struct {
	Id           string      `json:"id,omitempty"`
	Code         int         `json:"code,omitempty"`
	Message      string      `json:"message,omitempty"`
	LocalizedMsg string      `json:"localizedMsg,omitempty"`
	Data         accountData `json:"data,omitempty"`
}
