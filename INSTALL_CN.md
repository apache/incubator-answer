# Answer 安装指引

## 使用 docker 安装
### 步骤 1: 使用 docker 命令启动项目
```bash
docker run -d -p 9080:80 -v answer-data:/data --name answer answerdev/answer:latest
```
### 步骤 2: 访问安装路径进行项目安装
[http://127.0.0.1:9080/install](http://127.0.0.1:9080/install)

选择语言后点击下一步选择合适的数据库，如果当前只是想体验，建议直接选择 sqlite 作为数据库，如下图所示

![install-database](docs/img/install-database.png)

然后点击下一步会进行配置文件创建等操作，点击下一步输入网站基本信息和管理员信息，如下图所示

![install-site-info](docs/img/install-site-info.png)

点击下一步即可安装完成

### 步骤 3：安装完成后访问项目路径开始使用
[http://127.0.0.1:9080/](http://127.0.0.1:9080/)

使用刚才创建的管理员用户名密码即可登录。

## 使用 docker-compose 安装
### 步骤 1: 使用 docker-compose 命令启动项目
```bash
mkdir answer && cd answer
wget https://raw.githubusercontent.com/answerdev/answer/main/docker-compose.yaml
docker-compose up
```

### 步骤 2: 访问安装路径进行项目安装
[http://127.0.0.1:9080/install](http://127.0.0.1:9080/install)

具体配置与 docker 使用时相同

### 步骤 3：安装完成后访问项目路径开始使用
[http://127.0.0.1:9080/](http://127.0.0.1:9080/)

## 使用 二进制 安装
### 步骤 1: 下载二进制文件
[https://github.com/answerdev/answer/releases](https://github.com/answerdev/answer/releases)
请下载您当下系统所需要的对应版本

### 步骤 2: 使用命令行安装
> 以下命令中 -C 指定的是 answer 所需的数据目录，您可以根据实际需要进行修改

```bash
./answer init -C ./answer-data/
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
  show: true #是否显示swaggerapi文档，地址 /swagger/index.html
  protocol: http #swagger 协议头
  host: 127.0.0.1 #可被访问的ip地址或域名
  address: ':80'  #可被访问的端口号
service_config:
  secret_key: "answer" #加密key
  web_host: "http://127.0.0.1" #页面访问使用域名地址
  upload_path: "./upfiles" #上传目录
```

## 编译镜像
如果修改了源文件并且要重新打包镜像可以使用以下语句重新打包镜像
```
docker build -t  answer:v1.0.0 .
```
## 常见问题
 1. 项目无法启动，answer 主程序启动依赖配置文件 config.yaml 、国际化翻译目录 /i18n 、上传文件存放目录 /upfiles，需要确保项目启动时加载了配置文件 answer run -c config.yaml 以及在 config.yaml 正确的指定 i18n 和 upfiles 目录的配置项
