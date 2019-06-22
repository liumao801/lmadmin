package controllers

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"net/http"
)

type CommonController struct {

}

func init()  {
	// 创建验证码信息
	CreateCaptcha()
}

// 验证码
var cpt *captcha.Captcha
// 创建验证码
func CreateCaptcha() {
	// 使用 beego 的缓存系统存储 验证码数据
	//store := cache.NewFileCache()
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 3
	cpt.StdHeight = 34
	cpt.StdWidth = 100
}
// 检测验证码是否正确
func CheckCaptcha(request *http.Request) bool {
	verified := cpt.VerifyReq(request)
	return  verified
}
/*
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
}*/