package functions

import (
	"math/rand"
	"os"
	"time"
)

// 判断文件夹是否存在；不存在就创建
func IsDir(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 生成随机字符串
const (
	RAND_NUM	= 0	// 纯数字
	RAND_LOWER = 1 // 小写字母
	RAND_UPPER = 2 // 大写字母
	RAND_ALL 	= 3 // 数字、大小写字母
)
func RandCode(size, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}