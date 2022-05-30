/**
    @author: dongjs
    @date: 2022/5/24
    @description:
**/

package client

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dongjssd/aliyun_iotx_client/constant"
	error2 "github.com/dongjssd/aliyun_iotx_client/error"
	"github.com/dongjssd/aliyun_iotx_client/request"
	"github.com/dongjssd/aliyun_iotx_client/secret"
	"github.com/dongjssd/aliyun_iotx_client/util"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SyncApiClient struct {
	AppKey    string //云端唯一身份（AppKey）
	AppSecret string //云端密码（AppSecret）
	Host      string //API网关协议与地址
}

// InitSyncApiClient
/**
 * @Author dongjs
 * @Description //初始化api client
 * @Date 10:11 2022/5/24
 * @Param
 * @return
 **/
func InitSyncApiClient(appKey, appSecret, host string) (client *SyncApiClient, err error) {
	if appKey == "" {
		return nil, errors.New(error2.AliYunAppKeyNilErrorMsg)
	}

	if appSecret == "" {
		return nil, errors.New(error2.AliYunAppSecretNilErrorMsg)
	}

	if host == "" {
		host = "api.link.aliyun.com"
	}
	client = &SyncApiClient{AppSecret: appSecret, AppKey: appKey, Host: host}
	return
}

// PostBody
/**
 * @Author dongjs
 * @Description //发送请求
 * @Date 10:39 2022/5/24
 * @Param
 * @return
 **/
func (c *SyncApiClient) PostBody(path string, req request.Request) (body []byte, errMsg *error2.ErrorMsg) {
	var (
		reqBody []byte
		err     error
		resp    *http.Response
	)
	if reqBody, err = util.JsonReMarshal(req); err != nil {
		errMsg.Code = error2.AliYunApiRequestErrorCode
		errMsg.Message = err.Error()
		return
	}
	fmt.Println("reqBody:", string(reqBody))
	var apiHttpReq = request.InitIotApiRequest(reqBody, path)
	apiHttpReq.Host = c.Host
	httpReq := c.buildRequest(apiHttpReq)
	client := &http.Client{}
	if resp, err = client.Do(httpReq); err != nil {
		errMsg.Code = error2.AliYunApiDoErrorCode
		errMsg.Message = err.Error()
		return
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		errMsg.Code = error2.AliYunApiReadErrorCode
		errMsg.Message = err.Error()
		return
	}
	fmt.Println("resBody:", string(body))
	return
}

func (c *SyncApiClient) buildRequest(apiReq *request.IotApiRequest) *http.Request {
	c.makeIotApiRequest(apiReq)
	req, _ := http.NewRequest(apiReq.Method, apiReq.Url, bytes.NewBuffer(apiReq.Body))
	req.Header.Add("Content-Type", apiReq.GetFirstHeaderValue(constant.CloudApiHttpHeaderContentType))
	for key, values := range apiReq.Headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return req
}

func (c *SyncApiClient) makeIotApiRequest(apiReq *request.IotApiRequest) {
	/**
	 *  拼接URL
	 *  HTTP + HOST + PATH(With pathparameter) + Query Parameter
	 */
	var url = "https://" + c.Host + apiReq.Path
	apiReq.Url = url
	// 设置请求头中的时间戳
	currentTime := time.Now()
	location, _ := time.LoadLocation("GMT")
	date := currentTime.In(location).Format(time.RFC1123)
	apiReq.AddHeader(constant.CloudApiHttpHeaderDate, date)
	//设置请求头中的时间戳，以timeIntervalSince1970的形式
	apiReq.AddHeader(constant.CloudApiXCaTimestamp, strconv.FormatInt(currentTime.UnixMilli(), 10))
	//请求放重放Nonce,15分钟内保持唯一,建议使用UUID
	apiReq.AddHeader(constant.CloudApiXCaNonce, uuid.NewV4().String())
	//设置请求头中的UserAgent
	apiReq.AddHeader(constant.CloudApiHttpHeaderUserAgent, constant.CloudApiUserAgent)
	//设置请求头中的主机地址
	apiReq.AddHeader(constant.CloudApiHttpHeaderHost, apiReq.Host)
	//设置请求头中的Api绑定的的AppKey
	apiReq.AddHeader(constant.CloudApiXCaKey, c.AppKey)
	//设置签名版本号
	apiReq.AddHeader(constant.CloudApiXCaVersion, constant.CloudApiCaVersionValue)
	//设置请求数据类型
	apiReq.AddHeader(constant.CloudApiHttpHeaderContentType, constant.CloudApiContentTypeStream)
	//设置应答数据类型
	apiReq.AddHeader(constant.CloudApiHttpHeaderAccept, constant.CloudApiHttpHeaderAcceptValue)
	if apiReq.SignatureMethod != "" {
		apiReq.AddHeader(constant.CloudApiXCaSignatureMethod, apiReq.SignatureMethod)
	}
	/**
	 *  如果formParams不为空
	 *  将Form中的内容拼接成字符串后使用UTF8编码序列化成Byte数组后加入到Request中去
	 */
	if len(apiReq.Body) > 0 {
		apiReq.AddHeader(constant.CloudApiHttpHeaderContentMd5, secret.Base64AndMD5(apiReq.Body))
	}

	/**
	 *  将Request中的httpMethod、headers、path、queryParam、formParam合成一个字符串用hmacSha256算法双向加密进行签名
	 *  签名内容放到Http头中，用作服务器校验
	 */
	signature := c.sign(apiReq)
	apiReq.AddHeader(constant.CloudApiXCaSignature, signature)

	/**
	 *  凑齐所有HTTP头之后，将头中的数据全部放入Request对象中
	 *  Http头编码方式：先将字符串进行UTF-8编码，然后使用Iso-8859-1解码生成字符串
	 */
	for key, value := range apiReq.Headers {
		apiReq.Headers[key] = value
	}
}

//将Request中的httpMethod、headers、path、queryParam、formParam合成一个字符串用hmacSha256算法双向加密进行签名
func (c *SyncApiClient) sign(req *request.IotApiRequest) string {
	signStr := buildStringToSign(req)
	return secret.HmacSha256(signStr, c.AppSecret)
}

/**
 * 将Request中的httpMethod、headers、path、queryParam、formParam合成一个字符串
 */
func buildStringToSign(req *request.IotApiRequest) string {
	var strToSign []string
	strToSign = append(strToSign, req.Method)
	strToSign = append(strToSign, constant.CloudApiLf)

	//如果有@"Accept"头，这个头需要参与签名
	if len(req.Headers[constant.CloudApiHttpHeaderAccept]) > 0 {
		strToSign = append(strToSign, req.GetFirstHeaderValue(constant.CloudApiHttpHeaderAccept))
	}
	strToSign = append(strToSign, constant.CloudApiLf)

	//如果有@"Content-MD5"头，这个头需要参与签名
	if len(req.Headers[constant.CloudApiHttpHeaderContentMd5]) > 0 {
		strToSign = append(strToSign, req.GetFirstHeaderValue(constant.CloudApiHttpHeaderContentMd5))
	}
	strToSign = append(strToSign, constant.CloudApiLf)
	//如果有@"Content-Type"头，这个头需要参与签名
	if len(req.Headers[constant.CloudApiHttpHeaderContentType]) > 0 {
		strToSign = append(strToSign, req.GetFirstHeaderValue(constant.CloudApiHttpHeaderContentType))
	}
	strToSign = append(strToSign, constant.CloudApiLf)

	//签名优先读取HTTP_CA_HEADER_DATE，因为通过浏览器过来的请求不允许自定义Date（会被浏览器认为是篡改攻击）
	if len(req.Headers[constant.CloudApiHttpHeaderDate]) > 0 {
		strToSign = append(strToSign, req.GetFirstHeaderValue(constant.CloudApiHttpHeaderDate))
	}
	strToSign = append(strToSign, constant.CloudApiLf)

	//将headers合成一个字符串
	strToSign = append(strToSign, buildHeaders(req))
	strToSign = append(strToSign, constant.CloudApiLf)

	//将path、queryParam、formParam合成一个字符串
	strToSign = append(strToSign, buildResource(req))
	return strings.Join(strToSign, "")
}

/**
 *  将headers合成一个字符串
 *  需要注意的是，HTTP头需要按照字母排序加入签名字符串
 *  同时所有加入签名的头的列表，需要用逗号分隔形成一个字符串，加入一个新HTTP头@"X-Ca-Signature-Headers"
 */
func buildHeaders(req *request.IotApiRequest) string {
	var keys []string
	for key := range req.Headers {
		if len(key) > len(constant.CloudApiCaHeaderToSignPrefixSystem) && key[0:len(constant.CloudApiCaHeaderToSignPrefixSystem)] == constant.CloudApiCaHeaderToSignPrefixSystem {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	//同时所有加入签名的头的列表，需要用逗号分隔形成一个字符串，加入一个新HTTP头@"X-Ca-Signature-Headers"
	req.AddHeader(constant.CloudApiXCaSignatureHeaders, strings.Join(keys, ","))
	var keyValues []string
	for _, key := range keys {
		keyValues = append(keyValues, key+":"+req.GetFirstHeaderValue(key))
	}
	return strings.Join(keyValues, constant.CloudApiLf)
}

/**
 * 将path、queryParam、formParam合成一个字符串
 */
func buildResource(req *request.IotApiRequest) string {
	return req.Path
}
