package healthcheckin

type PostModel struct {
	Name          string `json:"name"`
	TimeStamp     int64  `json:"timestamp"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Country       string `json:"country"`
	AnswerJsonStr string `json:"answerJsonStr"`
}

type PersonInfo struct {
	Grade      string `json:"GRADE"`
	StaffID    string `json:"STAFFID"`
	StaffName  string `json:"STAFFNAME"`
	StaffType  string `json:"STAFFTYPE"`
	StaffState string `json:"STAFSTATE"`
	UnitCode   string `json:"UNITCODE"`
}

type DailyStatusData struct {
	Name          string `json:"name"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Country       string `json:"country"`
	AnswerJsonStr string `json:"answerJsonStr"`

	Creator    string `json:"creator"`
	ReportTime string `json:"reportTime"`
	XueYuan    string `json:"xueyuan"`
}

type DailyStatus struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    DailyStatusData `json:"data"`
}
