package queue

type BaseForm struct {
	Api string `json:"api"`
	Version string `json:"version"`
	TimeStamp int64 `json:"timestamp"`
	Data interface{} `json:"data"`
}

// 用户信息是三层结构需要多层解析
type ScrmUserReportFrom struct {
	BaseForm
	Data ScrmUserReportSecondFrom `json:"data"`
}

type ScrmUserReportSecondFrom struct {
	Api string `json:"api"`
	Exchange string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
	OnlineUserIds []int64 `json:"online_user_ids"`
	Data ScrmUserReportThirdFrom `json:"data"`
}

type ScrmUserReportThirdFrom struct {
	Mobile string `json:"mobile"`
	NickName string `json:"nick_name"`
	RoutingKey string `json:"routing_key"`
	Source string `json:"source"`
	TimeRegistration string `json:"time_registration"`
	UserNumber string `json:"user_number"`
}