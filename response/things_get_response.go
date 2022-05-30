/**
    @author: dongjs
    @date: 2022/5/26
    @description:批量获取物的列表 返回值
**/

package response

type ThingsGetResponse struct {
	Id      string        `json:"id,omitempty"`
	Code    int           `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Data    thingsGetData `json:"data"`
}

type thingsGetData struct {
	TotalNum int64              `json:"totalNum"`
	Items    []thingInfoGetData `json:"items,omitempty"`
}
