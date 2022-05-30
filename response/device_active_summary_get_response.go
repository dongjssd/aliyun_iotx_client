/**
    @author: dongjs
    @date: 2022/5/27
    @description:获取活跃设备的概览信息 返回值
**/

package response

type DeviceActiveSummaryGetResponse struct {
	Id           string `json:"id,omitempty"`
	Code         int    `json:"code,omitempty"`
	Message      string `json:"message,omitempty"`
	LocalizedMsg string `json:"localizedMsg,omitempty"`
	Data         struct {
		Data   []deviceActiveData `json:"data,omitempty"`
		Params []struct {
			Page []struct {
				Total int `json:"total,omitempty"`
				Size  int `json:"size,omitempty"`
				To    int `json:"to,omitempty"`
			} `json:"page,omitempty"`
		} `json:"params,omitempty"`
	} `json:"data,omitempty"`
}

type deviceActiveData struct {
	Num14Day     int    `json:"num_14day,omitempty"`
	Num7Day      int    `json:"num_7day,omitempty"`
	LastNum14Day int    `json:"last_num_14day,omitempty"`
	Rate14Day    string `json:"rate_14day,omitempty"`
	Rate30Day    string `json:"rate_30day,omitempty"`
	ProductKey   string `json:"product_key,omitempty"`
	DateTime     string `json:"date_time,omitempty"`
	Num          int    `json:"num,omitempty"`
	Rate7Day     string `json:"rate_7day,omitempty"`
	LastNum7Day  int    `json:"last_num_7day,omitempty"`
	Num30Day     int    `json:"num_30day,omitempty"`
	LastNum30Day int    `json:"last_num_30day,omitempty"`
}
