# Rank

demo 项目

# 技术选型

前端: React
后端: Go

# 运行指引

## 基础环境配置

开发工具: VS Code

环境依赖

Go: 1.3+ [下载地址](https://golang.org/dl/)

Node.js: 8.6+ [下载地址](https://nodejs.org/zh-cn/)

MySQL 配置

```
Host:     "localhost",
Port:     "3306",
User:     "root",
Password: "123456",
Database: "go2fe-demo",
```

# 项目初始化

1.  进入项目根目录
2.  启动服务`go run main.go`
3.  访问`http://localhost:5060/debug/init/database`初始化数据库
4.  安装开发工具
    1.  `go get github.com/pilu/fresh`
        1.  项目文件发生更改后自动重启, 类似 nodemon
    2.  `go get -v github.com/go-delve/delve/cmd/dlv`
        1.  go 语言 debug 工具

## 项目开发

### 启动

在项目根目录下

前端: `cd client && npm run start`
后端: `fresh`(支持自动重启) / `go run main.go`(传统启动方式)

访问`http://localhost:5060`, 调试项目
