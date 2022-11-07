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
