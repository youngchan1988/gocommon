#Gocommon 工具包

| 安装                                                   | 模块          | 说明                           |
| ------------------------------------------------------ | ------------- | ------------------------------ |
| go get g.newcoretech.com/mobile/gocommon               | gocommon      | 常用判空、相等、包含等判断     |
| go get g.newcoretech.com/mobile/gocommon/cast          | cast          | interface 对其他数据类型的转换 |
| go get g.newcoretech.com/mobile/gocommon/decimalutils  | decimalutils  | 浮点数操作                     |
| go get g.newcoretech.com/mobile/gocommon/fileutils     | fileutils     | 文件操作                       |
| go get g.newcoretech.com/mobile/gocommon/pageutils     | pageutils     | 分页操作                       |
| go get g.newcoretech.com/mobile/gocommon/safelist      | safelist      | 线程安全列表                   |
| go get g.newcoretech.com/mobile/gocommon/safemap       | safemap       | 线程安全字典                   |
| go get g.newcoretech.com/mobile/gocommon/securityutils | securityutils | 常用加/解密，md5等             |
| go get g.newcoretech.com/mobile/gocommon/sliceutils    | sliceutils    | Slice常用操作                  |
| go get g.newcoretech.com/mobile/gocommon/stringutils   | stringutils   | 常用字符串操作                 |
| go get g.newcoretech.com/mobile/gocommon/syncutils     | syncutils     | 同步锁操作                     |



# gocommon私有仓库配置

## 仓库使用SSH

`go get`  默认使用`https`  拉取代码，如果使用ssh，则需要做个替换：

```
git config --global url."git@g.newcoretech.com:".insteadOf "https://g.newcoretech.com/"
```

如果本地没有配置ssh key，则需要配置 Access Token, 在gitlab中，进入`Gitlab`—>`Settings`—>`Access Tokens`，然后创建一个`personal access token`，这里权限最好选择只读(read_repository)。

添加Access Token:

```
git config --global http.extraheader "PRIVATE-TOKEN: YOUR_PRIVATE_TOKEN"
```

## 配置GOPRIVATE环境变量

`go get`  默认从代理查找并下载module，对于私有module需要绕过代理直接从仓库下载，这时配置`GOPRIVATE`  环境变量即可：

```
go env -w GOPRIVATE="g.newcoretech.com"
```

 完成以上配置后，执行`go get -u -v g.newcoretech.com/mobile/gocommon`  即可下载module代码

在GOLand 中配置`GOPRIVATE` 变量：

![企业微信20210413-102928](https://tva1.sinaimg.cn/large/008eGmZEgy1gphwt4avtuj30ra0k8tam.jpg)
=======
# GoCommon

#### 介绍
Go 项目常用工具代码

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

