package consts

const (
	ServiceName string = "RecruitBe"
)

const (
	// 数据状态
	RowStatusNormal = 0
	RowStatusDelete = 4

	// 分页
	DefaultOffset   int64 = 0   //默认页码
	DefaultLimit    int64 = 10  //默认每页数量
	DefaultMaxLimit int64 = 500 //最大每页数量

	ROW_STATUS_NORMAL int = 0
	ROW_STATUS_DELETE int = 4

	// 格式化时间戳类型
	FormatTimestampType1 int = 1 //Y-m-d H:i:s
	FormatTimestampType2 int = 2 //Y-m-d H:i
	FormatTimestampType3 int = 3 //Y-m-d
	FormatTimestampType4 int = 4 //m-d-y
	FormatTimestampType5 int = 5 //Ym
	FormatTimestampType6 int = 6 //Y-m
)
