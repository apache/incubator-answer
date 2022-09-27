# answer

问答社区主项目代码

# Dependence
 github.com/segmentfault/pacman
 * config-file `viper` https://github.com/spf13/viper
 * web `gin` https://gin-gonic.com/zh-cn/
 * log `zap` https://github.com/uber-go/zap
 * orm `xorm` https://xorm.io/zh/
 * redis `go-redis` https://github.com/go-redis/redis

# module
 - email github.com/jordan-wright/email
 - session github.com/gin-contrib/sessions
 - Captcha github.com/mojocn/base64Captcha

# Run
```
cd cmd
export GOPRIVATE=git.backyard.segmentfault.com
go mod tidy
./dev.sh
```

# pprof

```
 # Installation dependency
 go get -u github.com/google/pprof
 brew install graphviz
```
```
pprof -http :8082 http://XXX/debug/pprof/profile\?seconds\=10
```
