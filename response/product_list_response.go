/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

// ProductListResponse 产品列表返回值
type ProductListResponse struct {
	Id      string        `json:"id,omitempty"`
	Code    int           `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
	Data    []productList `json:"data,omitempty"`
}

// ProductList 产品列表返回值
type productList struct {
	ProductKey    string `json:"productKey,omitempty"`    // 产品的Key，设备证书信息之一创建产品时，生活物联网平台为该产品颁发的全局唯一标识
	DataFormat    string `json:"dataFormat,omitempty"`    //产品的数据格式
	NetType       string `json:"netType,omitempty"`       //产品入网类型
	ProductSecret string `json:"productSecret,omitempty"` //产品的密钥，生活物联网平台颁发的产品密钥，通常与ProductKey成对出现
	NodeType      string `json:"nodeType,omitempty"`      //节点类型
	Name          string `json:"name,omitempty"`          //产品名称
	Region        string `json:"region,omitempty"`        //产品所属地域
	CategoryId    int64  `json:"categoryId,omitempty"`    //归属品类ID
	Status        int    `json:"status,omitempty"`        //产品状态0（表示开发中）；1（表示已发布）
	GmtCreate     int64  `json:"gmtCreate,omitempty"`     //创建产品的时间
	GmtModified   int64  `json:"gmtModified,omitempty"`   //修改产品的时间
	ProductId     int64  `json:"productId,omitempty"`     //产品ID，即productKey的十六进制字符串
	RbacTenantId  string `json:"rbacTenantId,omitempty"`  //租户ID
}
