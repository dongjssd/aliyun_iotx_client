/**
    @author: dongjs
    @date: 2022/5/26
    @description: 倒序分页查询用户列表 返回值
**/

package response

type AccountQueryIdentityByPageResponse struct {
	Id           string      `json:"id,omitempty"`
	Code         int         `json:"code,omitempty"`
	Message      string      `json:"message,omitempty"`
	LocalizedMsg string      `json:"localizedMsg,omitempty"`
	Data         accountData `json:"data,omitempty"`
}
