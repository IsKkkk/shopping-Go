package user

import "regexp"

// 用户名正则表达式，最小8个字符，最大 30个字符，用户名字母打头
var usernameRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9_]{7,29}$")

// 密码正则表达式，最小8个字符，至少一个字符一个数字
var passwordRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{7,29}$`)

func ValidateUserName(name string) bool {
	return usernameRegex.MatchString(name)
}

func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)

}
