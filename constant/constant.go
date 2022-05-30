/**
    @author: dongjs
    @date: 2022/5/25
    @description:
**/

package constant

// 定义常量

const (
	CloudApiXCaSignature               = "x-ca-signature"                          //签名Header
	CloudApiXCaSignatureHeaders        = "x-ca-signature-headers"                  //所有参与签名的Header
	CloudApiXCaTimestamp               = "x-ca-timestamp"                          //请求时间戳
	CloudApiXCaNonce                   = "x-ca-nonce"                              //请求放重放Nonce,15分钟内保持唯一,建议使用UUID
	CloudApiXCaKey                     = "x-ca-key"                                //APP KEY
	CloudApiXCaVersion                 = "CA_VERSION"                              //签名版本号
	CloudApiXCaSignatureMethod         = "X-Ca-Signature-Method"                   //签名算法
	CloudApiCaHeaderToSignPrefixSystem = "x-ca-"                                   //参与签名的系统Header前缀,只有指定前缀的Header才会参与到签名中
	CloudApiCaVersionValue             = "1"                                       //签名版本号
	CloudApiHttpHeaderHost             = "host"                                    //请求Header Host
	CloudApiHttpHeaderAccept           = "accept"                                  //请求Header Accept
	CloudApiHttpHeaderAcceptValue      = "application/json; charset=utf-8"         // 请求Header Accept value
	CloudApiHttpHeaderDate             = "date"                                    //请求Header Date
	CloudApiUserAgent                  = "ALIYUN-ANDROID-DEMO"                     //UserAgent
	CloudApiLf                         = "\n"                                      //换行符
	CloudApiHttpHeaderContentMd5       = "content-md5"                             //请求Body内容MD5 Header
	CloudApiHttpHeaderContentType      = "content-type"                            //请求Header Content-Type
	CloudApiHttpHeaderUserAgent        = "user-agent"                              //
	CloudApiContentTypeStream          = "application/octet-stream; charset=utf-8" //
	CloudApiXCaSignatureMethodValue    = "HmacSHA256"                              //
)
