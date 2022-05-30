/**
    @author: dongjs
    @date: 2022/5/26
    @description:获取物的事件timeline数据 返回值
**/

package response

type ThingEventTimelineGetResponse struct {
	Id      string                    `json:"id,omitempty"`
	Code    int                       `json:"code,omitempty"`
	Message string                    `json:"message,omitempty"`
	Data    thingEventTimelineGetData `json:"data,omitempty"`
}

type thingEventTimelineGetData struct {
	Items []thingEventTimelineGetDataItem `json:"items,omitempty"`
}

type thingEventTimelineGetDataItem struct {
	EventCode string `json:"eventCode,omitempty"`
	IotId     string `json:"iotId,omitempty"`
	EventName string `json:"eventName,omitempty"`
	EventType string `json:"eventType,omitempty"`
	BatchId   string `json:"batchId,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	EventBody struct {
		ErrorCode int `json:"ErrorCode,omitempty"`
	} `json:"eventBody,omitempty"`
}
