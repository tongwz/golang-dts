package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang-dts/pkg/setting"
	"golang-dts/routers"
	"net/http"
)

var ApiCmd = &cobra.Command{
	Use: "api",
	Short: "接口监听服务",
	Args: cobra.NoArgs,
	Long: `这里监听了中转中心需要的接口,
				包含但不限于，测试接口，解密接口，刷新用户绑定关系接口，修复数据接口等。
				基于cobra插件。`,
	Run: func(cmd *cobra.Command, args []string) {
		router := routers.InitRouter()

		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
			Handler:        router,
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
	},
}