# How to build and install

安装 Answer 之前,您需要先安装基本环境。
 - 数据库
     - [MySQL](http://dev.mysql.com)：版本 >= 5.7

然后，您可以通过以下几种种方式来安装 Answer：

 - 采用 Docker 部署
 - 二进制安装
 - 源码安装

## Docker for Answer
Visit Docker Hub or GitHub Container registry to see all available images and tags.

### Usage
To keep your data out of Docker container, we do a volume (/var/data -> /data) here, and you can change it based on your situation.

```
# Pull image from Docker Hub.
$ docker pull answer/answer

# Create local directory for volume.
$ mkdir -p /var/data

# Use `docker run` for the first time.
$ docker run --name=answer -p 9080:80 -v /var/data:/data answer/answer

# 第一次启动后会在/var/data 目录下生成配置文件
# /var/data/config.yaml
# 需要修改配置文件中的Mysql 数据库地址
vim /var/data/config.yaml

# 修改数据库连接 connection: [username]:[password]@tcp([host]:[port])/[DbName]
...

# Use `docker start` if you have stopped it.
$ docker start answer
```

## Binary for Answer
## 二进制安装

 1. 解压压缩包。
 2. 使用命令 cd 进入到刚刚创建的目录。
 3. 执行命令 ./answer init。
 4. Answer 会在当前目录生成./data 目录
 5. 进入data目录修改config.yaml文件
 6. 将数据库连接地址修改为你的数据库连接地址

     connection: [username]:[password]@tcp([host]:[port])/[DbName]
 7. 退出data 目录 执行 ./answer run -c ./data/config.yaml

## config.yaml 说明

```
server:
  http:
    addr: 0.0.0.0:80 #项目访问端口号
data:
  database:
    connection: root:root@tcp(127.0.0.1:3306)/answer #mysql数据库连接地址
  cache:
    file_path: "/tmp/cache/cache.db" #缓存文件存放路径
i18n:
  bundle_dir: "/data/i18n" #国际化文件存放目录
swaggerui:
  show: true #是否显示swaggerapi文档,地址 /swagger/index.html
  protocol: http #swagger 协议头
  host: 127.0.0.1 #可被访问的ip地址或域名
  address: ':80'  #可被访问的端口号
service_config:
  secret_key: "answer" #加密key
  web_host: "http://127.0.0.1" #页面访问使用域名地址
  upload_path: "./upfiles" #上传目录

```

# TODO
## 前端安装

## 后端安装

## 编译镜像

## 常见问题
