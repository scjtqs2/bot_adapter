# 一套配套demo 使用
> 这些demo可以直接使用，也可以用于参考开发。
> 
> 项目地址如下：
> 
+ go-cqhttp 机器人主体
  > https://github.com/scjtqs2/go-cqhttp
+ bot-adapter 适配器本体
  > https://github.com/scjtqs2/bot_adapter
+ bot-music 音乐插件demo
  > https://github.com/scjtqs2/bot_app_music
+ bot-request-add 自动同意加群插件demo
  > https://github.com/scjtqs2/bot_app_request_add
+ bot-search 取码demo
  > https://github.com/scjtqs2/bot_app_codefetch
+ bot-chat 图灵机器人
  > https://github.com/scjtqs2/bot_app_chat
  > 
+ bot-github github查询/事件推送
  > https://github.com/scjtqs2/bot_app_github
  > 
  > webhook地址 POST/JSON http://ip:8111/postreceive  

## 使用docker-compose
将当前的example目录放入到你的服务器目录

`docker-compose up -d` 启用服务

浏览器打开 http://ip:9999/admin 登录控制台 默认账户名和密码都是 `admin`，请及时更改

支持账号密码登录和扫码登录qq。后面的就请自行配置了。

请优先修改 http的token，改成你docker-compose.yaml中填些的token。

## config.yaml
默认已配置了一些，如果想修改成其他的配置，可以自行修改。不了解yaml的配置，可以在go-cqhttp的web控制台配置。
