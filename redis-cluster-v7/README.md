# 解决redis cluster v7 兼容问题

> 7.05 版本

## 背景

中间件版本升级，需要支持redis v7的集群模式, 项目用go-redis库 版本 v6.x


连接新版 redis集群会报错：
`got 4 elements in cluster info address, expected 2 or 3`


## 解决

https://github.com/redis/go-redis/issues/2085
https://help.aliyun.com/zh/redis/support/common-errors-and-troubleshooting#redis-external-sdk-goredis-1101


结论: 要么降redis版本 v6，要么升级go-redis库到v9。我们选了第二种，便于dba统一的管理，一律用最新版。



## 测试用例



## 附录工具

### go 模块依赖

这里介绍个工具 gmchart，找出直接依赖


进入工作项目

`cd xxx`

安装 gmchart

`go install github.com/PaulXu-cn/go-mod-graph-chart/gmchart@latest`


运行

`go mod graph | gmchart`


```
go mod graph 是官方工具命令。 可展示出了该项目所有的依赖关系，只不过是文本形式展示，输出的内容多了，人眼看不出啥来。这里借用 gmchart 工具，可以将其依赖关系组织为 树状 渲染 web 页面，也就是和 go 工具一样，跨平台的。
```
