/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

type ProjectListResponse struct {
	Id      string          `json:"id,omitempty"`
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    projectListData `json:"data,omitempty"`
}

type projectListData struct {
	Total    int64                 `json:"total"`
	PageNo   int                   `json:"pageNo"`
	PageSize int                   `json:"pageSize"`
	Data     []projectListInfoData `json:"data"`
}

type projectListInfoData struct {
	TenantId     string `json:"tenantId"`     //项目归属租户id
	ProjectId    string `json:"projectId"`    //项目id
	IsolationId  string `json:"isolationId"`  //隔离id
	Name         string `json:"name"`         //项目名称
	Description  string `json:"description"`  //项目描述
	ProductCount int64  `json:"productCount"` //关联产品数量，就是该项目下已经创建的产品数量
	AppCount     int64  `json:"appCount"`     //关联应用数量
	AccountCount int64  `json:"accountCount"` //项目成员数量
	DevelopCount int64  `json:"developCount"` //开发中的产品数量
	PublishCount int64  `json:"publishCount"` //已发布的产品数量
	GmtCreate    int64  `json:"gmtCreate"`    //是否为共享项目
	Share        bool   `json:"share"`        //是否为共享项目
}
