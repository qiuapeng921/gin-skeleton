# micro-scrm 

#### 目录结构
```
.
├── README.md
├── cmd                         
│   └── http
│       ├── api.go
│       ├── endpoints.go
│       ├── main.go
├── db                              数据库相关脚本
│   ├── migrations                  数据库版本管理脚本
│   └── templates                   程序中使用到的查询 sql
│       └── t_demo                  表名或模块名。
│           ├── insert.sqltpl       命名 sql。每个 sql 一个文件
│           ├── select-id.sqltpl
│           └── select.sqltpl
├── go.mod
├── go.sum
├── internal
│   ├── assets
│   │   └── assets.go
│   ├── config                      应用程序配置
│   │   ├── config.go
│   │   └── struct.go
│   ├── modules
│   │   └── demo
│   │       ├── delivery            对外发布接口
│   │       │   ├── http            http 方式发布
│   │       │   │   ├── handle.go   http 发布方式具体实现封装
│   │       │   │   └── route.go    返回 http 实现是路由描述，初始化领域服务，选择具体的实现 respository
│   │       │   └── translator.go   领域实现和传输转换类
│   │       ├── init.go             提供领域服务初始化方式，供 delivery 调用，注入 具体实现的 respository
│   │       ├── entity.go           实体
│   │       ├── mocks               测试 mock
│   │       ├── repository.go       存储接口
│   │       ├── respository         存储具体实现
│   │       │   └── mysql
│   │       │       ├── impl.go
│   │       │       └── init.go
│   │       └── usercase.go         应用服务
│   └── pkg                         内部类库
│       ├── cache
│       │   ├── cache.go
│       │   └── goodsCache.go
│       ├── db
│       │   └── init.go
│       ├── middleware
│       │   └── access.go
│       ├── mysqlRepositoyBase.go   
│       ├── redis
│       │   └── init.go
│       └── tracing
│           └── skywalking.go
├── pkg                         本服务对外提供的 sdk
│   └── README.md
```
