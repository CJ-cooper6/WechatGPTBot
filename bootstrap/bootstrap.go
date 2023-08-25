package bootstrap

import (
	"github.com/eatmoreapple/openwechat"
	"log"
	"wechatbot/handlers"
)

func Run() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	// 免扫码登录
	err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}
	bot.LoginCallBack = func(body openwechat.CheckLoginResponse) {
		log.Printf("login success!")
	}
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
