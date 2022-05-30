/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package response

// ProductResponse 单个产品返回数据
type ProductResponse struct {
	Id      string  `json:"id,omitempty"`
	Code    int     `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Data    product `json:"data,omitempty"`
}

// Product 单个产品返回数据
type product struct {
	ProductKey          string `json:"productKey,omitempty"`    //产品的Key
	DataFormat          string `json:"dataFormat,omitempty"`    //数据格式
	NetType             string `json:"netType,omitempty"`       //产品入网类型
	ProductSecret       string `json:"productSecret,omitempty"` //产品的密钥，生活物联网平台颁发的产品密钥，通常与ProductKey成对出现
	NodeType            string `json:"nodeType,omitempty"`      //节点类型
	Name                string `json:"name,omitempty"`          //产品的名称
	Region              string `json:"region,omitempty"`        //产品所属地域
	CategoryId          int64  `json:"categoryId,omitempty"`    //归属品类ID
	Status              string `json:"status,omitempty"`        //产品状态DEVELOPMENT_STATUS（表示开发中）；RELEASE_STATUS（表示已发布）；NOT_SPECIFY（表示未指定状态）
	GmtCreate           int64  `json:"gmtCreate,omitempty"`     //创建产品的时间
	GtModified          int64  `json:"gmtModified,omitempty"`   //修改产品的时间
	ProductId           int64  `json:"productId,omitempty"`     //产品ID，即productKey的十六进制字符串
	RbacTenantId        string `json:"rbacTenantId,omitempty"`  //租户ID
	CategoryName        string `json:"categoryName,omitempty"`
	DeviceNumLimit      int64  `json:"deviceNumLimit,omitempty"`
	Image               string `json:"image,omitempty"`
	CategoryKey         string `json:"categoryKey,omitempty"`
	Namespace           string `json:"namespace,omitempty"`
	Domain              string `json:"domain,omitempty"`
	AliyunCommodityCode string `json:"aliyunCommodityCode,omitempty"`
	ProductModel        string `json:"productModel,omitempty"`
	AccessMethod        string `json:"accessMethod,omitempty"`
}
