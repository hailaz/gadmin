#readme

Demo：http://gadmin.xiao3.top

后端使用

https://github.com/gogf/gf

前端使用

https://github.com/PanJiaChen/vue-element-admin

fork 修改后

https://github.com/hailaz/gadmin-vue

后端技术栈

gframe 框架

登录使用 RSA 加密

权限使用 casbin

API 授权使用 JWT

casbin RBAC
https://github.com/casbin/casbin

JWT json web token

login
登录需要先 get 请求/loginkey 获取 RSA 的加密公钥（单次有效），将密码使用 RSA-PKCS1v1.5 方式加密并使用 base64 编码
收到的数据

```js
{
    "code":0,
    "data":{
        "kid":"BTNKI5TJP5PGNE83JP",
        "cryptotype":"RSA-PKCS1v1.5",
        "key":"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAz101HTbyTYlVY8ID/SgR\ntH7FVKoT5wbnSpAPH7qoNb9zyEl4klvFAoJVWi5haatIG0EPvEmC+/mR05KN0SIe\nsOAeaq/l0v+8wY6VYlAI+sMp5NFNLQzczgvEPxy9404sxQWn8BOJ/tH+84yA23py\nJBiIWjRLLII9+t/R7rNM0TQ0Iu6T/BZfxv0QXuW5AU803CveEg0kE6/kpnNu1aNK\nb7p2gcBiVqGA36AwogqO6Rr0KXTEVxFFYspeVezLBkyY5g1ChWjor/kcr8jLttW4\nTzyiOLtDtAzdQUSg1ECGOTZnD/MYG4G7JT4KtSFlmAdTMn1+D8Jn2xTSYJDQA0iR\nnwIDAQAB\n-----END PUBLIC KEY-----\n",
        "timestamp":1556261061
    },
    "msg":"success"
}
```

例如：
使用 forge.js 进行加密数据
https://raw.githubusercontent.com/loveiset/RSAForPostman/master/forge.js

var publicKey = forge.pki.publicKeyFromPem(公钥);
var encryptedText = forge.util.encode64(publicKey.encrypt(真实密码));

最终提交的格式

```
Content-Type: application/json
POST
{
    "username":"admin",
    "password":encryptedText,
    "kid":"BTNK5IF7DYC8287814"//公钥的 id
}
```