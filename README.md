# golang-dts
php中转中心，golang重写项目，很可能烂尾

## 文件夹介绍：
1. command：是进程分组的初始化文件，暂时包含api（接口），queue（基于rabbitmq的队列）
2. conf：配置文件
3. middleware：中间件（验证登录token）
4. models：数据库表格操作
5. pkg：基础常量，配置，分页，日志等的基础函数
6. routers：api路由和部分测试controller代码
7. runtime：临时文件和日志文件存放地
8. vendor：基础包
