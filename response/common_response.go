/**
    @author: dongjs
    @date: 2022/5/26
    @description: 通用返回值
**/

package response

type CommonResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
}

// 账号详情信息
type accountData struct {
	LastLoginTime        int64  `json:"lastLoginTime,omitempty"`
	GmtModified          int64  `json:"gmtModified,omitempty"`
	ScopeId              string `json:"scopeId,omitempty"`
	CreateAppKey         string `json:"createAppKey,omitempty"`
	LoginId              string `json:"loginId,omitempty"`
	IdentityId           string `json:"identityId,omitempty"`
	NickName             string `json:"nickName,omitempty"`
	LoginName            string `json:"loginName,omitempty"`
	TenantId             string `json:"tenantId,omitempty"`
	GmtCreate            int64  `json:"gmtCreate,omitempty"`
	LoginSource          string `json:"loginSource,omitempty"`
	Status               string `json:"status,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Email                string `json:"email,omitempty"`
	AvatarUrl            string `json:"avatarUrl,omitempty"`
	SubIdentityLimitCode string `json:"subIdentityLimitCode,omitempty"`
}
