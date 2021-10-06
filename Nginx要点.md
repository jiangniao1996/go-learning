# Nginx



## 调优配置

###### 背景：

![image-20211004171501133](.\imgs\nginx-plan-before.png)

随着业务量增大并发量也随之提高，一台服务器满足不了业务需求

###### 方案：

![image-20211004172324840](.\imgs\nginx-plan-after.png)

**可做到反向代理以及负载均衡**



```nginx
upstream www_server_pool {
    server 10.0.0.7 weight=5;
    server 10.0.0.8 weight=10;
}
```























