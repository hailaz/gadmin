#### 项目不定期更新


### 简要说明

基于GoFrame框架的前后端分离管理后台，项目首次启动会自动插入管理员用户。

#### 使用方式

1. 准备数据库，数据结构在[docfile/sql](https://github.com/hailaz/gadmin/tree/master/docfile/sql)目录下，all.sql是全部结构。
2. 修改配置文件config/default.toml，也可指定其它配置文件 -c=xxx.toml
3. 启动编译后的二进制文件
4. 访问[http://localhost:8199](http://localhost:8199)
    用户名：admin
    密码：123456

**当前项目已包含构建后的前端文件。** 未构建的前端项目地址：[gadmin-vue](https://github.com/hailaz/gadmin-vue)

#### 自述
第一次设计，整体结构可能有些问题，望见谅，要是能给我个建议什么的就更好了。有bug可以提个issues。

![gadmin](/docfile/gadmin.png)


----

### 技术栈

#### 后端

框架：[GoFrame](https://github.com/gogf/gf)

登录加密：RSA

权限管理：[Casbin](https://github.com/casbin/casbin)的RBAC-restful

API授权方式：JWT，使用[gf-jwt](https://github.com/gogf/gf-jwt)

----
#### 前端
框架：[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)

----
----
### 如何访问接口文档？

访问[https://petstore.swagger.io](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/hailaz/gadmin/master/docfile/swagger.yaml)查看。

或者启动项目后访问[http://localhost:8199/swagger](http://localhost:8199/swagger)显示接口文档。

或将docfile目录下的swagger.yaml文件内容复制到[http://editor.swagger.io](http://editor.swagger.io)查看。


