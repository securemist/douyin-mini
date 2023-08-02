# douyin-mini

### 项目介绍

本项目为第六届字节跳动青训营大项目，点击查看[详情](https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc)。

### 启动方式

1. 编辑你的mysql配置和oss配置。


项目根目录创建两个文件

`mysql.properties`

```properties
addr=127.0.0.1:3306
username=root
password=root
database=douyin
```

`oss.properties`
```properties
endpoint=
acessKeyId=
accessKeySecret=
bucketName=
```


2. 运行SQL文件`doc/douyin.sql`。


3. 运行`generate/generate.go` 生成测试数据。


4. 启动项目

```shell
go build && ./douyin-mini.exe
```

### 开发进度

项目详细的接口文档在[这里](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707525)。

- [x] 基础接口
- [x] 互动接口
- [ ] 社交接口



### 参与项目

欢迎参与本项目。

联系方式：gongjiatian00@qq.com。

