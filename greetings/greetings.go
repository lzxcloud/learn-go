package greetings

import (
	"errors"
	"fmt"
)

// 这里是注释部分
func Hello(name string) (string, error) {
	// := 是 var message string的简写
	//增加对错误的提示
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
