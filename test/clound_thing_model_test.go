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

var productKeys = []string{"a1eniNgbtHM", "a1kVN8miPkX", "a1kM4sYClk8", "a1K06JjLXit", "a18EBnT9KIp"}
var statuses = []int{0, 1, 3, 8}

func TestThingPropertiesGet(t *testing.T) {
	apiClient, _ := aliyun_iotx_client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingPropertiesGet("1.0.2", "", "a1qbi9QsZPc", "HFbYWi7qdLlGSaYs8Z4B")
	fmt.Printf("%v", res)

}

func TestThingInfoGet(t *testing.T) {
	apiClient, _ := aliyun_iotx_client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingInfoGet("1.0.2", "", "a1STiUbGGhP", "TLDevNamdc29194e76a7")
	fmt.Printf("%+v", res)
}

func TestThingsGet(t *testing.T) {
	apiClient, _ := aliyun_iotx_client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	for _, productKey := range productKeys {
		var iotIds []string
		for _, status := range statuses {
			fmt.Println(status)
			res := apiClient.ThingsGet("1.0.2", productKey, status, 1, 20)
			if res.Code == 200 && res.Data.TotalNum > 0 {
				for _, item := range res.Data.Items {
					iotIds = append(iotIds, item.IotId)
				}
			}
			fmt.Printf("%+v", res)
			fmt.Println()
		}
		fmt.Println(productKey, ":", iotIds)
	}

}
