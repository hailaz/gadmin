```js
//路由说明
{
  path: '/permission',
  component: Layout,
  redirect: '/permission/index', //重定向地址，在面包屑中点击会重定向去的地址。当设置 noredirect 的时候该路由在面包屑导航中不可被点击
  hidden: true, // 不在侧边栏线上
  alwaysShow: true, //一直显示根路由
  meta: {
    roles: ['admin','editor']
  }, //你可以在根路由设置权限，这样它下面所以的子路由都继承了这个权限
  children: [{
    path: 'index',
    component: ()=>import('permission/index'),
    name: 'permission',
    meta: {
      title: 'permission',
      icon: 'lock', //图标
      role: ['admin','editor'], //或者你可以给每一个子路由设置自己的权限
      noCache: true // 不会被 <keep-alive> 缓存
    }
  }]
}

```
