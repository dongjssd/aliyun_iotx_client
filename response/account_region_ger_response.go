/**
    @author: dongjs
    @date: 2022/5/26
    @description:获取用户账号对应的 Region 信息 返回值
**/

package response

type AccountRegionGetResponse struct {
	Id           string               `json:"id,omitempty"`
	Code         int                  `json:"code,omitempty"`
	Message      string               `json:"message,omitempty"`
	LocalizedMsg string               `json:"localizedMsg,omitempty"`
	Data         accountRegionGetData `json:"data,omitempty"`
}

type accountRegionGetData struct {
	RegionId             string `json:"regionId,omitempty"`
	MqttEndpoint         string `json:"mqttEndpoint,omitempty"`
	ApiGatewayEndpoint   string `json:"apiGatewayEndpoint,omitempty"`
	OaApiGatewayEndpoint string `json:"oaApiGatewayEndpoint,omitempty"`
	RegionEnglishName    string `json:"regionEnglishName,omitempty"`
	PushChannelEndpoint  string `json:"pushChannelEndpoint,omitempty"`
	ShortRegionId        string `json:"shortRegionId,omitempty"`
}
