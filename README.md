# douyin-mini

### 项目介绍

本项目为第六届字节跳动青训营大项目，点击查看[详情](https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc)。

该分支为本地存储版，所有视频资源使用本地存储，经过部署测试，这种方式视频资源下载极慢，依赖于网络带宽，平均一个视频加载5s。


### 启动方式

1. 在`conf`编辑配置文件，如果用服务器部署务必标明ip。


2. 运行SQL文件`doc/douyin.sql`。


3. 在`static`目录下创建`video`和`cover`两个文件夹，运行以下命令生成测试数据(大约20-30s)

```shell
go build -o ./gen ./generate/generate.go
./gen
```

4. 启动项目
```shell
go build 
./douyin-mini
```


### 开发进度


项目详细的接口文档在[这里](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707525)。

- [x] 基础接口
- [x] 互动接口
- [x] 社交接口


项目目前只完成了最基本的接口，后续开发目标：

- [x] 本地存储版本
- [ ] docker一键部署版
- [ ] 微服务版本

### 参与项目

欢迎参与本项目。

联系方式：gongjiatian00@qq.com。

