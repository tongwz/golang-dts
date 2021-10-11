package command

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang-dts/pkg/e"
	"golang-dts/pkg/queue"
)

func init()  {
	QueueCmd.ResetFlags()
}

var QueueCmd = &cobra.Command{
	Use: "queue",
	Short: "队列监听服务",
	Aliases: []string{
		"test-queue",
		"user-upload",
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("你要监听的队列参数没有展示")
		}
		_, ok := e.CommandString[args[0]]
		if !ok {
			fmt.Println(args[0]+ "队列并不存在！")
		}
		return nil
	},
	Long: `这里监听了中转中心需要的队列,
				包含但不限于，测试队列，用户绑定队列，用户上报，订单上报等。
				基于cobra插件。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化命令行队列
		rabbitMq := queue.NewRabbitMQRouting(e.CommandString[args[0]])
		// 监听消费队列
		rabbitMq.ReceiveRouting()
	},
}

//func Execute()  {
//	if err := QueueCmd.Execute(); err != nil {
//		fmt.Println(os.Stderr, err)
//		logging.Fatal(os.Stderr, err)
//		os.Exit(1)
//	}
//}
