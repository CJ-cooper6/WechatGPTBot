# WechatGPTBot

最近发现了封装 OpenAI 网页版 ChatGPT 接口的项目，可以免费调用ChatGPT接口[chatgpt-api](https://github.com/zhuweiyou/chatgpt-api)。

想到要是接入微信，家人或朋友可以更方便的体验ChatGPT，于是有了这个项目。

## 目前实现功能

+ 私聊回复
+ 群聊@回复


## 快速开始
**1、运行chatgpt-api**
````
docker run -d -p 3000:3000 zhuweiyou/chatgpt-api:20230424
````
**2、获取网页版 access_token**

访问 https://chat.openai.com/chat 成功登录之后, 打开浏览器开发者工具 (F12) -> 刷新页面- > Network 找到 /api/auth/session 请求, 复制 accessToken 存到你本地配置

**3、修改配置文件**
````
# 获取项目
git clone https://github.com/CJ-cooper6/WechatGPTBot.git

# 进入项目目录
cd WechatGPTBot

# 替换config.json中access_token

````
**4、启动项目**

````
go run main.go
````
