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
	cmd.AddCommand(command.QueueCmd)
	cmd.AddCommand(command.ApiCmd)
	err := cmd.Execute()
	if err != nil {
		logging.Fatal("服务启动失败，有错误", err)
		fmt.Println("服务启动失败", err)
		os.Exit(1)
	}
}
