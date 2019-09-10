### 简要说明

基于GoFrame框架的前后端分离管理后台

项目数据库结构在[docfile/sql](https://github.com/hailaz/gadmin/tree/master/docfile/sql)目录下

在线示例：[gadmin](http://home.xiao3.top:8199)

对应的前端项目：[gadmin-vue](https://github.com/hailaz/gadmin-vue)

自述：第一次设计，整体结构可能有些问题，望见谅，要是能给我个建议什么的就更好了。暂时也没想到增加什么功能，有需要可以提个issues，看看我能不能做。

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


