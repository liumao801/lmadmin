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