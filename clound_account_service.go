/**
    @author: dongjs
    @date: 2022/5/26
    @description:云端用户服务
**/

package aliyun_iotx_client

import (
	"encoding/json"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/response"
)

//AccountGetByOpenId 通过三方外标查询账号信息
func (c *SyncApiClient) AccountGetByOpenId(apiVer, openId, openIdAppKey string) (res response.AccountGetByOpenIdResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.OpenId = openId
	req.Params.OpenIdAppKey = openIdAppKey
	if body, errMsg = c.PostBody("/cloud/account/getByOpenId", req); errMsg != nil {
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

// DeviceQueryByUser 调用该接口获取用户绑定的设备列表详情，包括设备详情。通过设置不同的偏移量来查询任意部分的记录，暂不提供总条数查询。
func (c *SyncApiClient) DeviceQueryByUser(apiVer, openId, openIdAppKey, identifyId string,
	limit, offset int) (res response.DeviceQueryByUserResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.OpenId = openId
	req.Params.OpenIdAppKey = openIdAppKey
	req.Params.IdentityId = identifyId
	req.Params.Limit = limit
	req.Params.Offset = offset
	if body, errMsg = c.PostBody("/cloud/device/queryByUser", req); errMsg != nil {
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

// UserDeviceUnbind 解绑用户和设备
func (c *SyncApiClient) UserDeviceUnbind(apiVer, openId, openIdAppKey, iotId, identifyId,
	productKey, deviceName string) (res response.CommonResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.OpenId = openId
	req.Params.OpenIdAppKey = openIdAppKey
	req.Params.IdentityId = identifyId
	req.Params.IotId = iotId
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	if body, errMsg = c.PostBody("/cloud/user/device/unbind", req); errMsg != nil {
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

// UserBindingDeviceUnbind 调用该接口强制解除指定设备上所有绑定的关系
func (c *SyncApiClient) UserBindingDeviceUnbind(apiVer, iotId, productKey, deviceName string) (res response.CommonResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	if body, errMsg = c.PostBody("/living/cloud/user/binding/device/unbind", req); errMsg != nil {
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

// OpenIdGetByIdentifyId 根据身份ID获取第三方Openid
func (c *SyncApiClient) OpenIdGetByIdentifyId(apiVer, identityId, openIdAppKey string) (res response.CommonResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IdentityId = identityId
	req.Params.OpenIdAppKey = openIdAppKey
	if body, errMsg = c.PostBody("/cloud/account/openId/getByIdentityId", req); errMsg != nil {
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

// UserDeviceBindingQuery 调用该接口根据设备ID查询该设备上所有绑定的账户
func (c *SyncApiClient) UserDeviceBindingQuery(apiVer, iotId string, pageNo, pageSize int) (res response.UserDeviceBindingQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.PageSize = pageSize
	req.Params.PageNo = pageNo
	if body, errMsg = c.PostBody("/living/user/device/binding/query", req); errMsg != nil {
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

// AccountGetByIdentityId 调用该接口通过IdentityId查询账户的详细信息，无内容值的字段不返回。IdentityId可以通过调用/cloud/account/queryIdentityByPage接口获得。
func (c *SyncApiClient) AccountGetByIdentityId(apiVer, identityId string) (res response.AccountGetByIdentityIdResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IdentityId = identityId
	if body, errMsg = c.PostBody("/cloud/account/getByIdentityId", req); errMsg != nil {
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

// 调用该接口在云端更新自有账号系统在平台内的用户信息。

func (c *SyncApiClient) AccountInfoUpdate(apiVer, openIdAppKey, openId, identityId,
	nickName string) (res response.CommonResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IdentityId = identityId
	req.Params.OpenIdAppKey = openIdAppKey
	req.Params.OpenId = openId
	req.Params.NickName = nickName
	if body, errMsg = c.PostBody("/user/account/info/update", req); errMsg != nil {
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

// UserRegionGet 调用该接口获取用户注册所在的区域，这个区域注册之后不会变动，除非用户注销重新注册。
func (c *SyncApiClient) UserRegionGet(apiVer, phone, phoneLocationCode, email, openId string) (res response.CommonResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.OpenId = openId
	req.Params.Phone = phone
	req.Params.PhoneLocationCode = phoneLocationCode
	req.Params.Email = email
	if body, errMsg = c.PostBody("/living/cloud/user/region/get", req); errMsg != nil {
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

// AccountQueryIdentityDescByPage 调用该接口分页查询租户下的用户列表，倒序排序
func (c *SyncApiClient) AccountQueryIdentityDescByPage(apiVer string, offset, count int) (res response.AccountQueryIdentityByPageResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.Offset = offset
	req.Params.Count = count
	if body, errMsg = c.PostBody("/cloud/account/queryidentitydescbypage", req); errMsg != nil {
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

// AccountRegionGet 调用该接口获取用户账号对应的 Region 信息。支持手机、邮箱、国家、OAuthCode。（此接口云端与客户端通用）
func (c *SyncApiClient) AccountRegionGet(apiVer, queryType, countryCode, phone, email, authCode,
	phoneLocationCode string) (res response.AccountRegionGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.QueryType = queryType
	req.Params.CountryCode = countryCode
	req.Params.Phone = phone
	req.Params.Email = email
	req.Params.AuthCode = authCode
	req.Params.PhoneLocationCode = phoneLocationCode
	if body, errMsg = c.PostBody("/living/account/region/get", req); errMsg != nil {
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

// DeviceRegionDistribute 调用该接口批量把设备迁移到当前region。
func (c *SyncApiClient) DeviceRegionDistribute(apiVer, productKey string, deviceNames []string,
	bindingCheck, pushDeviceRegionSwitch bool) (res response.DeviceRegionDistributeResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceNames = deviceNames
	req.Params.BindingCheck = bindingCheck
	req.Params.PushDeviceRegionSwitch = pushDeviceRegionSwitch
	if body, errMsg = c.PostBody("/living/account/region/get", req); errMsg != nil {
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

// UserBindingDeviceBatchBind 批量绑定设备
func (c *SyncApiClient) UserBindingDeviceBatchBind(apiVer, openId, openIdAppKey, identityId,
	projectId string, deviceList []map[string]string) (res response.UserBindingDeviceBatchBindResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.OpenId = openId
	req.Params.OpenIdAppKey = openIdAppKey
	req.Params.IdentityId = identityId
	req.Params.ProjectId = projectId
	req.Params.DeviceList = deviceList
	if body, errMsg = c.PostBody("/living/cloud/user/binding/device/batch/bind", req); errMsg != nil {
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
