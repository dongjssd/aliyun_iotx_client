/**
    @author: dongjs
    @date: 2022/5/24
    @description:
**/

package util

import (
	"encoding/json"
	"github.com/dongjssd/aliyun_iotx_client/request"
)

// JsonReMarshal
/**
 * @Author dongjs
 * @Description //json排序
 * @Date 18:04 2022/5/23
 * @Param
 * @return
 **/
func JsonReMarshal(req request.Request) (output []byte, err error) {
	var bytes []byte
	if bytes, err = json.Marshal(req); err != nil {
		return
	}
	var ice interface{}
	if err = json.Unmarshal(bytes, &ice); err != nil {
		return
	}
	if output, err = json.Marshal(ice); err != nil {
		return
	}
	return
}
