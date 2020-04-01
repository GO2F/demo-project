// 配置列表
// 这里可以随便写

package config

import (
	"os"
	"strings"
)

// Env 环境变量
const (
	Dev  = "development"
	Test = "testing"
	Prod = "production"
)

// Current 当前环境
var Current = Dev

// 导出结果
func envInit() {
	var args string = strings.Join(os.Args[1:], "")

	if strings.Contains(args, Dev) {
		Current = Dev
	}
	if strings.Contains(args, Test) {
		Current = Test
	}
	if strings.Contains(args, Prod) {
		Current = Prod
	}
}
