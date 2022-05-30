/**
    @author: dongjs
    @date: 2022/5/24
    @description:
**/

package test

import (
	"fmt"
	"github.com/dongjssd/aliyun_iotx_client/client"
	"testing"
	"time"
)

func TestTokenGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.TokenGet("1.0.1", "a123pTvF86e5yT7K")
	fmt.Printf("%+v", res)
}

func TestTokenRefresh(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.TokenRefresh("1.0.1", "dfa5f1a460a1428aaa561615505ab354")
	fmt.Printf("%+v", res)
}

func TestCloudThingProductList(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.CloudThingProductList("1.1.2", 1, 10)
	fmt.Printf("%+v", res)
}

//a1eniNgbtHM
func TestCloudThingProduct(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.CloudThingProduct("1.1.2", "a1eniNgbtHM")
	fmt.Printf("%+v", res)
}

func TestAmountDeviceGenerate(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.AmountDeviceGenerate("1.1.2", "a1eniNgbtHM", 1)
	fmt.Printf("%+v", res)
}

func TestDeviceNameUpload(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.DeviceNameUpload("1.1.2", "a1eniNgbtHM", []string{"deviceName1", "deviceName2"})
	fmt.Printf("%+v", res)
}

func TestNameDeviceGenerate(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.NameDeviceGenerate("1.1.2", "a1eniNgbtHM", "21312321")
	fmt.Printf("%+v", res)
}

func TestDeviceApplyQuery(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.DeviceApplyQuery("1.0.1", "21212", 1, 100)
	fmt.Printf("%+v", res)
}

func TestProductNvrDeviceGenerate(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ProductNvrDeviceGenerate("1.0.0", "a1eniNgbtHM", 1)
	fmt.Printf("%+v", res)
}

func TestProductNvrDeviceCertQuery(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ProductNvrDeviceCertQuery("1.0.0", "a1eniNgbtHM", "22", 1, 10)
	fmt.Printf("%+v", res)
}

func TestProductList(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ProductList("1.0.0", 1, 10)
	fmt.Printf("%+v", res)
}

func TestThingPropertiesGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingPropertiesGet("1.0.2", "", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y")
	fmt.Printf("%+v", res)
}

func TestThingTslGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingTslGet("1.0.2", "", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y")
	fmt.Printf("%+v", res)
}

func TestThingServiceInvoke(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingServiceInvoke("1.0.3", "", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", "DayNightMode", make(map[string]interface{}))
	fmt.Printf("%+v", res)
}

func TestThingPropertiesSet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	var items = make(map[string]interface{})
	items["LightSwitch1"] = 0
	res := apiClient.ThingPropertiesSet("1.0.2", "", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", items)
	fmt.Printf("%+v", res)
}

func TestThingStatusGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingStatusGet("1.0.2", "", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y")
	fmt.Printf("%+v", res)
}

func TestThingInfoGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingInfoGet("1.0.2", "", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y")
	fmt.Printf("%+v", res)
}

func TestThingsGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingsGet("1.0.2", "a1eniNgbtHM", 0, 1, 20)
	fmt.Printf("%+v", res)
}

func TestThingEventTimelineGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingEventTimelineGet("1.0.2", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", "",
		"22", "set", time.Now().AddDate(0, 0, -10).UnixMilli(), time.Now().UnixMilli(), 20, true)
	fmt.Printf("%+v", res)
}

func TestThingPropertyTimelineGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.ThingPropertyTimelineGet("1.0.2", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", "",
		"22", time.Now().AddDate(0, 0, -10).UnixMilli(), time.Now().UnixMilli(), 20, true)
	fmt.Printf("%+v", res)
}

func TestDeviceSubdeviceQuery(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.DeviceSubdeviceQuery("1.0.0", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", "", 1, 20)
	fmt.Printf("%+v", res)
}

func TestDeviceRegionGet(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.DeviceRegionGet("1.0.1", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y")
	fmt.Printf("%+v", res)
}

func TestDeviceCustomizedmessageNotify(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	var message = make(map[string]interface{})
	message["1111"] = "1231232"
	res := apiClient.DeviceCustomizedmessageNotify("1.0.0", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", "", message)
	fmt.Printf("%+v", res)
}

func TestAccountGetByOpenId(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.AccountGetByOpenId("1.0.5", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y")
	fmt.Printf("%+v", res)
}

func TestDeviceQueryByUser(t *testing.T) {
	apiClient, _ := client.InitSyncApiClient("29449075", "32891e4b647a8ea3e1cdd8f613792668", "")
	res := apiClient.DeviceQueryByUser("1.0.6", "a1eniNgbtHM", "v0wE8pvz0IagOSqEBe0y", "", 1, 10)
	fmt.Printf("%+v", res)
}
