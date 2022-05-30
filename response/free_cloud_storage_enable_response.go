/**
    @author: dongjs
    @date: 2022/5/27
    @description:设置免费云存储套餐立即生效 返回值
**/

package response

type FreeCloudStorageEnableResponse struct {
	Id           string                     `json:"id,omitempty"`
	Code         int                        `json:"code,omitempty"`
	Message      string                     `json:"message,omitempty"`
	LocalizedMsg string                     `json:"localizedMsg,omitempty"`
	Data         freeCloudStorageEnableData `json:"data,omitempty"`
}

type freeCloudStorageEnableData struct {
	StorageType  int    `json:"type,omitempty"`
	LifeCycle    int    `json:"lifeCycle,omitempty"`
	Months       int    `json:"months,omitempty"`
	Consumed     int    `json:"consumed,omitempty"`
	StartTime    string `json:"startTime,omitempty"`
	EndTime      string `json:"endTime,omitempty"`
	Expired      int    `json:"expired,omitempty"`
	StartTimeUTC string `json:"startTimeUTC,omitempty"`
	EndTimeUTC   string `json:"endTimeUTC,omitempty"`
}
