生活物联网平台云端api

### 介绍
阿里云IoT网关，旨在通过http/https协议将物联网平台及衍生服务开放给开发者。本文档主要介绍如何访问阿里云IoT网关。IoT网关依赖阿里云API网关，IoT网关在阿里云API网关的基础上扩展了IoT的特殊业务协议。

### 安装

go get github.com/dongjssd/aliyun_iotx_client@v0.1.3

### 方法命名规则

/cloud/token/refresh => TokenRefresh()

### 示例

```
var appKey = "" //阿里云appKey
var appSecret = "" //阿里云appSecret
var host = "" //API网关协议与地址 默认api.link.aliyun.com
client, _ := aliyun_iotx_client.InitSyncApiClient(appKey, appSecret, host)
var apiVer = "1.0.1" //版本号
res := client.TokenGet(apiVer, "")

```



