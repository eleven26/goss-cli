# goss-cli

上传下载对象存储文件的命令行工具，兼容 amazon s3 协议，支持但不限于阿里云、腾讯云、华为云、七牛云、amazon s3、minio。

## 安装

> 注：需要将 `$GOPATH/bin` 加到环境变量

```shell
go install github.com/eleven26/goss-cli@latest
```

## 配置

需要添加以下环境变量：

```shell
export GOSS_ENDPOINT=oss-cn-hangzhou.aliyuncs.com
export GOSS_ACCESS_KEY=xxx
export GOSS_SECRET_KEY=xxx
export GOSS_REGION=oss-cn-hangzhou
export GOSS_BUCKET=xxx
```

## 使用

1. 查看用法

```shell
goss-cli -h
```

输出：

```console
Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         获取指定文件
  delete      删除指定文件
  help        Help about any command
  list        列出指定目录下的文件
  put         上传文件
```

2. 列出指定前缀的文件：下面的 `mac` 是 `key` 的前缀

```shell
goss-cli list mac
```

3. 下载文件：这会下载文件到当前目录

```shell
goss-cli get xx.txt
```

4. 上传文件

* 上传当前目录的文件：

```shell
goss-cli put xx.txt
```

* 上传其他目录的文件：

```shell
goss-cli put /path/to/xx.txt
```

* 上传文件并指定 `key`：

```shell
goss-cli put /path/to/xx.txt --key=new_key.txt
```

5. 删除文件

```shell
goss-cli delete xx.txt
```
