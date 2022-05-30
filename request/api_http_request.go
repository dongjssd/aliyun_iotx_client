/**
    @author: dongjs
    @date: 2022/5/24
    @description:云端api请求对象
**/

package request

import (
	"github.com/dongjssd/aliyun_iotx_client/constant"
	"strings"
)

type IotApiRequestMessage struct {
	Body    []byte
	BodyStr string
	Headers map[string][]string
}

type IotApiRequest struct {
	IotApiRequestMessage
	Method          string
	Host            string
	Path            string
	Url             string
	SignatureMethod string
}

func InitIotApiRequest(body []byte, path string) *IotApiRequest {
	var req = &IotApiRequest{}
	var headers = make(map[string][]string)
	req.Headers = headers
	req.Body = body
	req.Path = path
	req.Method = "POST"
	req.SignatureMethod = constant.CloudApiXCaSignatureMethodValue
	return req
}

func (c *IotApiRequest) AddHeader(key, value string) {
	key = strings.ToLower(key)
	headers := c.Headers[key]
	headers = append(headers, value)
	c.Headers[key] = headers
}

func (c *IotApiRequest) GetFirstHeaderValue(key string) string {
	if len(c.Headers[key]) > 0 {
		return c.Headers[key][0]
	} else {
		return ""
	}

}
