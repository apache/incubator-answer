package checker

import (
	"fmt"
	"regexp"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

/*
 *  minLength: 指定密码的最小长度
 *  maxLength：指定密码的最大长度
 *  minLevel：指定密码最低要求的强度等级
 *  pwd：明文密码
 */
func PassWordCheck(minLength, maxLength, minLevel int, pwd string) error {
	// 首先校验密码长度是否在范围内
	if len(pwd) < minLength {
		return fmt.Errorf("BAD PASSWORD: The password is shorter than %d characters", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("BAD PASSWORD: The password is logner than %d characters", maxLength)
	}

	// 初始化密码强度等级为D，利用正则校验密码强度，若匹配成功则强度自增1
	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	// 如果最终密码强度低于要求的最低强度，返回并报错
	if level < minLevel {
		return fmt.Errorf("The password does not satisfy the current policy requirements. ")
	}
	return nil
}
