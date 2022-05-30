/**
    @author: dongjs
    @date: 2022/5/26
    @description:获取物的属性返回值
**/

package response

type ThingPropertiesGetResponse struct {
	Id      string                   `json:"id,omitempty"`
	Code    int                      `json:"code,omitempty"`
	Message string                   `json:"message,omitempty"`
	Data    []thingPropertiesGetData `json:"data,omitempty"`
}

type thingPropertiesGetData struct {
	Attribute   string `json:"attribute,omitempty"`
	BatchId     string `json:"batchId,omitempty"`
	GmtModified int64  `json:"gmtModified,omitempty"`
	IotId       string `json:"iotId,omitempty"`
}
