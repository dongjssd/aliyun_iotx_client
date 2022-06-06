/**
    @author: dongjs
    @date: 2022/5/24
    @description:
**/

package request

import (
	"fmt"
	"time"
)

// Request 阿里云请求数据
type Request struct {
	Id         string     `json:"id,omitempty"`      //请求ID，如使用生活物联网平台提供的SDK，则会自动生成；如果您自行调用API，则需要自己生成。该参数主要用于关联请求应答及问题定位。
	Version    string     `json:"version,omitempty"` //开放平台的版本号，当前固定为1.0。
	ApiRequest ApiRequest `json:"request,omitempty"` //系统请求参数。
	Params     ApiParam   `json:"params"`            //业务请求参数，具体内容因不同的接口而不同，部分接口可能不需要入参，仅通过默认的用户鉴权信息即可操作，具体请参见相应的接口文档。
}

// ApiRequest 系统请求参数
type ApiRequest struct {
	ApiVer string `json:"apiVer,omitempty"` //所调用接口对应的版本号，具体参见相应的接口文档。
	//IotToken   string `json:"iotToken,omitempty"`   //该参数仅调用客户端API时生效，主要用于确认请求发起者的系统参数，使用生活物联网平台提供的账号SDK时该值会自动生成。
	//CloudToken string `json:"cloudToken,omitempty"` //该参数仅调用云端API时生效，需要您使用项目ID自行获取对应的Token，具体API请参见云端资源服务的获取云端资源token。
	//Language   string `json:"language,omitempty"`   //用于传递多语言信息，该参数仅调用客户端API时生效。
}

// ApiParam 业务请求参数，具体内容因不同的接口而不同，部分接口可能不需要入参，仅通过默认的用户鉴权信息即可操作，具体请参见相应的接口文档。
type ApiParam struct {
	GrantType  string `json:"grantType,omitempty"`
	Res        string `json:"res,omitempty"`
	CloudToken string `json:"cloudToken,omitempty"` //该参数仅调用云端API时生效，需要您使用项目ID自行获取对应的Token，具体API请参见云端资源服务的获取云端资源token。
	ApiParamProductList
	ApiParamProduct
	ApiParamDeviceGenerate
	ApiParamDeviceNameUpload
	ApiParamNameDeviceGenerate
	ApiParamDeviceApplyQuery
	ApiParamThingPropertiesGet
	ApiParamThingServiceInvoke
	ApiParamThingPropertiesSet
	ApiParamThingsGet
	ApiParamThingEventTimelineGet
	ApiParamDeviceCustomizemessageNotify
	ApiParamAccountGetByOpenId
	ApiParamDeviceQueryByUser
	ApiParamAccountInfoUpdate
	ApiParamUserRegionGet
	ApiParamAccountByPage
	ApiParamAccountRegionGet
	ApiParamDeviceRegionDistribute
	ApiParamUserBindingDeviceBatchBind
	ApiParamFreeCloudStorageConsume
	ApiParamCloudStorageCommodityCheck
	ApiParamCloudStorageCommodityBuy
	ApiParamCloudStorageOrderGet
	ApiParamCloudStorageTransfer
	ApiParamStreamingOrderCreate
	ApiParamDeviceActiveSummaryGet
}

// InitRequest
/**
 * @Author dongjs
 * @Description //初始化请求
 * @Date 10:56 2022/5/24
 * @Param
 * @return
 **/
func InitRequest() Request {
	return Request{Id: fmt.Sprint(time.Now().UnixMicro()), Version: "1.0"}
}

// ApiParamProductList 产品列表请求参数
type ApiParamProductList struct {
	PageNo   int `json:"pageNo,omitempty"`   //从1开始的分页页码
	PageSize int `json:"pageSize,omitempty"` //分页大小。
}

// ApiParamProduct 查询单个产品
type ApiParamProduct struct {
	ProductKey string `json:"productKey,omitempty"` //物的产品Key，设备证书信息之一。创建产品时，生活物联网平台为该产品颁发的全局唯一标识。
}

// ApiParamDeviceGenerate 量产生成设备证书请求参数
type ApiParamDeviceGenerate struct {
	//ProductKey string `json:"productKey"` //量产产品的Key，设备证书信息之一。创建产品时，生活物联网平台为该产品颁发的全局唯一标识。 重复存在
	Amount int64 `json:"amount,omitempty"` //量产的设备数量。
}

// ApiParamDeviceNameUpload 创建量产批次接口请求参数
type ApiParamDeviceNameUpload struct {
	DeviceNames []string `json:"deviceNames,omitempty"`
}

// ApiParamNameDeviceGenerate 根据量产批次ID生成设备证书请求参数
type ApiParamNameDeviceGenerate struct {
	BatchId string `json:"batchId,omitempty"`
}

// ApiParamDeviceApplyQuery 根据批次ID查询设备证书列表请求参数
type ApiParamDeviceApplyQuery struct {
	ApplyId string `json:"applyId,omitempty"` //申请的批次ID
	//pageNo   int    `json:"pageNo,omitempty"` //从1开始的查询页码
	//pageSize int    `json:"pageSize,omitempty"` //分页大小，单页的记录上限为 200
}

// ApiParamThingPropertiesGet 获取物的属性 请求参数
type ApiParamThingPropertiesGet struct {
	IotId      string `json:"iotId,omitempty"`      //设备ID，生活物联网平台为设备颁发的ID，设备的唯一标识符。
	DeviceName string `json:"deviceName,omitempty"` //设备的名称，设备证书信息之一。在注册设备时，自定义的或系统生成的设备名称，具备产品维度内的唯一性
}

// ApiParamThingServiceInvoke 触发物的服务 请求参数
type ApiParamThingServiceInvoke struct {
	Identifier string      `json:"identifier,omitempty"` //服务的标识符
	Args       interface{} `json:"args,omitempty"`       //服务入参
}

// ApiParamThingPropertiesSet 设置物的属性 请求参数
type ApiParamThingPropertiesSet struct {
	Items interface{} `json:"items,omitempty"` //设置功能属性的参数列表，列表是由功能属性标识符（key）和要设置的值（value）组成的标准Map
}

// ApiParamThingsGet 批量获取物的列表 请求列表
type ApiParamThingsGet struct {
	Status      int `json:"status"`                //设备状态。0（表示未激活）；1（表示在线）；3（表示离线）；8（表示禁用）
	CurrentPage int `json:"currentPage,omitempty"` //分页查询
}

// ApiParamThingEventTimelineGet 获取物的事件timeline数据 请求参数
type ApiParamThingEventTimelineGet struct {
	EventType string `json:"eventType,omitempty"` //事件的类型
	Start     int64  `json:"start,omitempty"`     //timeline起始时间
	End       int64  `json:"end,omitempty"`       //timeline终止时间
	Ordered   bool   `json:"ordered,omitempty"`   //查询方式。true（表示顺序）；false（表示逆序）
}

// ApiParamDeviceCustomizemessageNotify 向指定设备发送自定义消息 请求参数
type ApiParamDeviceCustomizemessageNotify struct {
	MessageContent interface{} `json:"messageContent,omitempty"`
}

// ApiParamAccountGetByOpenId 通过三方外标查询账号信息 请求参数
type ApiParamAccountGetByOpenId struct {
	OpenId       string `json:"openId,omitempty"`       //三方账号的外标或ID
	OpenIdAppKey string `json:"openIdAppKey,omitempty"` //开放平台的AppKey
}

// ApiParamDeviceQueryByUser 获取用户绑定的设备列表详情 请求参数
type ApiParamDeviceQueryByUser struct {
	IdentityId string `json:"identityId,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
}

// ApiParamAccountInfoUpdate 更新用户账号信息 请求参数
type ApiParamAccountInfoUpdate struct {
	NickName string `json:"nickName,omitempty"`
}

// ApiParamUserRegionGet 获取用户注册的区域 请求参数
type ApiParamUserRegionGet struct {
	Phone             string `json:"phone,omitempty"`
	PhoneLocationCode string `json:"phoneLocationCode,omitempty"`
	Email             string `json:"email,omitempty"`
}

// ApiParamAccountByPage 倒序分页查询用户列表 请求参数
type ApiParamAccountByPage struct {
	Count int `json:"count,omitempty"`
}

// ApiParamAccountRegionGet 获取用户账号对应的 Region 信息 请求参数
type ApiParamAccountRegionGet struct {
	QueryType   string `json:"type,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	AuthCode    string `json:"authCode,omitempty"`
}

// ApiParamDeviceRegionDistribute 批量把设备迁移到当前region 请求参数
type ApiParamDeviceRegionDistribute struct {
	BindingCheck           bool `json:"bindingCheck,omitempty"`
	PushDeviceRegionSwitch bool `json:"pushDeviceRegionSwitch,omitempty"`
}

// ApiParamUserBindingDeviceBatchBind  批量绑定设备 请求参数
type ApiParamUserBindingDeviceBatchBind struct {
	DeviceList []map[string]string `json:"deviceList,omitempty"`
	ProjectId  string              `json:"projectId,omitempty"`
}

// ApiParamFreeCloudStorageConsume  领取免费的云存储套餐 请求参数
type ApiParamFreeCloudStorageConsume struct {
	EnableDefaultPlan  bool `json:"enableDefaultPlan,omitempty"`
	ImmediateUse       bool `json:"immediateUse,omitempty"`
	PreRecordDuration  int  `json:"preRecordDuration,omitempty"`
	RecordDuration     int  `json:"recordDuration,omitempty"`
	Quota              int  `json:"quota,omitempty"`
	EventRecordProlong bool `json:"eventRecordProlong,omitempty"`
}

// ApiParamCloudStorageCommodityCheck 查询云存储套餐是否可以购买 请求参数
type ApiParamCloudStorageCommodityCheck struct {
	CommodityCode string `json:"commodityCode,omitempty"`
	Specification string `json:"specification,omitempty"`
}

// ApiParamCloudStorageCommodityBuy 购买云存储套餐 请求参数
type ApiParamCloudStorageCommodityBuy struct {
	UserName              string `json:"userName,omitempty"`
	Copies                int    `json:"copies,omitempty"`
	MaxRecordFileDuration int    `json:"maxRecordFileDuration,omitempty"`
}

// ApiParamCloudStorageOrderGet 获取云存储套餐的订单详情 请求参数
type ApiParamCloudStorageOrderGet struct {
	OrderId string `json:"orderId,omitempty"`
}

// ApiParamCloudStorageTransfer 用户云存套餐转移 请求参数
type ApiParamCloudStorageTransfer struct {
	SrcIotId   string `json:"srcIotId,omitempty"`
	SrcOrderId string `json:"srcOrderId,omitempty"`
	DstIotId   string `json:"dstIotId,omitempty"`
}

// ApiParamStreamingOrderCreate 创建直播分享套餐订单 请求参数
type ApiParamStreamingOrderCreate struct {
	UserId string `json:"userId,omitempty"`
}

// ApiParamDeviceActiveSummaryGet 获取活跃设备的概览信息 请求参数
type ApiParamDeviceActiveSummaryGet struct {
	ProductKeys []string `json:"productKeys,omitempty"`
}
