package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang-dts/pkg/logging"
	"os"

	"golang-dts/command"
)

func main() {

	var cmd = &cobra.Command{Use:"comm"}
	// 添加command 的备用选项
	cmd.AddCommand(command.QueueCmd)
	cmd.AddCommand(command.ApiCmd)
	err := cmd.Execute()
	// 启动失败的后续处理
	if err != nil {
		logging.Fatal("服务启动失败，有错误", err)
		fmt.Println("服务启动失败", err)
		os.Exit(1)
	}
}
