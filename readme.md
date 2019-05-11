### 简要说明

Demo：[gadmin](http://gadmin.xiao3.top)

后端使用框架：[GoFrame](https://github.com/gogf/gf)

前端使用框架：[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)

前端项目地址：[https://github.com/hailaz/gadmin-vue](https://github.com/hailaz/gadmin-vue)

----

### 技术栈

#### 后端

框架：GoFrame

登录加密：RSA

权限管理：[Casbin](https://github.com/casbin/casbin)的RBAC

API授权方式：JWT，使用[gf-jwt](https://github.com/gogf/gf-jwt)

----
#### 前端
框架：vue-element-admin

----
----
### 其它说明

login

登录需要先 get 请求/loginkey 获取 RSA 的加密公钥（单次有效），将密码使用 RSA-PKCS1v1.5 方式加密并使用 base64 编码

收到的数据

```js
{
  "code": 20000,
  "data": {
    "kid": "BU06PZLJPUNK90L44T",
    "cryptotype": "RSA-PKCS1v1.5",
    "key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsAj30qPj3YT7n4bm2uyr\nglOugXyvv34k/kwHkrS/KlLNsvbbGbmAv8fMakuPxm64LKXjX/16+P5IQgNuqWjg\nqoqllB+9Ex3fftEa+YE9hxIlkfBVrhV3SFxJe1bR2QbdgN/tNmnrmwjxCpgeyrCM\nKE/ifAkImgNfC6KCavY7idklRMv/Xji53H1TUoarN2kHFfde5NGfd/Z+9GoB633U\nWi6zJ6vMhinlCaPVQlBHHwIRqMlzC3CWXR3Tx42o4GgPktkZjWK0+DqUUhF7amXC\n/diypzSBsVmCEn68U52byW9vSHiWlRA0N9/stHCZ5M3NcnnZF9JV6XcYjlgHj3ZD\nEwIDAQAB\n-----END PUBLIC KEY-----\n",
    "timestamp": 1557542458
  },
  "message": "success"
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
