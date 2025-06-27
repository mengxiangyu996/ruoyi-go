<h1 align="center" style="margin: 30px 0 30px; font-weight: bold; font-size: 30px">RuoYi-Go</h1>
<h4 align="center">基于Go+Gin+Gorm实现的若依服务端脚手架</h4>

[中文](https://www.readme-i18n.com/mengxiangyu996/ruoyi-go?lang=zh) | 
[English](https://www.readme-i18n.com/mengxiangyu996/ruoyi-go?lang=en)


## 平台简介
* 本仓库为后端技术栈 [Gin](https://gin-gonic.com/zh-cn/docs) + [Gorm](https://gorm.io/zh_CN/docs/index.html) 的 `golang` 版本。
* 配套前端代码仓库地址 [RuoYi-Vue3](https://github.com/yangzongzhuan/RuoYi-Vue3) 或使用 [RuoYi-Vue3-ts](https://github.com/zzh948498/RuoYi-Vue3-ts)
* 其他生态组件请访问 **[若依官网](http://ruoyi.vip/)**

## 后端运行
> **提示：** 运行前请先安装好 `go` 环境，版本 `1.21` 以上。

    # 克隆项目
    git clone https://github.com/mengxiangyu996/ruoyi-go.git

    # 进入项目目录
    cd ruoyi-go

    # 修改配置文件
    cp application-example.yaml application.yaml

    # 安装依赖
    go mod tidy

    # 启动服务
    go run main.go

## 前端运行
    # 调整 .env 文件
    VUE_APP_BASE_API = '/dev-api' 改为 VITE_APP_BASE_API = '/api'

    # 调整 vite.config.js 文件
    server: {
      port: 8000,
      open: false,
      proxy: {
        // https://cn.vitejs.dev/config/#server-proxy
        '/api': {
          target: 'http://localhost:3000',
          changeOrigin: true,
          rewrite: (p) => p.replace(/^\/api/, '/api')
        }
      }
    },

    # 安装依赖
    npm install

    # 启动服务
    npm run dev

## 后端打包
    # 打包
    go build main.go

## 内置功能
1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9.  登录日志：系统登录日志记录查询包含登录异常。

## 特别感谢（排名不分先后）
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm)
- [Redis](https://github.com/redis/go-redis)
- [RuoYi-Vue3](https://github.com/yangzongzhuan/RuoYi-Vue3)
- [RuoYi-Vue3-ts](https://github.com/zzh948498/RuoYi-Vue3-ts)

## Star History
<a href="https://www.star-history.com/#mengxiangyu996/ruoyi-go&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=mengxiangyu996/ruoyi-go&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=mengxiangyu996/ruoyi-go&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=mengxiangyu996/ruoyi-go&type=Date" />
 </picture>
</a>
