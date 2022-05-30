/**
    @author: dongjs
    @date: 2022/5/26
    @description: 获取物的连接状态 返回值
**/

package response

type ThingStatusGetResponse struct {
	Id      string             `json:"id,omitempty"`
	Code    int                `json:"code,omitempty"`
	Message string             `json:"message,omitempty"`
	Data    thingStatusGetData `json:"data,omitempty"`
}

type thingStatusGetData struct {
	Status int   `json:"status,omitempty"`
	Time   int64 `json:"time,omitempty"`
}
