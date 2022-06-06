/**
    @author: dongjs
    @date: 2022/5/31
    @description:
**/

package test

import (
	"fmt"
	"github.com/dongjssd/aliyun_iotx_client"
	"testing"
)

func TestDeviceApplyQuery(t *testing.T) {
	apiClient, _ := aliyun_iotx_client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.DeviceApplyQuery("1.0.1", "4627293", 1, 10)
	fmt.Printf("%+v", res)
}

func TestProductGet(t *testing.T) {
	apiClient, _ := aliyun_iotx_client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.CloudThingProduct("1.1.2", "a1qbi9QsZPc")
	fmt.Printf("%v", res)
}
