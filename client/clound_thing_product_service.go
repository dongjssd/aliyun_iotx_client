/**
    @author: dongjs
    @date: 2022/5/25
    @description:云端物的产品管理
**/

package client

import (
	"encoding/json"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/response"
)

// CloudThingProductList 调用该接口查询当前项目下的产品列表
func (c *SyncApiClient) CloudThingProductList(apiVer string, pageNo, pageSize int) (res response.ProductListResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.PageNo = pageNo
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/cloud/thing/productList/get", req); errMsg != nil {
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

// CloudThingProduct  调用该接口查询指定产品的详细信息
func (c *SyncApiClient) CloudThingProduct(apiVer, productKey string) (res response.ProductResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	if body, errMsg = c.PostBody("/cloud/thing/product/get", req); errMsg != nil {
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

// AmountDeviceGenerate 云端开放的产品量产接口（动态生成设备名称）调用该接口可自动生成设备证书，设备名称由系统分配。生成设备证书后，可以直接调用/living/device/apply/query接口下载设备证书。
func (c *SyncApiClient) AmountDeviceGenerate(apiVer, productKey string, amount int64) (res response.DeviceGenerateResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.Amount = amount
	if body, errMsg = c.PostBody("/cloud/amount/device/generate", req); errMsg != nil {
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

// DeviceNameUpload 创建量产批次接口（用户上传设备名称方式）
func (c *SyncApiClient) DeviceNameUpload(apiVer, productKey string, deviceNames []string) (res response.DeviceNameUploadResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.DeviceNames = deviceNames
	if body, errMsg = c.PostBody("/cloud/device/name/upload", req); errMsg != nil {
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

// NameDeviceGenerate 根据量产批次ID生成设备证书
//调用该接口根据批次ID（batchId）生成设备证书信息。在调用该接口前请确保已调用/cloud/device/name/upload接口，通过主动上传要量产设备的名称信息，获取量产的批次ID。
func (c *SyncApiClient) NameDeviceGenerate(apiVer, productKey, batchId string) (res response.NameDeviceGenerateResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.BatchId = batchId
	if body, errMsg = c.PostBody("/cloud/name/device/generate", req); errMsg != nil {
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

// DeviceApplyQuery 根据批次ID查询设备证书列表
func (c *SyncApiClient) DeviceApplyQuery(apiVer, applyId string, pageNo, pageSize int) (res response.DeviceApplyQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ApplyId = applyId
	req.Params.PageNo = pageNo
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/living/device/apply/query", req); errMsg != nil {
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

// ProductNvrDeviceGenerate 获取量产批次ID（NVR产品专用）
func (c *SyncApiClient) ProductNvrDeviceGenerate(apiVer, productKey string, amount int64) (res response.ProductNvrDeviceGenerateResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.Amount = amount
	if body, errMsg = c.PostBody("/living/cloud/product/nvr/device/generate", req); errMsg != nil {
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

// ProductNvrDeviceCertQuery 根据批次ID查询设备证书列表（NVR产品专用）
func (c *SyncApiClient) ProductNvrDeviceCertQuery(apiVer, productKey, batchId string, pageNo, pageSize int) (res response.ProductNvrDeviceCertQueryResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.ProductKey = productKey
	req.Params.BatchId = batchId
	req.Params.PageNo = pageNo
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/living/cloud/product/nvr/device/cert/query", req); errMsg != nil {
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

// ProductList 带分页查询项目信息列表
func (c *SyncApiClient) ProductList(apiVer string, pageNo, pageSize int) (res response.ProjectListResponse) {
	var (
		body   []byte
		errMsg *error2.ErrorMsg
		err    error
	)
	req := request.InitRequest()
	req.ApiRequest.ApiVer = apiVer
	req.Params.PageNo = pageNo
	req.Params.PageSize = pageSize
	if body, errMsg = c.PostBody("/living/cloud/project/list", req); errMsg != nil {
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
