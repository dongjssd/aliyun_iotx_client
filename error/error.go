/**
    @author: dongjs
    @date: 2022/5/24
    @description:
**/

package error

const (
	AliYunAppKeyNilErrorMsg          = "云端唯一身份（AppKey）不能为空"
	AliYunAppSecretNilErrorMsg       = "云端密码（AppSecret）不能为空"
	AliYunApiRequestErrorCode        = 1000001
	AliYunApiDoErrorCode             = 1000002
	AliYunApiReadErrorCode           = 1000003
	AliYunApiResponseDecodeErrorCode = 1000004
)

type ErrorMsg struct {
	Code    int
	Message string
}
