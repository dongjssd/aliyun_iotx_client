/**
    @author: dongjs
    @date: 2022/5/26
    @description: 领取免费的云存储套餐 返回值
**/

package response

type FreeCloudStorageConsumeResponse struct {
	Id           string               `json:"id,omitempty"`
	Code         int                  `json:"code,omitempty"`
	Message      string               `json:"message,omitempty"`
	LocalizedMsg string               `json:"localizedMsg,omitempty"`
	Data         freeCloudStorageData `json:"data,omitempty"`
}

type freeCloudStorageData struct {
	StorageType  int    `json:"type,omitempty"`
	LifeCycle    int    `json:"lifecycle,omitempty"`
	Months       int    `json:"months,omitempty"`
	Consumed     int    `json:"consumed,omitempty"`
	StartTime    string `json:"startTime,omitempty"`
	EndTime      string `json:"endTime,omitempty"`
	Expired      int    `json:"expired,omitempty"`
	StartTimeUTC string `json:"startTimeUTC,omitempty"`
	EndTimeUTC   string `json:"endTimeUTC,omitempty"`
	RemainQuota  int    `json:"remainQuota,omitempty"`
}
