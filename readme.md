# cqhttp 的适配器

1. 采用webservice的方式集成插件，每个插件需要一个appid和appsecret。
2. 对每个app都采用推送事件的方式上报。可以注册关键词拦截、推送权限赛选。
3. app使用grpc的接口调用bot_adapter实现的cqhttp的onebot11接口。带权限管理

参数配置参考 `demo.config.yaml`。以 config.yaml命名

## 推送说明：

+ 以json的方式进行post请求

```json
{
  "encrypt": "VFFFQ0xGUUhFVFpXR1JVV6wT7C5/iBIlGatuAwn/GIgAVTLImK8Aufo//pE5MDUJNjOHJY8FYN7MoInqvb2adNrPpWdzodlugpRE7/Xi0QZonYoqRHGgYEmp2aqZ2/qAVTvZ6HvwLVZuNZ1DX7SGf+CSEL6EquhI97ucU1sZq3f3P6MNEJA6ZHJq+yFvCLW66MOu7NH6RDldBzEjDfcbrbZp7E+kezty4GSQPKCkI0fMmUwn1PUUuCEL4OiKFQj6DLTkctnGr7wv+MlofDHUOCu8r5Txgr5yHy0jHJs57+lyeL7CNXBNzEacXxqtu8yO4fTzCrA0EUq+cnZ8ZNRV7Y77pY2K6bFFT+g8rmQn3cPZvg8CSikwTAbPtvgIfj8hukMgjAJx5PyrpdVYceeGNmP92zw5vuBq0X9xVISzMFsd+Q0PK9/gWHJKZhAS+cpAYpf4Z89kMGnpiee73Aw2CledB3ntky1iY4XnQGLreZiyfduZhvqI6b1c1O8X/WKYgb7bG7vkun5s29ufA==",
  "app_id": "32"
}
```

    1. `encrypt`:进行 aes256-cbc 加密后的字符串。以plugins中的`encrypt_key`字段为key
    2. `app_id`:该推送对应的appid，对于多appid的服务来匹配`encrypt_key`用。

+ 推送签名校验 参数从header中提取：
    1. `X-Lark-Request-Timestamp` 时间戳
    2. `X-Lark-Request-Nonce` nonce 随机字符串
    3. `X-Lark-Signature` 当前推送的签名，用于和自己计算的签名做对比

  加密和签名方法，在本项目的sha256里面，参考自`飞书`的推送

## example demo

docker-compose方式的 使用 [demo](example)