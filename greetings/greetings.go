package greetings

import "fmt"

// 这里是注释部分
func Hello(name string) string {
	// := 是 var message string的简写
	message := fmt.Sprintln("Hi, %v. Welcome!", name)
	return message
}
