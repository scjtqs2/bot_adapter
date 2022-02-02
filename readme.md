# cqhttp 的适配器

1. 采用webservice的方式集成插件，每个插件需要一个appid和appsecret。
2. 对每个app都采用推送事件的方式上报。可以注册关键词拦截、推送权限赛选。
3. app使用grpc的接口调用bot_adapter实现的cqhttp的onebot11接口。带权限管理

参数配置参考 `demo.config.yaml`。以 config.yaml命名