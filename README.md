# goss-cli

上传下载对象存储文件的命令行工具，支持阿里云、腾讯云、华为云、七牛云。

## 安装

> 注：需要将 `$GOPATH/bin` 加到环境变量

```shell
go install github.com/eleven26/goss-cli@latest
```

## 配置

1. 配置文件路径为：`~/.goss.yml`，`goss-cli` 执行的时候会去读取 `~/.goss.yml`，因此需要保证该文件为有效配置文件。

2. 参考配置文件：

* `driver`: 指明使用哪个云存储提供商的对象存储服务。
* `aliyun`、`tencent`、`qiniu`、`huawei`、`s3`、`minio` 等为不同厂商对象存储的配置。

```yaml
# 云存储类型
driver: aliyun

aliyun:
  # oss 的链接，不同区域不同
  endpoint:
  # bucket
  bucket:
  access_key_id:
  access_key_secret:

tencent:
  url:
  secret_id:
  secret_key:

qiniu:
  bucket:
  access_key:
  secret_key:
  domain:
  private:

huawei:
  endpoint:
  location:
  bucket:
  access_key:
  secret_key:

s3:
  endpoint:
  region:
  bucket:
  access_key:
  secret_key:

minio:
  endpoint:
  bucket:
  access_key:
  secret_key:
  use_ssl:
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
  debug       调试命令
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
