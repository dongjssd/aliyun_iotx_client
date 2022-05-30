/**
    @author: dongjs
    @date: 2022/5/26
    @description:获取物的属性timeline数据 返回值
**/

package response

type ThingPropertyTimelineGetResponse struct {
	Id      string                       `json:"id,omitempty"`
	Code    int                          `json:"code,omitempty"`
	Message string                       `json:"message,omitempty"`
	Data    thingPropertyTimelineGetData `json:"data,omitempty"`
}

type thingPropertyTimelineGetData struct {
	Items thingPropertyTimelineGetDataItem `json:"items"`
}

type thingPropertyTimelineGetDataItem struct {
	IotId      string `json:"iotId,omitempty"`
	ModifyTime string `json:"modifytime,omitempty"`
	Property   string `json:"property,omitempty"`
	BatchId    string `json:"batchId,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty"`
	Group      string
}
