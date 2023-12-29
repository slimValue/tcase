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



