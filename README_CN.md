<a href="https://answer.apache.org">
    <img alt="logo" src="docs/img/logo.svg" height="99px">
</a>

# Answer - 构建问答社区

一款问答形式的知识社区开源软件，你可以使用它快速建立你的问答社区，用于产品技术支持、客户支持、用户交流等。

了解更多关于该项目的内容，请访问 [answer.apache.org](https://answer.apache.org).

[![LICENSE](https://img.shields.io/github/license/apache/incubator-answer)](https://github.com/apache/incubator-answer/blob/main/LICENSE)
[![Language](https://img.shields.io/badge/language-go-blue.svg)](https://golang.org/)
[![Language](https://img.shields.io/badge/language-react-blue.svg)](https://reactjs.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/apache/incubator-answer)](https://goreportcard.com/report/github.com/apache/incubator-answer)
[![Discord](https://img.shields.io/badge/discord-chat-5865f2?logo=discord&logoColor=f5f5f5)](https://discord.gg/Jm7Y4cbUej)

## 截图

![screenshot](docs/img/screenshot.png)

## 快速开始

### 使用 docker 快速搭建

```bash
docker run -d -p 9080:80 -v answer-data:/data --name answer answerdev/answer:latest
```

其他安装配置细节请参考 [Installation](https://answer.apache.org/docs/installation)

## 从源码构建

### 环境要求
- Golang >= 1.18
- Node.js >= 16.17
- pnpm >= 8

### 构建
```bash
$ make ui
$ make build
```

## 从源码运行

### 前端构建

```bash
make ui
```

### 系统初始化和安装

```bash
INSTALL_PORT=9080 go run cmd/answer/main.go --data-path=/data init
```

其他安装配置细节请参考 [Installation](https://answer.apache.org/docs/installation)

### 本地运行

```bash
go run cmd/answer/main.go --data-path=/data run
```

## 贡献

我们随时欢迎你的贡献!

参考 [CONTRIBUTING](https://answer.apache.org/docs/development/contributing/) 开始贡献。

## License

[Apache License 2.0](https://github.com/apache/incubator-answer/blob/main/LICENSE)
