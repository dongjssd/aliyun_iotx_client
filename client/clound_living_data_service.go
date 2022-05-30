/**
    @author: dongjs
    @date: 2022/5/26
    @description: 运营中心API服务
**/

package client

import (
	"encoding/json"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/response"
)

/***********************************活跃设备信息服务********************************************/

//DeviceActiveSummaryGet 调用该接口获取活跃设备的概览信息
func (c *SyncApiClient) DeviceActiveSummaryGet(apiVer, projectId string, productKeys []string) (res response.DeviceActiveSummaryGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	req.Params.ProductKeys = productKeys
	if body, errMsg = c.PostBody("/living/data/device/activesummary/get", req); errMsg != nil {
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

// DeviceActiveDetailGet 调用该接口获取活跃设备的详细信息
func (c *SyncApiClient) DeviceActiveDetailGet(apiVer, projectId string, productKeys []string) (res response.DeviceActiveDetailGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	req.Params.ProductKeys = productKeys
	if body, errMsg = c.PostBody("/living/data/device/activedetail/get", req); errMsg != nil {
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

/***********************************活跃设备信息服务********************************************/

/***********************************激活设备信息服务********************************************/

// DeviceActivateDetailGet 调用该接口获取设备激活数据详细信息
func (c *SyncApiClient) DeviceActivateDetailGet(apiVer, projectId string, productKeys []string) (res response.DeviceActivateDetailGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	req.Params.ProductKeys = productKeys
	if body, errMsg = c.PostBody("/living/data/device/activatedetail/get", req); errMsg != nil {
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

// 调用该接口获取设备激活数据概览的信息

func (c *SyncApiClient) DeviceActivateSummaryGet(apiVer, projectId string, productKeys []string) (res response.DeviceActivateSummaryGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	req.Params.ProductKeys = productKeys
	if body, errMsg = c.PostBody("/living/data/device/activatesummary/get", req); errMsg != nil {
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

/***********************************激活设备信息服务********************************************/

/***********************************活跃用户信息服务********************************************/

//UserActiveSummaryGet 调用该接口获取活跃用户的概括信息
func (c *SyncApiClient) UserActiveSummaryGet(apiVer, projectId string) (res response.UserActiveSummaryGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	if body, errMsg = c.PostBody("/living/data/user/activesummary/get", req); errMsg != nil {
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

//UserActiveDetailGet 调用该接口获取活跃用户的详细信息
func (c *SyncApiClient) UserActiveDetailGet(apiVer, projectId string) (res response.UserActiveDetailGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	if body, errMsg = c.PostBody("/living/data/user/activedetail/get", req); errMsg != nil {
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

/***********************************活跃用户信息服务********************************************/

/***********************************注册用户信息服务********************************************/

//UserActivateSummaryGet 调用该接口获取注册用户概况信息
func (c *SyncApiClient) UserActivateSummaryGet(apiVer, projectId string) (res response.UserActivateSummaryGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	if body, errMsg = c.PostBody("/living/data/user/activatesummary/get", req); errMsg != nil {
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

//UserActivateDetailGet 调用该接口获取注册用户详细信息
func (c *SyncApiClient) UserActivateDetailGet(apiVer, projectId string) (res response.UserActivateDetailGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProjectId = projectId
	if body, errMsg = c.PostBody("/living/data/user/activatedetail/get", req); errMsg != nil {
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

/***********************************注册用户信息服务********************************************/
