/**
    @author: dongjs
    @date: 2022/5/27
    @description: 获取注册用户概况信息 返回值
**/

package response

type UserActivateSummaryGetResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         struct {
		Data deviceActiveData `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
