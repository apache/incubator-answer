package captcha

import (
	"fmt"
	"image/color"
	"sync"

	"github.com/mojocn/base64Captcha"
)

var store base64Captcha.Store
var once sync.Once

func NewCaptcha() {
	once.Do(func() {
		//var err error
		//RedisDb, err = arch.App.Cache.GetQuestion("cache")
		//if err != nil {
		//	store = base64Captcha.DefaultMemStore
		//	return
		//}
		//var ctx = context.Background()
		//_, err = RedisDb.Ping(ctx).Result()
		//
		//if err != nil {
		//	store = base64Captcha.DefaultMemStore
		//	return
		//}
		store = RedisStore{}
	})
}

// CaptchaClient
type CaptchaClient struct {
}

// NewCaptchaClient
func NewCaptchaClient() *CaptchaClient {
	return &CaptchaClient{}
}

func MakeCaptcha() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	//Configure the parameters of the CAPTCHA
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	//ConvertFonts Load fonts by name
	driver = driverString.ConvertFonts()
	//Create Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	//Generate
	id, b64s, err = captcha.Generate()
	return id, b64s, err

}

// VerifyCaptcha Verification code
func VerifyCaptcha(id string, VerifyValue string) bool {
	fmt.Println(id, VerifyValue)
	if store.Verify(id, VerifyValue, true) {
		//verify successfully
		return true
	} else {
		//Verification failed
		return false
	}
}
