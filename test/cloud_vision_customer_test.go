/**
    @author: dongjs
    @date: 2022/5/27
    @description:
**/

package test

import (
	"fmt"
	"github.com/dongjssd/aliyun_iotx_client"
	"testing"
)

func TestFreeCloudStorageGet(t *testing.T) {
	apClient, _ := aliyun_iotx_client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apClient.FreeCloudStorageGet("1.0.1", "a123pTvF86e5yT7K")
	fmt.Printf("%+v", res)
}
