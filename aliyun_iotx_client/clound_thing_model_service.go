/**
    @author: dongjs
    @date: 2022/5/25
    @description:云端物的模型服务
**/

package aliyun_iotx_client

import (
	"encoding/json"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/response"
)

// ThingPropertiesGet 调用该接口获取物的所有属性快照数据
func (c *SyncApiClient) ThingPropertiesGet(apiVer, iotId, productKey, deviceName string) (res response.ThingPropertiesGetResponse) {
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
	if body, errMsg = c.PostBody("/cloud/thing/properties/get", req); errMsg != nil {
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

// ThingTslGet 调用该接口获取物的TSL功能模板，包含属性、事件、服务的定义
func (c *SyncApiClient) ThingTslGet(apiVer, iotId, productKey, deviceName string) (res map[string]interface{}) {
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
	if body, errMsg = c.PostBody("/cloud/thing/tsl/get", req); errMsg != nil {
		res["code"] = errMsg.Code
		res["message"] = errMsg.Message
		return
	}
	if err = json.Unmarshal(body, &res); err != nil {
		res["code"] = error2.AliYunApiResponseDecodeErrorCode
		res["message"] = err.Error()
		return
	}
	return
}

// ThingServiceInvoke 调用该接口触发物的服务
func (c *SyncApiClient) ThingServiceInvoke(apiVer, iotId, productKey, deviceName, identifier string, args map[string]interface{}) (res response.CommonResponse) {
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
	req.Params.Identifier = identifier
	req.Params.Args = args
	if body, errMsg = c.PostBody("/cloud/thing/service/invoke", req); errMsg != nil {
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

// ThingPropertiesSet 设置物的属性
func (c *SyncApiClient) ThingPropertiesSet(apiVer, iotId, productKey, deviceName string, items map[string]interface{}) (res response.CommonResponse) {
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
	req.Params.Items = items
	if body, errMsg = c.PostBody("/cloud/thing/properties/set", req); errMsg != nil {
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

// ThingStatusGet 调用该接口获取物的连接状态
func (c *SyncApiClient) ThingStatusGet(apiVer, iotId, productKey, deviceName string) (res response.ThingStatusGetResponse) {
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
	if body, errMsg = c.PostBody("/cloud/thing/status/get", req); errMsg != nil {
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

// ThingInfoGet 获取物的基本信息
func (c *SyncApiClient) ThingInfoGet(apiVer, iotId, productKey, deviceName string) (res response.ThingInfoGetResponse) {
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
	if body, errMsg = c.PostBody("/cloud/thing/info/get", req); errMsg != nil {
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

// ThingsGet 批量获取物的列表
func (c *SyncApiClient) ThingsGet(apiVer, productKey string, status, currentPage, pageSize int) (res response.ThingsGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.Status = status
	req.Params.CurrentPage = currentPage
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/cloud/things/get", req); errMsg != nil {
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

// ThingEventTimelineGet 获取物的事件timeline数据
func (c *SyncApiClient) ThingEventTimelineGet(apiVer, productKey, deviceName, iotId, identifier, eventType string,
	start, end int64, pageSize int, ordered bool) (res response.ThingEventTimelineGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	req.Params.IotId = iotId
	req.Params.Identifier = identifier
	req.Params.EventType = eventType
	req.Params.Start = start
	req.Params.End = end
	req.Params.PageSize = pageSize
	req.Params.Ordered = ordered
	if body, errMsg = c.PostBody("/cloud/thing/event/timeline/get", req); errMsg != nil {
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

// ThingPropertyTimelineGet 获取物的属性timeline数据
func (c *SyncApiClient) ThingPropertyTimelineGet(apiVer, productKey, deviceName, iotId, identifier string,
	start, end int64, pageSize int, ordered bool) (res response.ThingPropertyTimelineGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	req.Params.IotId = iotId
	req.Params.Identifier = identifier
	req.Params.Start = start
	req.Params.End = end
	req.Params.PageSize = pageSize
	req.Params.Ordered = ordered
	if body, errMsg = c.PostBody("/cloud/thing/property/timeline/get", req); errMsg != nil {
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

// DeviceSubdeviceQuery 调用该接口提供根据网关设备ID查询子设备列表。
func (c *SyncApiClient) DeviceSubdeviceQuery(apiVer, productKey, deviceName, iotId string, pageNo, pageSize int) (res response.DeviceSubdeviceQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	req.Params.IotId = iotId
	req.Params.PageSize = pageSize
	req.Params.PageNo = pageNo
	if body, errMsg = c.PostBody("/living/cloud/device/subdevice/query", req); errMsg != nil {
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

// DeviceRegionGet 获取设备当前所在的站点
func (c *SyncApiClient) DeviceRegionGet(apiVer, productKey, deviceName string) (res response.DeviceRegionGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	if body, errMsg = c.PostBody("/living/cloud/device/region/get", req); errMsg != nil {
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

// DeviceCustomizedmessageNotify 向指定设备发送自定义消息
func (c *SyncApiClient) DeviceCustomizedmessageNotify(apiVer, productKey, deviceName, iotId string,
	messageContent map[string]interface{}) (res response.DeviceCustomizemessageNotifyResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceName = deviceName
	req.Params.MessageContent = messageContent
	req.Params.IotId = iotId
	if body, errMsg = c.PostBody("/living/cloud/device/customizedmessage/notify", req); errMsg != nil {
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
