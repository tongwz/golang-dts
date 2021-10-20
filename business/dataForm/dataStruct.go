package dataForm

type BaseForm struct {
	Api string `json:"api"`
	Version int `json:"version"`
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
	OnlineUserIds []int `json:"online_user_ids"`
	Data map[string]interface{}
}

type ScrmUserReportThirdFrom struct {
	AppVersion string `json:"app_version"`
	BusinessId string `json:"business_id"`
	Channel string `json:"channel"`
	//IsProtected string `json:"is_protected,omitempty"`
	ReportTarget string `json:"report_target"`
	Mobile string `json:"mobile"`
	NickName string `json:"nick_name"`
	Source string `json:"source"`
	TimeRegistration string `json:"time_registration"`
	UserNumber string `json:"user_number"`
}