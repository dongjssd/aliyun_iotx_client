/**
    @author: dongjs
    @date: 2022/5/24
    @description: 云端管理基础服务
**/

package client

import (
	"encoding/json"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/response"
)

// TokenGet 调用该接口获取云端资源Token。该接口对于同一个res，在Token失效前仅需调用一次。如果再次调用，会生成新的cloudToken，并导致之前的cloudToken失效。
func (c *SyncApiClient) TokenGet(apiVer, projectId string) (res response.TokenResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	// 初始化请求参数
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.Res = projectId
	req.Params.GrantType = "project"
	if body, errMsg = c.PostBody("/cloud/token", req); errMsg != nil {
		res.Code = errMsg.Code
		res.Message = errMsg.Message
		return
	}
	if err = json.Unmarshal(body, &res); err != nil {
		res.Code = error2.AliYunApiResponseDecodeErrorCode
		res.Message = err.Error()
		return
	}
	return
}

// TokenRefresh 调用该接口刷新云端资源的Token。
func (c *SyncApiClient) TokenRefresh(apiVer, cloudToken string) (res response.TokenResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	// 初始化请求参数
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.CloudToken = cloudToken
	if body, errMsg = c.PostBody("/cloud/token/refresh", req); errMsg != nil {
		res.Code = errMsg.Code
		res.Message = errMsg.Message
		return
	}
	if err = json.Unmarshal(body, &res); err != nil {
		res.Code = error2.AliYunApiResponseDecodeErrorCode
		res.Message = err.Error()
		return
	}
	return
}
