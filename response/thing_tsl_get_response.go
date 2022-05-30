/**
    @author: dongjs
    @date: 2022/5/26
    @description: 获取物的模板返回值
**/

package response

type ThingTslGetResponse struct {
	Id      string          `json:"id,omitempty"`
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    thingTslGetData `json:"data,omitempty"`
}

type thingTslGetData struct {
	Schema     string             `json:"schema,omitempty"`
	Profile    thingTslProfile    `json:"profile,omitempty"`
	Link       string             `json:"link,omitempty"`
	Services   []thingTslService  `json:"services,omitempty"`
	Properties []thingTslProperty `json:"properties,omitempty"`
	Events     []thingTslEvent    `json:"events,omitempty"`
}
type thingTslProfile struct {
	ProductKey string `json:"productKey,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
}

type thingTslService struct {
	OutputData []thingTslOutPutData `json:"outputData,omitempty"`
	Identifier string               `json:"identifier,omitempty"`
	InputData  []interface{}        `json:"inputData,omitempty"`
	Method     string               `json:"method,omitempty"`
	Name       string               `json:"name,omitempty"`
	Required   bool                 `json:"required,omitempty"`
	CallType   string               `json:"callType,omitempty"`
}

type thingTslProperty struct {
	Identifier string `json:"identifier,omitempty"`
	DataType   struct {
		DataTypeType string `json:"type,omitempty"`
		Specs        struct {
			Spec0 string `json:"0,omitempty"`
			Spec1 string `json:"1,omitempty"`
		} `json:"specs,omitempty"`
	} `json:"dataType,omitempty"`
	Name       string `json:"name,omitempty"`
	AccessMode string `json:"accessMode,omitempty"`
	Required   bool   `json:"required,omitempty"`
}

type thingTslEvent struct {
	OutputData []thingTslOutPutData `json:"outputData,omitempty"`
	Identifier string               `json:"identifier,omitempty"`
	Method     string               `json:"method,omitempty"`
	Name       string               `json:"name,omitempty"`
	EventType  string               `json:"type,omitempty"`
	Required   bool                 `json:"required,omitempty"`
}

type thingTslOutPutData struct {
	Identifier string `json:"identifier,omitempty"`
	DataType   struct {
		Specs struct {
			Spec0 string `json:"0,omitempty"`
		} `json:"specs,omitempty"`
		DataTypeType string `json:"type,omitempty"`
	} `json:"dataType,omitempty"`
}
