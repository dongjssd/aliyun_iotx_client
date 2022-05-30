/**
    @author: dongjs
    @date: 2022/5/24
    @description:
**/

package secret

import (
	"crypto/md5"
	"encoding/base64"
)

// Base64AndMD5
/**
 * @Author dongjs
 * @Description //Content-MD5是指Body的MD5值，只有当Body非Form表单时才计算 MD5，计算方式如下。
 * @Date 16:15 2022/5/23
 * @Param
 * @return
 **/
func Base64AndMD5(req []byte) string {
	m := md5.New()
	m.Write(req)
	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}
