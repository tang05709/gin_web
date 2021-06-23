package index

import (
	"festival/app/model"
	"net/http"

	"festival/app/common/response"
	"festival/app/request"
	systemService "festival/app/service/system"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func AdminIndex(c *gin.Context) {
	response.BuildTpl(c, "backend/home").ToTpl()
}

func LoginPage(c *gin.Context) {
	if systemService.IsLogin(c) {
		c.Redirect(http.StatusFound, "/admin/index")
		return
	}
	response.BuildTpl(c, "backend/login").ToTpl()
}

//验证登陆
func CheckLogin(c *gin.Context) {
	var req *request.UserReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).ToJson()
		return
	}

	if req == nil {
		response.ErrorResp(c).SetMsg("用户名或密码错误").ToJson()
		return
	}

	//比对验证码
	verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)

	if !verifyResult {
		response.ErrorResp(c).SetMsg("验证码不正确").ToJson()
		return
	}

	//验证账号密码
	if _, err := systemService.Login(req.UserName, req.Password, c); err != nil {
		response.ErrorResp(c).SetMsg("账号或密码不正确").ToJson()
	} else {
		response.SucessResp(c).SetMsg("登陆成功").ToJson()
	}
}

// 图形验证码
func CaptchaImage(c *gin.Context) {
	//config struct for digits
	//数字验证码配置
	//var configD = base64Captcha.ConfigDigit{
	//	Height:     80,
	//	Width:      240,
	//	MaxSkew:    0.7,
	//	DotCount:   80,
	//	CaptchaLen: 5,
	//}
	//config struct for audio
	//声音验证码配置
	//var configA = base64Captcha.ConfigAudio{
	//	CaptchaLen: 6,
	//	Language:   "zh",
	//}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumberAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	c.JSON(http.StatusOK, model.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

//注销
func Logout(c *gin.Context) {
	if systemService.IsLogin(c) {
		systemService.LoginOut(c)
	}

	c.Redirect(http.StatusFound, "/admin/login")
	c.Abort()
}
