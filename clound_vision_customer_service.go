/**
    @author: dongjs
    @date: 2022/5/26
    @description: 视频服务云存储购买
**/

package aliyun_iotx_client

import (
	"encoding/json"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/response"
)

// FreeCloudStorageGet 调用该接口查询免费的云存储套餐详情
func (c *SyncApiClient) FreeCloudStorageGet(apiVer, iotId string) (res response.FreeCloudStorageGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	if body, errMsg = c.PostBody("/vision/customer/freecloudstorage/get", req); errMsg != nil {
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

// FreeCloudStorageConsume 调用该接口领取免费的云存储套餐
func (c *SyncApiClient) FreeCloudStorageConsume(apiVer, iotId string, enableDefaultPlan, immediateUse bool,
	preRecordDuration, recordDuration, quota int, eventRecordProlong bool) (res response.FreeCloudStorageConsumeResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.EnableDefaultPlan = enableDefaultPlan
	req.Params.ImmediateUse = immediateUse
	req.Params.PreRecordDuration = preRecordDuration
	req.Params.RecordDuration = recordDuration
	req.Params.Quota = quota
	req.Params.EventRecordProlong = eventRecordProlong
	if body, errMsg = c.PostBody("/vision/customer/freecloudstorage/consume", req); errMsg != nil {
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

// CloudStorageCommodityQuery 调用该接口查询视频云存储套餐列表
func (c *SyncApiClient) CloudStorageCommodityQuery(apiVer string) (res response.CloudStorageCommodityQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/commodity/query", req); errMsg != nil {
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

// CloudStorageCommodityCheck 调用该接口查询云存储套餐是否可以购买
func (c *SyncApiClient) CloudStorageCommodityCheck(apiVer, iotId, commodityCode, specification string) (res response.CloudStorageCommodityCheckResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.CommodityCode = commodityCode
	req.Params.Specification = specification
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/commodity/check", req); errMsg != nil {
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

// CloudStorageCommodityBuy 调用该接口购买云存储套餐
func (c *SyncApiClient) CloudStorageCommodityBuy(apiVer, iotId, userName, commodityCode, specification string,
	copies int, enableDefaultPlan, immediateUse bool, preRecordDuration, recordDuration int,
	eventRecordProlong bool, maxRecordFileDuration int) (res response.CloudStorageCommodityBuyResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.CommodityCode = commodityCode
	req.Params.Specification = specification
	req.Params.EnableDefaultPlan = enableDefaultPlan
	req.Params.ImmediateUse = immediateUse
	req.Params.PreRecordDuration = preRecordDuration
	req.Params.RecordDuration = recordDuration
	req.Params.EventRecordProlong = eventRecordProlong
	req.Params.UserName = userName
	req.Params.Copies = copies
	req.Params.MaxRecordFileDuration = maxRecordFileDuration
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/commodity/buy", req); errMsg != nil {
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

// CloudStorageOrderGet 调用该接口获取云存储套餐的订单详情
func (c *SyncApiClient) CloudStorageOrderGet(apiVer, iotId, orderId string) (res response.CloudStorageOrderGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.OrderId = orderId
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/order/get", req); errMsg != nil {
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

// CloudStorageOrderQuery 调用该接口查询云存储套餐的订单列表信息
func (c *SyncApiClient) CloudStorageOrderQuery(apiVer, iotId string, pageNo, pageSize int) (res response.CloudStorageOrderQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.PageNo = pageNo
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/order/query", req); errMsg != nil {
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

// FreeCloudStorageEnable 调用该接口设置免费云存储套餐立即生效。
func (c *SyncApiClient) FreeCloudStorageEnable(apiVer, iotId string) (res response.FreeCloudStorageEnableResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	if body, errMsg = c.PostBody("/vision/customer/freecloudstorage/enable", req); errMsg != nil {
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

// CloudStorageCommodityEnable 调用该接口设置购买的云存储套餐立即生效
func (c *SyncApiClient) CloudStorageCommodityEnable(apiVer, iotId, orderId string) (res response.CloudStorageCommodityEnableResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.OrderId = orderId
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/commodity/enable", req); errMsg != nil {
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

// CloudStorageTransfer 调用该接口将用户已购买的云存储套餐转移。当前仅支持将云存储套餐从同一个管理员账号下的一个设备转移到另一个设备。
func (c *SyncApiClient) CloudStorageTransfer(apiVer, srcIotId, srcOrderId, dstIotId string, enableDefaultPlan, immediateUse bool) (res response.CloudStorageTransferResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.SrcIotId = srcIotId
	req.Params.SrcOrderId = srcOrderId
	req.Params.DstIotId = dstIotId
	req.Params.EnableDefaultPlan = enableDefaultPlan
	req.Params.ImmediateUse = immediateUse
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/transfer", req); errMsg != nil {
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

// CloudStorageStatusSet 设置用户的云存储套餐状态
func (c *SyncApiClient) CloudStorageStatusSet(apiVer, iotId, orderId string, status int) (res response.CloudStorageOrderGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.OrderId = orderId
	req.Params.Status = status
	if body, errMsg = c.PostBody("/vision/customer/cloudstorage/status/set", req); errMsg != nil {
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

// StreamingCommodityQuery 调用该接口查询直播分享套餐列表
func (c *SyncApiClient) StreamingCommodityQuery(apiVer string) (res response.StreamingCommodityQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	if body, errMsg = c.PostBody("/vision/customer/streaming/commodity/query", req); errMsg != nil {
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

// StreamingOrderCreate 调用该接口创建直播分享套餐订单
func (c *SyncApiClient) StreamingOrderCreate(apiVer, iotId, commodityCode, userName, userId string, copies int,
	immediateUse bool) (res response.StreamingOrderCreateResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.CommodityCode = commodityCode
	req.Params.UserName = userName
	req.Params.UserId = userId
	req.Params.ImmediateUse = immediateUse
	req.Params.Copies = copies
	if body, errMsg = c.PostBody("/vision/customer/streaming/order/create", req); errMsg != nil {
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

// StreamingOrderPay 调用该接口支付直播分享套餐订单。
func (c *SyncApiClient) StreamingOrderPay(apiVer, iotId, orderId string) (res response.StreamingOrderPayResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.OrderId = orderId
	if body, errMsg = c.PostBody("/vision/customer/streaming/order/pay", req); errMsg != nil {
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

// StreamingOrderGet 调用该接口查询直播分享套餐订单
func (c *SyncApiClient) StreamingOrderGet(apiVer, iotId, orderId string) (res response.StreamingOrderGetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.OrderId = orderId
	if body, errMsg = c.PostBody("/vision/customer/streaming/order/get", req); errMsg != nil {
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

// StreamingOrderQuery 调用该接口查询直播分享套餐订单列表。
func (c *SyncApiClient) StreamingOrderQuery(apiVer, iotId string, pageNo, pageSize int) (res response.StreamingOrderQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.PageNo = pageNo
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/vision/customer/streaming/order/query", req); errMsg != nil {
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

// StreamingOrderEnable 调用该接口使直播分享套餐订单立即生效
func (c *SyncApiClient) StreamingOrderEnable(apiVer, iotId, orderId string) (res response.StreamingOrderEnable) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.OrderId = orderId
	if body, errMsg = c.PostBody("/vision/customer/streaming/order/enable", req); errMsg != nil {
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

// FreeCloudStorageStatusSet 调用该接口设置用户的免费云存储套餐状态
func (c *SyncApiClient) FreeCloudStorageStatusSet(apiVer, iotId string, status int) (res response.FreeCloudStorageStatusSetResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	req.Params.Status = status
	if body, errMsg = c.PostBody("/vision/customer/freecloudstorage/status/set", req); errMsg != nil {
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

// FreeStorageQuotaCheck 调用该接口检查设备是否可以领用免费云存储。用户绑定设备正常使用过程中，只能领用一次；考虑到退换货的情形，一个设备最多支持被3个用户绑定领用免费云存。
func (c *SyncApiClient) FreeStorageQuotaCheck(apiVer, iotId string) (res response.CommonResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.IotId = iotId
	if body, errMsg = c.PostBody("/vision/customer/freestorage/quota/check", req); errMsg != nil {
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
