package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 将字符串加密成 md5
func Str2md5(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)  //将[]byte转成16进制
}

// RandomStr 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomStr(len int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, len)
	for i := 0; i < len; i++ {
		t := rand.Intn(62)
		switch {
		case t < 10:
			result = append(result, strconv.Itoa(rand.Intn(10)))
			break
		case t < 36:
			result = append(result, string(rand.Intn(26)+65))
			break
		default:
			result = append(result, string(rand.Intn(26)+97))
			break
		}
	}
	return strings.Join(result, "")
}