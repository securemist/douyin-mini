# douyin-mini

### 项目介绍

本项目为第六届字节跳动青训营大项目，点击查看[详情](https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc)。


注意：本项目为oss版本，需要你开通阿里云oss对象存储才能运行起来。
### 启动方式

1. 在`conf`编辑你的mysql配置和oss配置。


2. 运行SQL文件`doc/douyin.sql`。


3. 运行`generate/generate.go` 生成测试数据。


```shell
go build -0 gen ./generate/generate.go
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

