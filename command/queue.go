package command

import (
	"errors"
	"github.com/spf13/cobra"
	"golang-dts/pkg/e"
	"golang-dts/pkg/logging"
	"golang-dts/pkg/queue"
)

func init()  {
	QueueCmd.ResetFlags()
}

var QueueCmd = &cobra.Command{
	Use: "queue",
	Short: "队列监听服务",
	Aliases: aliasList(),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("你要监听的队列参数没有展示")
		}
		_, ok := e.CommandString[args[0]]
		if !ok {
			return errors.New(args[0]+ "队列并不存在！")
		}
		return nil
	},
	Long: `这里监听了中转中心需要的队列,
				包含但不限于，测试队列，用户绑定队列，用户上报，订单上报等。
				基于cobra插件。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化命令行队列
		rabbitMq := queue.NewRabbitMQRouting(e.CommandString[args[0]])
		// 获取实现消费的逻辑 实例化的结构体
		var receiveObj = e.Receive(e.CommandString[args[0]])
		// 设置日志名称
		logging.SetName(receiveObj.SetLogName())
		// 监听消费队列
		rabbitMq.ReceiveRouting(receiveObj)
	},
}

//  获取所有的queue的入参
func aliasList() []string {
	// 设置长度可以不需要每次append 都申请新的内存空间
	var aList = make([]string, 0, len(e.CommandString))
	for aliasStr :=  range  e.CommandString {
		aList = append(aList, aliasStr)
	}
	return aList
}

//func Execute()  {
//	if err := QueueCmd.Execute(); err != nil {
//		fmt.Println(os.Stderr, err)
//		logging.Fatal(os.Stderr, err)
//		os.Exit(1)
//	}
//}
